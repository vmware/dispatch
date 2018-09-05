///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package secret

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

// NewUpdateSecretParams creates a new UpdateSecretParams object
// with the default values initialized.
func NewUpdateSecretParams() *UpdateSecretParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &UpdateSecretParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSecretParamsWithTimeout creates a new UpdateSecretParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateSecretParamsWithTimeout(timeout time.Duration) *UpdateSecretParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &UpdateSecretParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		timeout: timeout,
	}
}

// NewUpdateSecretParamsWithContext creates a new UpdateSecretParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateSecretParamsWithContext(ctx context.Context) *UpdateSecretParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &UpdateSecretParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,

		Context: ctx,
	}
}

// NewUpdateSecretParamsWithHTTPClient creates a new UpdateSecretParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateSecretParamsWithHTTPClient(client *http.Client) *UpdateSecretParams {
	var (
		xDispatchOrgDefault     = string("default")
		xDispatchProjectDefault = string("default")
	)
	return &UpdateSecretParams{
		XDispatchOrg:     &xDispatchOrgDefault,
		XDispatchProject: &xDispatchProjectDefault,
		HTTPClient:       client,
	}
}

/*UpdateSecretParams contains all the parameters to send to the API endpoint
for the update secret operation typically these are written to a http.Request
*/
type UpdateSecretParams struct {

	/*XDispatchOrg*/
	XDispatchOrg *string
	/*XDispatchProject*/
	XDispatchProject *string
	/*Secret*/
	Secret *v1.Secret
	/*SecretName*/
	SecretName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update secret params
func (o *UpdateSecretParams) WithTimeout(timeout time.Duration) *UpdateSecretParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update secret params
func (o *UpdateSecretParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update secret params
func (o *UpdateSecretParams) WithContext(ctx context.Context) *UpdateSecretParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update secret params
func (o *UpdateSecretParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update secret params
func (o *UpdateSecretParams) WithHTTPClient(client *http.Client) *UpdateSecretParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update secret params
func (o *UpdateSecretParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDispatchOrg adds the xDispatchOrg to the update secret params
func (o *UpdateSecretParams) WithXDispatchOrg(xDispatchOrg *string) *UpdateSecretParams {
	o.SetXDispatchOrg(xDispatchOrg)
	return o
}

// SetXDispatchOrg adds the xDispatchOrg to the update secret params
func (o *UpdateSecretParams) SetXDispatchOrg(xDispatchOrg *string) {
	o.XDispatchOrg = xDispatchOrg
}

// WithXDispatchProject adds the xDispatchProject to the update secret params
func (o *UpdateSecretParams) WithXDispatchProject(xDispatchProject *string) *UpdateSecretParams {
	o.SetXDispatchProject(xDispatchProject)
	return o
}

// SetXDispatchProject adds the xDispatchProject to the update secret params
func (o *UpdateSecretParams) SetXDispatchProject(xDispatchProject *string) {
	o.XDispatchProject = xDispatchProject
}

// WithSecret adds the secret to the update secret params
func (o *UpdateSecretParams) WithSecret(secret *v1.Secret) *UpdateSecretParams {
	o.SetSecret(secret)
	return o
}

// SetSecret adds the secret to the update secret params
func (o *UpdateSecretParams) SetSecret(secret *v1.Secret) {
	o.Secret = secret
}

// WithSecretName adds the secretName to the update secret params
func (o *UpdateSecretParams) WithSecretName(secretName string) *UpdateSecretParams {
	o.SetSecretName(secretName)
	return o
}

// SetSecretName adds the secretName to the update secret params
func (o *UpdateSecretParams) SetSecretName(secretName string) {
	o.SecretName = secretName
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSecretParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Secret != nil {
		if err := r.SetBodyParam(o.Secret); err != nil {
			return err
		}
	}

	// path param secretName
	if err := r.SetPathParam("secretName", o.SecretName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
