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

	"github.com/vmware/dispatch/pkg/api/v1"
)

// NewUpdateDriverTypeParams creates a new UpdateDriverTypeParams object
// with the default values initialized.
func NewUpdateDriverTypeParams() *UpdateDriverTypeParams {
	var ()
	return &UpdateDriverTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDriverTypeParamsWithTimeout creates a new UpdateDriverTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDriverTypeParamsWithTimeout(timeout time.Duration) *UpdateDriverTypeParams {
	var ()
	return &UpdateDriverTypeParams{

		timeout: timeout,
	}
}

// NewUpdateDriverTypeParamsWithContext creates a new UpdateDriverTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDriverTypeParamsWithContext(ctx context.Context) *UpdateDriverTypeParams {
	var ()
	return &UpdateDriverTypeParams{

		Context: ctx,
	}
}

// NewUpdateDriverTypeParamsWithHTTPClient creates a new UpdateDriverTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDriverTypeParamsWithHTTPClient(client *http.Client) *UpdateDriverTypeParams {
	var ()
	return &UpdateDriverTypeParams{
		HTTPClient: client,
	}
}

/*UpdateDriverTypeParams contains all the parameters to send to the API endpoint
for the update driver type operation typically these are written to a http.Request
*/
type UpdateDriverTypeParams struct {

	/*XDISPATCHORGID*/
	XDISPATCHORGID string
	/*Body
	  driver object

	*/
	Body *v1.EventDriverType
	/*DriverTypeName
	  Name of the driver type to work on

	*/
	DriverTypeName string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update driver type params
func (o *UpdateDriverTypeParams) WithTimeout(timeout time.Duration) *UpdateDriverTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update driver type params
func (o *UpdateDriverTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update driver type params
func (o *UpdateDriverTypeParams) WithContext(ctx context.Context) *UpdateDriverTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update driver type params
func (o *UpdateDriverTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update driver type params
func (o *UpdateDriverTypeParams) WithHTTPClient(client *http.Client) *UpdateDriverTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update driver type params
func (o *UpdateDriverTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDISPATCHORGID adds the xDISPATCHORGID to the update driver type params
func (o *UpdateDriverTypeParams) WithXDISPATCHORGID(xDISPATCHORGID string) *UpdateDriverTypeParams {
	o.SetXDISPATCHORGID(xDISPATCHORGID)
	return o
}

// SetXDISPATCHORGID adds the xDISPATCHORGId to the update driver type params
func (o *UpdateDriverTypeParams) SetXDISPATCHORGID(xDISPATCHORGID string) {
	o.XDISPATCHORGID = xDISPATCHORGID
}

// WithBody adds the body to the update driver type params
func (o *UpdateDriverTypeParams) WithBody(body *v1.EventDriverType) *UpdateDriverTypeParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update driver type params
func (o *UpdateDriverTypeParams) SetBody(body *v1.EventDriverType) {
	o.Body = body
}

// WithDriverTypeName adds the driverTypeName to the update driver type params
func (o *UpdateDriverTypeParams) WithDriverTypeName(driverTypeName string) *UpdateDriverTypeParams {
	o.SetDriverTypeName(driverTypeName)
	return o
}

// SetDriverTypeName adds the driverTypeName to the update driver type params
func (o *UpdateDriverTypeParams) SetDriverTypeName(driverTypeName string) {
	o.DriverTypeName = driverTypeName
}

// WithTags adds the tags to the update driver type params
func (o *UpdateDriverTypeParams) WithTags(tags []string) *UpdateDriverTypeParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the update driver type params
func (o *UpdateDriverTypeParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDriverTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-DISPATCH-ORG-ID
	if err := r.SetHeaderParam("X-DISPATCH-ORG-ID", o.XDISPATCHORGID); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param driverTypeName
	if err := r.SetPathParam("driverTypeName", o.DriverTypeName); err != nil {
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
