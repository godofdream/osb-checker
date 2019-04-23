// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/openservicebrokerapi/osb-checker/autogenerated/models"
)

// ServiceBindingGetOKCode is the HTTP code returned for type ServiceBindingGetOK
const ServiceBindingGetOKCode int = 200

/*ServiceBindingGetOK OK

swagger:response serviceBindingGetOK
*/
type ServiceBindingGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.ServiceBindingResource `json:"body,omitempty"`
}

// NewServiceBindingGetOK creates ServiceBindingGetOK with default headers values
func NewServiceBindingGetOK() *ServiceBindingGetOK {

	return &ServiceBindingGetOK{}
}

// WithPayload adds the payload to the service binding get o k response
func (o *ServiceBindingGetOK) WithPayload(payload *models.ServiceBindingResource) *ServiceBindingGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding get o k response
func (o *ServiceBindingGetOK) SetPayload(payload *models.ServiceBindingResource) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingGetUnauthorizedCode is the HTTP code returned for type ServiceBindingGetUnauthorized
const ServiceBindingGetUnauthorizedCode int = 401

/*ServiceBindingGetUnauthorized Unauthorized

swagger:response serviceBindingGetUnauthorized
*/
type ServiceBindingGetUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingGetUnauthorized creates ServiceBindingGetUnauthorized with default headers values
func NewServiceBindingGetUnauthorized() *ServiceBindingGetUnauthorized {

	return &ServiceBindingGetUnauthorized{}
}

// WithPayload adds the payload to the service binding get unauthorized response
func (o *ServiceBindingGetUnauthorized) WithPayload(payload *models.Error) *ServiceBindingGetUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding get unauthorized response
func (o *ServiceBindingGetUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingGetNotFoundCode is the HTTP code returned for type ServiceBindingGetNotFound
const ServiceBindingGetNotFoundCode int = 404

/*ServiceBindingGetNotFound Not Found

swagger:response serviceBindingGetNotFound
*/
type ServiceBindingGetNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingGetNotFound creates ServiceBindingGetNotFound with default headers values
func NewServiceBindingGetNotFound() *ServiceBindingGetNotFound {

	return &ServiceBindingGetNotFound{}
}

// WithPayload adds the payload to the service binding get not found response
func (o *ServiceBindingGetNotFound) WithPayload(payload *models.Error) *ServiceBindingGetNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding get not found response
func (o *ServiceBindingGetNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ServiceBindingGetPreconditionFailedCode is the HTTP code returned for type ServiceBindingGetPreconditionFailed
const ServiceBindingGetPreconditionFailedCode int = 412

/*ServiceBindingGetPreconditionFailed Precondition Failed

swagger:response serviceBindingGetPreconditionFailed
*/
type ServiceBindingGetPreconditionFailed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewServiceBindingGetPreconditionFailed creates ServiceBindingGetPreconditionFailed with default headers values
func NewServiceBindingGetPreconditionFailed() *ServiceBindingGetPreconditionFailed {

	return &ServiceBindingGetPreconditionFailed{}
}

// WithPayload adds the payload to the service binding get precondition failed response
func (o *ServiceBindingGetPreconditionFailed) WithPayload(payload *models.Error) *ServiceBindingGetPreconditionFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the service binding get precondition failed response
func (o *ServiceBindingGetPreconditionFailed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ServiceBindingGetPreconditionFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(412)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
