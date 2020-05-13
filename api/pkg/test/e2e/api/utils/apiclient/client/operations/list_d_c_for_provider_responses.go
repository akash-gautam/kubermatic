// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// ListDCForProviderReader is a Reader for the ListDCForProvider structure.
type ListDCForProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListDCForProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListDCForProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListDCForProviderUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListDCForProviderForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListDCForProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListDCForProviderOK creates a ListDCForProviderOK with default headers values
func NewListDCForProviderOK() *ListDCForProviderOK {
	return &ListDCForProviderOK{}
}

/*ListDCForProviderOK handles this case with default header values.

Datacenter
*/
type ListDCForProviderOK struct {
	Payload []*models.Datacenter
}

func (o *ListDCForProviderOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc][%d] listDCForProviderOK  %+v", 200, o.Payload)
}

func (o *ListDCForProviderOK) GetPayload() []*models.Datacenter {
	return o.Payload
}

func (o *ListDCForProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListDCForProviderUnauthorized creates a ListDCForProviderUnauthorized with default headers values
func NewListDCForProviderUnauthorized() *ListDCForProviderUnauthorized {
	return &ListDCForProviderUnauthorized{}
}

/*ListDCForProviderUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type ListDCForProviderUnauthorized struct {
}

func (o *ListDCForProviderUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc][%d] listDCForProviderUnauthorized ", 401)
}

func (o *ListDCForProviderUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListDCForProviderForbidden creates a ListDCForProviderForbidden with default headers values
func NewListDCForProviderForbidden() *ListDCForProviderForbidden {
	return &ListDCForProviderForbidden{}
}

/*ListDCForProviderForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type ListDCForProviderForbidden struct {
}

func (o *ListDCForProviderForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc][%d] listDCForProviderForbidden ", 403)
}

func (o *ListDCForProviderForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListDCForProviderDefault creates a ListDCForProviderDefault with default headers values
func NewListDCForProviderDefault(code int) *ListDCForProviderDefault {
	return &ListDCForProviderDefault{
		_statusCode: code,
	}
}

/*ListDCForProviderDefault handles this case with default header values.

errorResponse
*/
type ListDCForProviderDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list d c for provider default response
func (o *ListDCForProviderDefault) Code() int {
	return o._statusCode
}

func (o *ListDCForProviderDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/{provider_name}/dc][%d] listDCForProvider default  %+v", o._statusCode, o.Payload)
}

func (o *ListDCForProviderDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListDCForProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
