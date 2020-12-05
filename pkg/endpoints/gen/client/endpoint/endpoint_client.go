///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new endpoint API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for endpoint API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddEndpoint adds a new endpoint
*/
func (a *Client) AddEndpoint(params *AddEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*AddEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddEndpointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addEndpoint",
		Method:             "POST",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddEndpointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddEndpointOK), nil

}

/*
DeleteEndpoint deletes an endpoint
*/
func (a *Client) DeleteEndpoint(params *DeleteEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEndpointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEndpoint",
		Method:             "DELETE",
		PathPattern:        "/{endpoint}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEndpointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEndpointOK), nil

}

/*
GetEndpoint finds endpoint by name

get an Endpoint by name
*/
func (a *Client) GetEndpoint(params *GetEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*GetEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEndpoint",
		Method:             "GET",
		PathPattern:        "/{endpoint}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEndpointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEndpointOK), nil

}

/*
GetEndpoints lists all existing endpoints
*/
func (a *Client) GetEndpoints(params *GetEndpointsParams, authInfo runtime.ClientAuthInfoWriter) (*GetEndpointsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEndpointsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEndpoints",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEndpointsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEndpointsOK), nil

}

/*
UpdateEndpoint updates an endpoint
*/
func (a *Client) UpdateEndpoint(params *UpdateEndpointParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateEndpointOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEndpointParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateEndpoint",
		Method:             "PUT",
		PathPattern:        "/{endpoint}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateEndpointReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateEndpointOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}