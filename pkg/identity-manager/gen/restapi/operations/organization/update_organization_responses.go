///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package organization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// UpdateOrganizationOKCode is the HTTP code returned for type UpdateOrganizationOK
const UpdateOrganizationOKCode int = 200

/*UpdateOrganizationOK Successful update

swagger:response updateOrganizationOK
*/
type UpdateOrganizationOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Organization `json:"body,omitempty"`
}

// NewUpdateOrganizationOK creates UpdateOrganizationOK with default headers values
func NewUpdateOrganizationOK() *UpdateOrganizationOK {

	return &UpdateOrganizationOK{}
}

// WithPayload adds the payload to the update organization o k response
func (o *UpdateOrganizationOK) WithPayload(payload *v1.Organization) *UpdateOrganizationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization o k response
func (o *UpdateOrganizationOK) SetPayload(payload *v1.Organization) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrganizationBadRequestCode is the HTTP code returned for type UpdateOrganizationBadRequest
const UpdateOrganizationBadRequestCode int = 400

/*UpdateOrganizationBadRequest Invalid input

swagger:response updateOrganizationBadRequest
*/
type UpdateOrganizationBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateOrganizationBadRequest creates UpdateOrganizationBadRequest with default headers values
func NewUpdateOrganizationBadRequest() *UpdateOrganizationBadRequest {

	return &UpdateOrganizationBadRequest{}
}

// WithPayload adds the payload to the update organization bad request response
func (o *UpdateOrganizationBadRequest) WithPayload(payload *v1.Error) *UpdateOrganizationBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization bad request response
func (o *UpdateOrganizationBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrganizationNotFoundCode is the HTTP code returned for type UpdateOrganizationNotFound
const UpdateOrganizationNotFoundCode int = 404

/*UpdateOrganizationNotFound Organization not found

swagger:response updateOrganizationNotFound
*/
type UpdateOrganizationNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateOrganizationNotFound creates UpdateOrganizationNotFound with default headers values
func NewUpdateOrganizationNotFound() *UpdateOrganizationNotFound {

	return &UpdateOrganizationNotFound{}
}

// WithPayload adds the payload to the update organization not found response
func (o *UpdateOrganizationNotFound) WithPayload(payload *v1.Error) *UpdateOrganizationNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization not found response
func (o *UpdateOrganizationNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateOrganizationInternalServerErrorCode is the HTTP code returned for type UpdateOrganizationInternalServerError
const UpdateOrganizationInternalServerErrorCode int = 500

/*UpdateOrganizationInternalServerError Internal error

swagger:response updateOrganizationInternalServerError
*/
type UpdateOrganizationInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewUpdateOrganizationInternalServerError creates UpdateOrganizationInternalServerError with default headers values
func NewUpdateOrganizationInternalServerError() *UpdateOrganizationInternalServerError {

	return &UpdateOrganizationInternalServerError{}
}

// WithPayload adds the payload to the update organization internal server error response
func (o *UpdateOrganizationInternalServerError) WithPayload(payload *v1.Error) *UpdateOrganizationInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update organization internal server error response
func (o *UpdateOrganizationInternalServerError) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateOrganizationInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}