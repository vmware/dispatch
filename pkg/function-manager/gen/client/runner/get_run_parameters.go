///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package runner

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

// NewGetRunParams creates a new GetRunParams object
// with the default values initialized.
func NewGetRunParams() *GetRunParams {
	var ()
	return &GetRunParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunParamsWithTimeout creates a new GetRunParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunParamsWithTimeout(timeout time.Duration) *GetRunParams {
	var ()
	return &GetRunParams{

		timeout: timeout,
	}
}

// NewGetRunParamsWithContext creates a new GetRunParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunParamsWithContext(ctx context.Context) *GetRunParams {
	var ()
	return &GetRunParams{

		Context: ctx,
	}
}

// NewGetRunParamsWithHTTPClient creates a new GetRunParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunParamsWithHTTPClient(client *http.Client) *GetRunParams {
	var ()
	return &GetRunParams{
		HTTPClient: client,
	}
}

/*GetRunParams contains all the parameters to send to the API endpoint
for the get run operation typically these are written to a http.Request
*/
type GetRunParams struct {

	/*XDISPATCHORGID*/
	XDISPATCHORGID string
	/*FunctionName
	  Name of function to retreive a run for

	*/
	FunctionName *string
	/*RunName
	  name of run to retrieve

	*/
	RunName strfmt.UUID
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get run params
func (o *GetRunParams) WithTimeout(timeout time.Duration) *GetRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get run params
func (o *GetRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get run params
func (o *GetRunParams) WithContext(ctx context.Context) *GetRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get run params
func (o *GetRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get run params
func (o *GetRunParams) WithHTTPClient(client *http.Client) *GetRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get run params
func (o *GetRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDISPATCHORGID adds the xDISPATCHORGID to the get run params
func (o *GetRunParams) WithXDISPATCHORGID(xDISPATCHORGID string) *GetRunParams {
	o.SetXDISPATCHORGID(xDISPATCHORGID)
	return o
}

// SetXDISPATCHORGID adds the xDISPATCHORGId to the get run params
func (o *GetRunParams) SetXDISPATCHORGID(xDISPATCHORGID string) {
	o.XDISPATCHORGID = xDISPATCHORGID
}

// WithFunctionName adds the functionName to the get run params
func (o *GetRunParams) WithFunctionName(functionName *string) *GetRunParams {
	o.SetFunctionName(functionName)
	return o
}

// SetFunctionName adds the functionName to the get run params
func (o *GetRunParams) SetFunctionName(functionName *string) {
	o.FunctionName = functionName
}

// WithRunName adds the runName to the get run params
func (o *GetRunParams) WithRunName(runName strfmt.UUID) *GetRunParams {
	o.SetRunName(runName)
	return o
}

// SetRunName adds the runName to the get run params
func (o *GetRunParams) SetRunName(runName strfmt.UUID) {
	o.RunName = runName
}

// WithTags adds the tags to the get run params
func (o *GetRunParams) WithTags(tags []string) *GetRunParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get run params
func (o *GetRunParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-DISPATCH-ORG-ID
	if err := r.SetHeaderParam("X-DISPATCH-ORG-ID", o.XDISPATCHORGID); err != nil {
		return err
	}

	if o.FunctionName != nil {

		// query param functionName
		var qrFunctionName string
		if o.FunctionName != nil {
			qrFunctionName = *o.FunctionName
		}
		qFunctionName := qrFunctionName
		if qFunctionName != "" {
			if err := r.SetQueryParam("functionName", qFunctionName); err != nil {
				return err
			}
		}

	}

	// path param runName
	if err := r.SetPathParam("runName", o.RunName.String()); err != nil {
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
