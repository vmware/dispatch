///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package image

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

// NewDeleteImageByNameParams creates a new DeleteImageByNameParams object
// with the default values initialized.
func NewDeleteImageByNameParams() *DeleteImageByNameParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &DeleteImageByNameParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteImageByNameParamsWithTimeout creates a new DeleteImageByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteImageByNameParamsWithTimeout(timeout time.Duration) *DeleteImageByNameParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &DeleteImageByNameParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: timeout,
	}
}

// NewDeleteImageByNameParamsWithContext creates a new DeleteImageByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteImageByNameParamsWithContext(ctx context.Context) *DeleteImageByNameParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &DeleteImageByNameParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		Context: ctx,
	}
}

// NewDeleteImageByNameParamsWithHTTPClient creates a new DeleteImageByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteImageByNameParamsWithHTTPClient(client *http.Client) *DeleteImageByNameParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &DeleteImageByNameParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,
		HTTPClient:       client,
	}
}

/*DeleteImageByNameParams contains all the parameters to send to the API endpoint
for the delete image by name operation typically these are written to a http.Request
*/
type DeleteImageByNameParams struct {

	/*XDispatchOrg*/
	XDispatchOrg *string
	/*XDispatchProject*/
	XDispatchProject *string
	/*ImageName
	  Name of image to return

	*/
	ImageName string
	/*Tags
	  Filter on image tags

	*/
	Tags []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete image by name params
func (o *DeleteImageByNameParams) WithTimeout(timeout time.Duration) *DeleteImageByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete image by name params
func (o *DeleteImageByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete image by name params
func (o *DeleteImageByNameParams) WithContext(ctx context.Context) *DeleteImageByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete image by name params
func (o *DeleteImageByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete image by name params
func (o *DeleteImageByNameParams) WithHTTPClient(client *http.Client) *DeleteImageByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete image by name params
func (o *DeleteImageByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the delete image by name params
func (o *DeleteImageByNameParams) WithXDispatchOrg(xDispatchOrg *string) *DeleteImageByNameParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the delete image by name params
func (o *DeleteImageByNameParams) SetXDispatchOrg(xDispatchOrg *string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithXDispatchProject adds the xDispatchProject to the delete image by name params
func (o *DeleteImageByNameParams) WithXDispatchProject(xDispatchProject *string) *DeleteImageByNameParams {
	o.SetXDispatchProject(xDispatchProject)
	return o
}

// SetXDispatchProject adds the xDispatchProject to the delete image by name params
func (o *DeleteImageByNameParams) SetXDispatchProject(xDispatchProject *string) {
	o.XDispatchProject = xDispatchProject
}

// WithImageName adds the imageName to the delete image by name params
func (o *DeleteImageByNameParams) WithImageName(imageName string) *DeleteImageByNameParams {
	o.SetImageName(imageName)
	return o
}

// SetImageName adds the imageName to the delete image by name params
func (o *DeleteImageByNameParams) SetImageName(imageName string) {
	o.ImageName = imageName
}

// WithTags adds the tags to the delete image by name params
func (o *DeleteImageByNameParams) WithTags(tags []string) *DeleteImageByNameParams {
	o.SetTags(tags)
	return o
}

// SetTags adds the tags to the delete image by name params
func (o *DeleteImageByNameParams) SetTags(tags []string) {
	o.Tags = tags
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteImageByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param imageName
	if err := r.SetPathParam("imageName", o.ImageName); err != nil {
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
