// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/openservicebrokerapi/osb-checker/models"
)

// ServiceBindingUnbindingReader is a Reader for the ServiceBindingUnbinding structure.
type ServiceBindingUnbindingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBindingUnbindingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceBindingUnbindingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewServiceBindingUnbindingAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServiceBindingUnbindingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 410:
		result := NewServiceBindingUnbindingGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceBindingUnbindingOK creates a ServiceBindingUnbindingOK with default headers values
func NewServiceBindingUnbindingOK() *ServiceBindingUnbindingOK {
	return &ServiceBindingUnbindingOK{}
}

/*ServiceBindingUnbindingOK handles this case with default header values.

OK
*/
type ServiceBindingUnbindingOK struct {
	Payload models.Object
}

func (o *ServiceBindingUnbindingOK) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingOK  %+v", 200, o.Payload)
}

func (o *ServiceBindingUnbindingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingUnbindingAccepted creates a ServiceBindingUnbindingAccepted with default headers values
func NewServiceBindingUnbindingAccepted() *ServiceBindingUnbindingAccepted {
	return &ServiceBindingUnbindingAccepted{}
}

/*ServiceBindingUnbindingAccepted handles this case with default header values.

Accepted
*/
type ServiceBindingUnbindingAccepted struct {
	Payload *models.AsyncOperation
}

func (o *ServiceBindingUnbindingAccepted) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingAccepted  %+v", 202, o.Payload)
}

func (o *ServiceBindingUnbindingAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AsyncOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingUnbindingBadRequest creates a ServiceBindingUnbindingBadRequest with default headers values
func NewServiceBindingUnbindingBadRequest() *ServiceBindingUnbindingBadRequest {
	return &ServiceBindingUnbindingBadRequest{}
}

/*ServiceBindingUnbindingBadRequest handles this case with default header values.

Bad Request
*/
type ServiceBindingUnbindingBadRequest struct {
	Payload *models.Error
}

func (o *ServiceBindingUnbindingBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBindingUnbindingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingUnbindingGone creates a ServiceBindingUnbindingGone with default headers values
func NewServiceBindingUnbindingGone() *ServiceBindingUnbindingGone {
	return &ServiceBindingUnbindingGone{}
}

/*ServiceBindingUnbindingGone handles this case with default header values.

Gone
*/
type ServiceBindingUnbindingGone struct {
	Payload *models.Error
}

func (o *ServiceBindingUnbindingGone) Error() string {
	return fmt.Sprintf("[DELETE /v2/service_instances/{instance_id}/service_bindings/{binding_id}][%d] serviceBindingUnbindingGone  %+v", 410, o.Payload)
}

func (o *ServiceBindingUnbindingGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
