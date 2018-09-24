///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package endpoint

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

// NewGetEndpointsParams creates a new GetEndpointsParams object
// with the default values initialized.
func NewGetEndpointsParams() *GetEndpointsParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &GetEndpointsParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEndpointsParamsWithTimeout creates a new GetEndpointsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEndpointsParamsWithTimeout(timeout time.Duration) *GetEndpointsParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &GetEndpointsParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: timeout,
	}
}

// NewGetEndpointsParamsWithContext creates a new GetEndpointsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEndpointsParamsWithContext(ctx context.Context) *GetEndpointsParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &GetEndpointsParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		Context: ctx,
	}
}

// NewGetEndpointsParamsWithHTTPClient creates a new GetEndpointsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEndpointsParamsWithHTTPClient(client *http.Client) *GetEndpointsParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &GetEndpointsParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,
		HTTPClient:       client,
	}
}

/*GetEndpointsParams contains all the parameters to send to the API endpoint
for the get endpoints operation typically these are written to a http.Request
*/
type GetEndpointsParams struct {

	/*XDispatchOrg*/
	XDispatchOrg *string
	/*XDispatchProject*/
	XDispatchProject *string
	/*Function
	  Filter based on function names

	*/
	Function *string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get endpoints params
func (o *GetEndpointsParams) WithTimeout(timeout time.Duration) *GetEndpointsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get endpoints params
func (o *GetEndpointsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get endpoints params
func (o *GetEndpointsParams) WithContext(ctx context.Context) *GetEndpointsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get endpoints params
func (o *GetEndpointsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get endpoints params
func (o *GetEndpointsParams) WithHTTPClient(client *http.Client) *GetEndpointsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get endpoints params
func (o *GetEndpointsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the get endpoints params
func (o *GetEndpointsParams) WithXDispatchOrg(xDispatchOrg *string) *GetEndpointsParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the get endpoints params
func (o *GetEndpointsParams) SetXDispatchOrg(xDispatchOrg *string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithXDispatchProject adds the xDispatchProject to the get endpoints params
func (o *GetEndpointsParams) WithXDispatchProject(xDispatchProject *string) *GetEndpointsParams {
	o.SetXDispatchProject(xDispatchProject)
	return o
}

// SetXDispatchProject adds the xDispatchProject to the get endpoints params
func (o *GetEndpointsParams) SetXDispatchProject(xDispatchProject *string) {
	o.XDispatchProject = xDispatchProject
}

// WithFunction adds the function to the get endpoints params
func (o *GetEndpointsParams) WithFunction(function *string) *GetEndpointsParams {
	o.SetFunction(function)
	return o
}

// SetFunction adds the function to the get endpoints params
func (o *GetEndpointsParams) SetFunction(function *string) {
	o.Function = function
}

// WithTags adds the tags to the get endpoints params
func (o *GetEndpointsParams) WithTags(tags []string) *GetEndpointsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get endpoints params
func (o *GetEndpointsParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetEndpointsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XDispatchOrg != nil {

		// header param X-Dispatch-Org
		if err := r.SetHeaderParam("X-Dispatch-Org", *o.XDispatchOrg); err != nil {
			return err
		}

	}

	if o.XDispatchProject != nil {

		// header param X-Dispatch-Project
		if err := r.SetHeaderParam("X-Dispatch-Project", *o.XDispatchProject); err != nil {
			return err
		}

	}

	if o.Function != nil {

		// query param function
		var qrFunction string
		if o.Function != nil {
			qrFunction = *o.Function
		}
		qFunction := qrFunction
		if qFunction != "" {
			if err := r.SetQueryParam("function", qFunction); err != nil {
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
