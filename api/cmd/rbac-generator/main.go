package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/golang/glog"
	rbaccontroller "github.com/kubermatic/kubermatic/api/pkg/controller/rbac"
	kubermaticclientset "github.com/kubermatic/kubermatic/api/pkg/crd/client/clientset/versioned"
	"github.com/kubermatic/kubermatic/api/pkg/crd/client/informers/externalversions"
	"github.com/kubermatic/kubermatic/api/pkg/signals"

	kuberinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type controllerRunOptions struct {
	kubeconfig string
	masterURL  string

	workerName  string
	workerCount int
}

type controllerContext struct {
	runOptions                      controllerRunOptions
	stopCh                          <-chan struct{}
	kubeMasterClient                kubernetes.Interface
	kubermaticMasterClient          kubermaticclientset.Interface
	kubermaticMasterInformerFactory externalversions.SharedInformerFactory
	kubeMasterInformerFactory       kuberinformers.SharedInformerFactory
	seedClusterProviders            []*rbaccontroller.ClusterProvider
}

func main() {
	ctrlCtx := controllerContext{}
	flag.StringVar(&ctrlCtx.runOptions.kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&ctrlCtx.runOptions.masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&ctrlCtx.runOptions.workerName, "worker-name", "", "The name of the worker that will only processes resources with label=worker-name.")
	flag.IntVar(&ctrlCtx.runOptions.workerCount, "worker-count", 4, "Number of workers which process the clusters in parallel.")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags(ctrlCtx.runOptions.masterURL, ctrlCtx.runOptions.kubeconfig)
	if err != nil {
		glog.Fatal(err)
	}

	ctrlCtx.stopCh = signals.SetupSignalHandler()
	ctrlCtx.kubeMasterClient = kubernetes.NewForConfigOrDie(config)
	ctrlCtx.kubermaticMasterClient = kubermaticclientset.NewForConfigOrDie(config)
	ctrlCtx.kubermaticMasterInformerFactory = externalversions.NewSharedInformerFactory(ctrlCtx.kubermaticMasterClient, time.Minute*5)
	ctrlCtx.kubeMasterInformerFactory = kuberinformers.NewSharedInformerFactory(ctrlCtx.kubeMasterClient, time.Minute*5)
	ctrlCtx.seedClusterProviders = []*rbaccontroller.ClusterProvider{}
	{
		clientcmdConfig, err := clientcmd.LoadFromFile(ctrlCtx.runOptions.kubeconfig)
		if err != nil {
			glog.Fatal(err)
		}

		for ctxName := range clientcmdConfig.Contexts {
			clientConfig := clientcmd.NewNonInteractiveClientConfig(
				*clientcmdConfig,
				ctxName,
				&clientcmd.ConfigOverrides{CurrentContext: ctxName},
				nil,
			)
			cfg, err := clientConfig.ClientConfig()
			if err != nil {
				glog.Fatal(err)
			}
			if cfg.Host == config.Host && cfg.Username == config.Username && cfg.Password == config.Password {
				glog.V(2).Infof("Skipping adding %s as a seed cluster. It is exactly the same as existing kubernetes master client", ctxName)
				continue
			}

			glog.V(2).Infof("Adding %s as seed cluster", ctxName)
			kubeClient, err := kubernetes.NewForConfig(cfg)
			if err != nil {
				glog.Fatal(err)
			}
			kubeInformerFactory := kuberinformers.NewSharedInformerFactory(kubeClient, time.Minute*5)
			kubermaticClient := kubermaticclientset.NewForConfigOrDie(config)
			kubermaticInformerFactory := externalversions.NewSharedInformerFactory(kubermaticClient, time.Minute*5)
			ctrlCtx.seedClusterProviders = append(ctrlCtx.seedClusterProviders, rbaccontroller.NewClusterProvider(fmt.Sprintf("seed/%s", ctxName), kubeClient, kubeInformerFactory, kubermaticClient, kubermaticInformerFactory))
		}
	}

	err = startRBACGeneratorController(ctrlCtx)
	if err != nil {
		glog.Fatal(err)
	}

	ctrlCtx.kubermaticMasterInformerFactory.Start(ctrlCtx.stopCh)
	ctrlCtx.kubeMasterInformerFactory.Start(ctrlCtx.stopCh)
	for _, seedClusterProvider := range ctrlCtx.seedClusterProviders {
		seedClusterProvider.StartInformers((ctrlCtx.stopCh))
	}

	<-ctrlCtx.stopCh
}

func startRBACGeneratorController(ctrlCtx controllerContext) error {
	ctrl, err := rbaccontroller.New(
		rbaccontroller.NewMetrics(),
		ctrlCtx.runOptions.workerName,
		ctrlCtx.kubermaticMasterClient,
		ctrlCtx.kubermaticMasterInformerFactory,
		ctrlCtx.kubeMasterClient,
		ctrlCtx.kubeMasterInformerFactory.Rbac().V1().ClusterRoles(),
		ctrlCtx.kubeMasterInformerFactory.Rbac().V1().ClusterRoleBindings(),
		ctrlCtx.seedClusterProviders)
	if err != nil {
		return err
	}
	go ctrl.Run(ctrlCtx.runOptions.workerCount, ctrlCtx.stopCh)
	return nil
}
