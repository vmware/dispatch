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

// NewGetRunsParams creates a new GetRunsParams object
// with the default values initialized.
func NewGetRunsParams() *GetRunsParams {
	var ()
	return &GetRunsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRunsParamsWithTimeout creates a new GetRunsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRunsParamsWithTimeout(timeout time.Duration) *GetRunsParams {
	var ()
	return &GetRunsParams{

		timeout: timeout,
	}
}

// NewGetRunsParamsWithContext creates a new GetRunsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRunsParamsWithContext(ctx context.Context) *GetRunsParams {
	var ()
	return &GetRunsParams{

		Context: ctx,
	}
}

// NewGetRunsParamsWithHTTPClient creates a new GetRunsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRunsParamsWithHTTPClient(client *http.Client) *GetRunsParams {
	var ()
	return &GetRunsParams{
		HTTPClient: client,
	}
}

/*GetRunsParams contains all the parameters to send to the API endpoint
for the get runs operation typically these are written to a http.Request
*/
type GetRunsParams struct {

	/*XDISPATCHORGID*/
	XDISPATCHORGID string
	/*FunctionName
	  Name of function to run or retreive runs for

	*/
	FunctionName *string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get runs params
func (o *GetRunsParams) WithTimeout(timeout time.Duration) *GetRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get runs params
func (o *GetRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get runs params
func (o *GetRunsParams) WithContext(ctx context.Context) *GetRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get runs params
func (o *GetRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get runs params
func (o *GetRunsParams) WithHTTPClient(client *http.Client) *GetRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get runs params
func (o *GetRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDISPATCHORGID adds the xDISPATCHORGID to the get runs params
func (o *GetRunsParams) WithXDISPATCHORGID(xDISPATCHORGID string) *GetRunsParams {
	o.SetXDISPATCHORGID(xDISPATCHORGID)
	return o
}

// SetXDISPATCHORGID adds the xDISPATCHORGId to the get runs params
func (o *GetRunsParams) SetXDISPATCHORGID(xDISPATCHORGID string) {
	o.XDISPATCHORGID = xDISPATCHORGID
}

// WithFunctionName adds the functionName to the get runs params
func (o *GetRunsParams) WithFunctionName(functionName *string) *GetRunsParams {
	o.SetFunctionName(functionName)
	return o
}

// SetFunctionName adds the functionName to the get runs params
func (o *GetRunsParams) SetFunctionName(functionName *string) {
	o.FunctionName = functionName
}

// WithTags adds the tags to the get runs params
func (o *GetRunsParams) WithTags(tags []string) *GetRunsParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the get runs params
func (o *GetRunsParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *GetRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
