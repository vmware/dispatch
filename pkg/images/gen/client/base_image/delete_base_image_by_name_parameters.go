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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteBaseImageByNameParams creates a new DeleteBaseImageByNameParams object
// with the default values initialized.
func NewDeleteBaseImageByNameParams() *DeleteBaseImageByNameParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &DeleteBaseImageByNameParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBaseImageByNameParamsWithTimeout creates a new DeleteBaseImageByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBaseImageByNameParamsWithTimeout(timeout time.Duration) *DeleteBaseImageByNameParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &DeleteBaseImageByNameParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: timeout,
	}
}

// NewDeleteBaseImageByNameParamsWithContext creates a new DeleteBaseImageByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBaseImageByNameParamsWithContext(ctx context.Context) *DeleteBaseImageByNameParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &DeleteBaseImageByNameParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		Context: ctx,
	}
}

// NewDeleteBaseImageByNameParamsWithHTTPClient creates a new DeleteBaseImageByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBaseImageByNameParamsWithHTTPClient(client *http.Client) *DeleteBaseImageByNameParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &DeleteBaseImageByNameParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,
		HTTPClient:       client,
	}
}

/*DeleteBaseImageByNameParams contains all the parameters to send to the API endpoint
for the delete base image by name operation typically these are written to a http.Request
*/
type DeleteBaseImageByNameParams struct {

	/*XDispatchOrg*/
	XDispatchOrg *string
	/*XDispatchProject*/
	XDispatchProject *string
	/*BaseImageName
	  Name of base image to return

	*/
	BaseImageName string
	/*Tags
	  Filter based on tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete base image by name params
func (o *DeleteBaseImageByNameParams) WithTimeout(timeout time.Duration) *DeleteBaseImageByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete base image by name params
func (o *DeleteBaseImageByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete base image by name params
func (o *DeleteBaseImageByNameParams) WithContext(ctx context.Context) *DeleteBaseImageByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete base image by name params
func (o *DeleteBaseImageByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete base image by name params
func (o *DeleteBaseImageByNameParams) WithHTTPClient(client *http.Client) *DeleteBaseImageByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete base image by name params
func (o *DeleteBaseImageByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the delete base image by name params
func (o *DeleteBaseImageByNameParams) WithXDispatchOrg(xDispatchOrg *string) *DeleteBaseImageByNameParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the delete base image by name params
func (o *DeleteBaseImageByNameParams) SetXDispatchOrg(xDispatchOrg *string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithXDispatchProject adds the xDispatchProject to the delete base image by name params
func (o *DeleteBaseImageByNameParams) WithXDispatchProject(xDispatchProject *string) *DeleteBaseImageByNameParams {
	o.SetXDispatchProject(xDispatchProject)
	return o
}

// SetXDispatchProject adds the xDispatchProject to the delete base image by name params
func (o *DeleteBaseImageByNameParams) SetXDispatchProject(xDispatchProject *string) {
	o.XDispatchProject = xDispatchProject
}

// WithBaseImageName adds the baseImageName to the delete base image by name params
func (o *DeleteBaseImageByNameParams) WithBaseImageName(baseImageName string) *DeleteBaseImageByNameParams {
	o.SetBaseImageName(baseImageName)
	return o
}

// SetBaseImageName adds the baseImageName to the delete base image by name params
func (o *DeleteBaseImageByNameParams) SetBaseImageName(baseImageName string) {
	o.BaseImageName = baseImageName
}

// WithTags adds the tags to the delete base image by name params
func (o *DeleteBaseImageByNameParams) WithTags(tags []string) *DeleteBaseImageByNameParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the delete base image by name params
func (o *DeleteBaseImageByNameParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBaseImageByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param baseImageName
	if err := r.SetPathParam("baseImageName", o.BaseImageName); err != nil {
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