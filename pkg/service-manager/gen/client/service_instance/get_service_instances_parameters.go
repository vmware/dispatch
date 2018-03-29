///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package service_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetServiceInstancesParams creates a new GetServiceInstancesParams object
// with the default values initialized.
func NewGetServiceInstancesParams() *GetServiceInstancesParams {
	var ()
	return &GetServiceInstancesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceInstancesParamsWithTimeout creates a new GetServiceInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceInstancesParamsWithTimeout(timeout time.Duration) *GetServiceInstancesParams {
	var ()
	return &GetServiceInstancesParams{

		timeout: timeout,
	}
}

// NewGetServiceInstancesParamsWithContext creates a new GetServiceInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceInstancesParamsWithContext(ctx context.Context) *GetServiceInstancesParams {
	var ()
	return &GetServiceInstancesParams{

		Context: ctx,
	}
}

// NewGetServiceInstancesParamsWithHTTPClient creates a new GetServiceInstancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceInstancesParamsWithHTTPClient(client *http.Client) *GetServiceInstancesParams {
	var ()
	return &GetServiceInstancesParams{
		HTTPClient: client,
	}
}

/*GetServiceInstancesParams contains all the parameters to send to the API endpoint
for the get service instances operation typically these are written to a http.Request
*/
type GetServiceInstancesParams struct {

	/*Serviceclass
	  service class name

	*/
	Serviceclass *string
	/*Tags
	  Filter on service instance tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service instances params
func (o *GetServiceInstancesParams) WithTimeout(timeout time.Duration) *GetServiceInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service instances params
func (o *GetServiceInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service instances params
func (o *GetServiceInstancesParams) WithContext(ctx context.Context) *GetServiceInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service instances params
func (o *GetServiceInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service instances params
func (o *GetServiceInstancesParams) WithHTTPClient(client *http.Client) *GetServiceInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service instances params
func (o *GetServiceInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithServiceclass adds the serviceclass to the get service instances params
func (o *GetServiceInstancesParams) WithServiceclass(serviceclass *string) *GetServiceInstancesParams {
	o.SetServiceclass(serviceclass)
	return o
}

// SetServiceclass adds the serviceclass to the get service instances params
func (o *GetServiceInstancesParams) SetServiceclass(serviceclass *string) {
	o.Serviceclass = serviceclass
}

// WithTags adds the tags to the get service instances params
func (o *GetServiceInstancesParams) WithTags(tags []string) *GetServiceInstancesParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get service instances params
func (o *GetServiceInstancesParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Serviceclass != nil {

		// query param serviceclass
		var qrServiceclass string
		if o.Serviceclass != nil {
			qrServiceclass = *o.Serviceclass
		}
		qServiceclass := qrServiceclass
		if qServiceclass != "" {
			if err := r.SetQueryParam("serviceclass", qServiceclass); err != nil {
				return err
			}
		}

	}

	valuesTags := o.Tags

	joinedTags := swag.JoinByFormat(valuesTags, "multi")
	// query array param tags
	if err := r.SetQueryParam("tags", joinedTags...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
