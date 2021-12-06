/*
Copyright 2021 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	corev1 "kubevirt.io/api/core/v1"
)

// FakeVirtualMachineInstancePresets implements VirtualMachineInstancePresetInterface
type FakeVirtualMachineInstancePresets struct {
	Fake *FakeKubevirtV1
	ns   string
}

var virtualmachineinstancepresetsResource = schema.GroupVersionResource{Group: "kubevirt.io", Version: "v1", Resource: "virtualmachineinstancepresets"}

var virtualmachineinstancepresetsKind = schema.GroupVersionKind{Group: "kubevirt.io", Version: "v1", Kind: "VirtualMachineInstancePreset"}

// Get takes name of the virtualMachineInstancePreset, and returns the corresponding virtualMachineInstancePreset object, and an error if there is any.
func (c *FakeVirtualMachineInstancePresets) Get(ctx context.Context, name string, options v1.GetOptions) (result *corev1.VirtualMachineInstancePreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineinstancepresetsResource, c.ns, name), &corev1.VirtualMachineInstancePreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.VirtualMachineInstancePreset), err
}

// List takes label and field selectors, and returns the list of VirtualMachineInstancePresets that match those selectors.
func (c *FakeVirtualMachineInstancePresets) List(ctx context.Context, opts v1.ListOptions) (result *corev1.VirtualMachineInstancePresetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineinstancepresetsResource, virtualmachineinstancepresetsKind, c.ns, opts), &corev1.VirtualMachineInstancePresetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &corev1.VirtualMachineInstancePresetList{ListMeta: obj.(*corev1.VirtualMachineInstancePresetList).ListMeta}
	for _, item := range obj.(*corev1.VirtualMachineInstancePresetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineInstancePresets.
func (c *FakeVirtualMachineInstancePresets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineinstancepresetsResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineInstancePreset and creates it.  Returns the server's representation of the virtualMachineInstancePreset, and an error, if there is any.
func (c *FakeVirtualMachineInstancePresets) Create(ctx context.Context, virtualMachineInstancePreset *corev1.VirtualMachineInstancePreset, opts v1.CreateOptions) (result *corev1.VirtualMachineInstancePreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineinstancepresetsResource, c.ns, virtualMachineInstancePreset), &corev1.VirtualMachineInstancePreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.VirtualMachineInstancePreset), err
}

// Update takes the representation of a virtualMachineInstancePreset and updates it. Returns the server's representation of the virtualMachineInstancePreset, and an error, if there is any.
func (c *FakeVirtualMachineInstancePresets) Update(ctx context.Context, virtualMachineInstancePreset *corev1.VirtualMachineInstancePreset, opts v1.UpdateOptions) (result *corev1.VirtualMachineInstancePreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineinstancepresetsResource, c.ns, virtualMachineInstancePreset), &corev1.VirtualMachineInstancePreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.VirtualMachineInstancePreset), err
}

// Delete takes name of the virtualMachineInstancePreset and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineInstancePresets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(virtualmachineinstancepresetsResource, c.ns, name), &corev1.VirtualMachineInstancePreset{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineInstancePresets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachineinstancepresetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &corev1.VirtualMachineInstancePresetList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineInstancePreset.
func (c *FakeVirtualMachineInstancePresets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *corev1.VirtualMachineInstancePreset, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineinstancepresetsResource, c.ns, name, pt, data, subresources...), &corev1.VirtualMachineInstancePreset{})

	if obj == nil {
		return nil, err
	}
	return obj.(*corev1.VirtualMachineInstancePreset), err
}
