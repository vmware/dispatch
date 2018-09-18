///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package base_image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// NewAddBaseImageParams creates a new AddBaseImageParams object
// with the default values initialized.
func NewAddBaseImageParams() *AddBaseImageParams {
	var (
		xDispatchProjectDefault = string("default")
	)
	return &AddBaseImageParams{
		XDispatchProject: &xDispatchProjectDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewAddBaseImageParamsWithTimeout creates a new AddBaseImageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddBaseImageParamsWithTimeout(timeout time.Duration) *AddBaseImageParams {
	var (
		xDispatchProjectDefault = string("default")
	)
	return &AddBaseImageParams{
		XDispatchProject: &xDispatchProjectDefault,

		timeout: timeout,
	}
}

// NewAddBaseImageParamsWithContext creates a new AddBaseImageParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddBaseImageParamsWithContext(ctx context.Context) *AddBaseImageParams {
	var (
		xDispatchProjectDefault = string("default")
	)
	return &AddBaseImageParams{
		XDispatchProject: &xDispatchProjectDefault,

		Context: ctx,
	}
}

// NewAddBaseImageParamsWithHTTPClient creates a new AddBaseImageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddBaseImageParamsWithHTTPClient(client *http.Client) *AddBaseImageParams {
	var (
		xDispatchProjectDefault = string("default")
	)
	return &AddBaseImageParams{
		XDispatchProject: &xDispatchProjectDefault,
		HTTPClient:       client,
	}
}

/*AddBaseImageParams contains all the parameters to send to the API endpoint
for the add base image operation typically these are written to a http.Request
*/
type AddBaseImageParams struct {

	/*XDispatchOrg*/
	XDispatchOrg string
	/*XDispatchProject*/
	XDispatchProject *string
	/*Body
	  Base image object

	*/
	Body *v1.BaseImage

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add base image params
func (o *AddBaseImageParams) WithTimeout(timeout time.Duration) *AddBaseImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add base image params
func (o *AddBaseImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add base image params
func (o *AddBaseImageParams) WithContext(ctx context.Context) *AddBaseImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add base image params
func (o *AddBaseImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add base image params
func (o *AddBaseImageParams) WithHTTPClient(client *http.Client) *AddBaseImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add base image params
func (o *AddBaseImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the add base image params
func (o *AddBaseImageParams) WithXDispatchOrg(xDispatchOrg string) *AddBaseImageParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the add base image params
func (o *AddBaseImageParams) SetXDispatchOrg(xDispatchOrg string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithXDispatchProject adds the xDispatchProject to the add base image params
func (o *AddBaseImageParams) WithXDispatchProject(xDispatchProject *string) *AddBaseImageParams {
	o.SetXDispatchProject(xDispatchProject)
	return o
}

// SetXDispatchProject adds the xDispatchProject to the add base image params
func (o *AddBaseImageParams) SetXDispatchProject(xDispatchProject *string) {
	o.XDispatchProject = xDispatchProject
}

// WithBody adds the body to the add base image params
func (o *AddBaseImageParams) WithBody(body *v1.BaseImage) *AddBaseImageParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add base image params
func (o *AddBaseImageParams) SetBody(body *v1.BaseImage) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddBaseImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Dispatch-Org
	if err := r.SetHeaderParam("X-Dispatch-Org", o.XDispatchOrg); err != nil {
		return err
	}

	if o.XDispatchProject != nil {

		// header param X-Dispatch-Project
		if err := r.SetHeaderParam("X-Dispatch-Project", *o.XDispatchProject); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
