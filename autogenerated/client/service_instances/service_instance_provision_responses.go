// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/openservicebrokerapi/osb-checker/models"
)

// ServiceInstanceProvisionReader is a Reader for the ServiceInstanceProvision structure.
type ServiceInstanceProvisionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceInstanceProvisionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewServiceInstanceProvisionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 201:
		result := NewServiceInstanceProvisionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewServiceInstanceProvisionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewServiceInstanceProvisionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewServiceInstanceProvisionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewServiceInstanceProvisionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServiceInstanceProvisionOK creates a ServiceInstanceProvisionOK with default headers values
func NewServiceInstanceProvisionOK() *ServiceInstanceProvisionOK {
	return &ServiceInstanceProvisionOK{}
}

/*ServiceInstanceProvisionOK handles this case with default header values.

OK
*/
type ServiceInstanceProvisionOK struct {
	Payload *models.ServiceInstanceProvision
}

func (o *ServiceInstanceProvisionOK) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionOK  %+v", 200, o.Payload)
}

func (o *ServiceInstanceProvisionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstanceProvision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceProvisionCreated creates a ServiceInstanceProvisionCreated with default headers values
func NewServiceInstanceProvisionCreated() *ServiceInstanceProvisionCreated {
	return &ServiceInstanceProvisionCreated{}
}

/*ServiceInstanceProvisionCreated handles this case with default header values.

Created
*/
type ServiceInstanceProvisionCreated struct {
	Payload *models.ServiceInstanceProvision
}

func (o *ServiceInstanceProvisionCreated) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionCreated  %+v", 201, o.Payload)
}

func (o *ServiceInstanceProvisionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstanceProvision)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceProvisionAccepted creates a ServiceInstanceProvisionAccepted with default headers values
func NewServiceInstanceProvisionAccepted() *ServiceInstanceProvisionAccepted {
	return &ServiceInstanceProvisionAccepted{}
}

/*ServiceInstanceProvisionAccepted handles this case with default header values.

Accepted
*/
type ServiceInstanceProvisionAccepted struct {
	Payload *models.ServiceInstanceAsyncOperation
}

func (o *ServiceInstanceProvisionAccepted) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionAccepted  %+v", 202, o.Payload)
}

func (o *ServiceInstanceProvisionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstanceAsyncOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceProvisionBadRequest creates a ServiceInstanceProvisionBadRequest with default headers values
func NewServiceInstanceProvisionBadRequest() *ServiceInstanceProvisionBadRequest {
	return &ServiceInstanceProvisionBadRequest{}
}

/*ServiceInstanceProvisionBadRequest handles this case with default header values.

Bad Request
*/
type ServiceInstanceProvisionBadRequest struct {
	Payload *models.Error
}

func (o *ServiceInstanceProvisionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceInstanceProvisionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceProvisionConflict creates a ServiceInstanceProvisionConflict with default headers values
func NewServiceInstanceProvisionConflict() *ServiceInstanceProvisionConflict {
	return &ServiceInstanceProvisionConflict{}
}

/*ServiceInstanceProvisionConflict handles this case with default header values.

Conflict
*/
type ServiceInstanceProvisionConflict struct {
	Payload *models.Error
}

func (o *ServiceInstanceProvisionConflict) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionConflict  %+v", 409, o.Payload)
}

func (o *ServiceInstanceProvisionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceProvisionUnprocessableEntity creates a ServiceInstanceProvisionUnprocessableEntity with default headers values
func NewServiceInstanceProvisionUnprocessableEntity() *ServiceInstanceProvisionUnprocessableEntity {
	return &ServiceInstanceProvisionUnprocessableEntity{}
}

/*ServiceInstanceProvisionUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type ServiceInstanceProvisionUnprocessableEntity struct {
	Payload *models.Error
}

func (o *ServiceInstanceProvisionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v2/service_instances/{instance_id}][%d] serviceInstanceProvisionUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ServiceInstanceProvisionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
