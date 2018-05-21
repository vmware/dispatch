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

// NewGetApisParams creates a new GetApisParams object
// with the default values initialized.
func NewGetApisParams() *GetApisParams {
	var ()
	return &GetApisParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetApisParamsWithTimeout creates a new GetApisParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetApisParamsWithTimeout(timeout time.Duration) *GetApisParams {
	var ()
	return &GetApisParams{

		timeout: timeout,
	}
}

// NewGetApisParamsWithContext creates a new GetApisParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetApisParamsWithContext(ctx context.Context) *GetApisParams {
	var ()
	return &GetApisParams{

		Context: ctx,
	}
}

// NewGetApisParamsWithHTTPClient creates a new GetApisParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetApisParamsWithHTTPClient(client *http.Client) *GetApisParams {
	var ()
	return &GetApisParams{
		HTTPClient: client,
	}
}

/*GetApisParams contains all the parameters to send to the API endpoint
for the get apis operation typically these are written to a http.Request
*/
type GetApisParams struct {

	/*XDISPATCHORGID*/
	XDISPATCHORGID string
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

// WithTimeout adds the timeout to the get apis params
func (o *GetApisParams) WithTimeout(timeout time.Duration) *GetApisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get apis params
func (o *GetApisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get apis params
func (o *GetApisParams) WithContext(ctx context.Context) *GetApisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get apis params
func (o *GetApisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get apis params
func (o *GetApisParams) WithHTTPClient(client *http.Client) *GetApisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get apis params
func (o *GetApisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDISPATCHORGID adds the xDISPATCHORGID to the get apis params
func (o *GetApisParams) WithXDISPATCHORGID(xDISPATCHORGID string) *GetApisParams {
	o.SetXDISPATCHORGID(xDISPATCHORGID)
	return o
}

// SetXDISPATCHORGID adds the xDISPATCHORGId to the get apis params
func (o *GetApisParams) SetXDISPATCHORGID(xDISPATCHORGID string) {
	o.XDISPATCHORGID = xDISPATCHORGID
}

// WithFunction adds the function to the get apis params
func (o *GetApisParams) WithFunction(function *string) *GetApisParams {
	o.SetFunction(function)
	return o
}

// SetFunction adds the function to the get apis params
func (o *GetApisParams) SetFunction(function *string) {
	o.Function = function
}

// WithTags adds the tags to the get apis params
func (o *GetApisParams) WithTags(tags []string) *GetApisParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get apis params
func (o *GetApisParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetApisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-DISPATCH-ORG-ID
	if err := r.SetHeaderParam("X-DISPATCH-ORG-ID", o.XDISPATCHORGID); err != nil {
		return err
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
