// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new service bindings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service bindings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ServiceBindingBinding generations of a service binding
*/
func (a *Client) ServiceBindingBinding(params *ServiceBindingBindingParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBindingBindingOK, *ServiceBindingBindingCreated, *ServiceBindingBindingAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingBindingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBinding.binding",
		Method:             "PUT",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingBindingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ServiceBindingBindingOK:
		return value, nil, nil, nil
	case *ServiceBindingBindingCreated:
		return nil, value, nil, nil
	case *ServiceBindingBindingAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
ServiceBindingGet gets a service binding
*/
func (a *Client) ServiceBindingGet(params *ServiceBindingGetParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBindingGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBinding.get",
		Method:             "GET",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBindingGetOK), nil

}

/*
ServiceBindingLastOperationGet lasts requested operation state for service binding
*/
func (a *Client) ServiceBindingLastOperationGet(params *ServiceBindingLastOperationGetParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBindingLastOperationGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingLastOperationGetParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBinding.lastOperation.get",
		Method:             "GET",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingLastOperationGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ServiceBindingLastOperationGetOK), nil

}

/*
ServiceBindingUnbinding deprovisions of a service binding
*/
func (a *Client) ServiceBindingUnbinding(params *ServiceBindingUnbindingParams, authInfo runtime.ClientAuthInfoWriter) (*ServiceBindingUnbindingOK, *ServiceBindingUnbindingAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingUnbindingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serviceBinding.unbinding",
		Method:             "DELETE",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBindingUnbindingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ServiceBindingUnbindingOK:
		return value, nil, nil
	case *ServiceBindingUnbindingAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
