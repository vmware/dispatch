///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package drivers

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

// NewGetDriverParams creates a new GetDriverParams object
// with the default values initialized.
func NewGetDriverParams() *GetDriverParams {
	var (
		xDispatchProjectDefault = string("default")
	)
	return &GetDriverParams{
		XDispatchProject: &xDispatchProjectDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDriverParamsWithTimeout creates a new GetDriverParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDriverParamsWithTimeout(timeout time.Duration) *GetDriverParams {
	var (
		xDispatchProjectDefault = string("default")
	)
	return &GetDriverParams{
		XDispatchProject: &xDispatchProjectDefault,

		timeout: timeout,
	}
}

// NewGetDriverParamsWithContext creates a new GetDriverParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDriverParamsWithContext(ctx context.Context) *GetDriverParams {
	var (
		xDispatchProjectDefault = string("default")
	)
	return &GetDriverParams{
		XDispatchProject: &xDispatchProjectDefault,

		Context: ctx,
	}
}

// NewGetDriverParamsWithHTTPClient creates a new GetDriverParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDriverParamsWithHTTPClient(client *http.Client) *GetDriverParams {
	var (
		xDispatchProjectDefault = string("default")
	)
	return &GetDriverParams{
		XDispatchProject: &xDispatchProjectDefault,
		HTTPClient:       client,
	}
}

/*GetDriverParams contains all the parameters to send to the API endpoint
for the get driver operation typically these are written to a http.Request
*/
type GetDriverParams struct {

	/*XDispatchOrg*/
	XDispatchOrg string
	/*XDispatchProject*/
	XDispatchProject *string
	/*DriverName
	  Name of the driver to work on

	*/
	DriverName string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get driver params
func (o *GetDriverParams) WithTimeout(timeout time.Duration) *GetDriverParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get driver params
func (o *GetDriverParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get driver params
func (o *GetDriverParams) WithContext(ctx context.Context) *GetDriverParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get driver params
func (o *GetDriverParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get driver params
func (o *GetDriverParams) WithHTTPClient(client *http.Client) *GetDriverParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get driver params
func (o *GetDriverParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the get driver params
func (o *GetDriverParams) WithXDispatchOrg(xDispatchOrg string) *GetDriverParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the get driver params
func (o *GetDriverParams) SetXDispatchOrg(xDispatchOrg string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithXDispatchProject adds the xDispatchProject to the get driver params
func (o *GetDriverParams) WithXDispatchProject(xDispatchProject *string) *GetDriverParams {
	o.SetXDispatchProject(xDispatchProject)
	return o
}

// SetXDispatchProject adds the xDispatchProject to the get driver params
func (o *GetDriverParams) SetXDispatchProject(xDispatchProject *string) {
	o.XDispatchProject = xDispatchProject
}

// WithDriverName adds the driverName to the get driver params
func (o *GetDriverParams) WithDriverName(driverName string) *GetDriverParams {
	o.SetDriverName(driverName)
	return o
}

// SetDriverName adds the driverName to the get driver params
func (o *GetDriverParams) SetDriverName(driverName string) {
	o.DriverName = driverName
}

// WithTags adds the tags to the get driver params
func (o *GetDriverParams) WithTags(tags []string) *GetDriverParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get driver params
func (o *GetDriverParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetDriverParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param driverName
	if err := r.SetPathParam("driverName", o.DriverName); err != nil {
		return err
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
