///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/api/v1"
)

// DeletePolicyOKCode is the HTTP code returned for type DeletePolicyOK
const DeletePolicyOKCode int = 200

/*DeletePolicyOK Successful operation

swagger:response deletePolicyOK
*/
type DeletePolicyOK struct {

	/*
	  In: Body
	*/
	Payload *v1.Policy `json:"body,omitempty"`
}

// NewDeletePolicyOK creates DeletePolicyOK with default headers values
func NewDeletePolicyOK() *DeletePolicyOK {

	return &DeletePolicyOK{}
}

// WithPayload adds the payload to the delete policy o k response
func (o *DeletePolicyOK) WithPayload(payload *v1.Policy) *DeletePolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy o k response
func (o *DeletePolicyOK) SetPayload(payload *v1.Policy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePolicyBadRequestCode is the HTTP code returned for type DeletePolicyBadRequest
const DeletePolicyBadRequestCode int = 400

/*DeletePolicyBadRequest Invalid Name supplied

swagger:response deletePolicyBadRequest
*/
type DeletePolicyBadRequest struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeletePolicyBadRequest creates DeletePolicyBadRequest with default headers values
func NewDeletePolicyBadRequest() *DeletePolicyBadRequest {

	return &DeletePolicyBadRequest{}
}

// WithPayload adds the payload to the delete policy bad request response
func (o *DeletePolicyBadRequest) WithPayload(payload *v1.Error) *DeletePolicyBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy bad request response
func (o *DeletePolicyBadRequest) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePolicyNotFoundCode is the HTTP code returned for type DeletePolicyNotFound
const DeletePolicyNotFoundCode int = 404

/*DeletePolicyNotFound Policy not found

swagger:response deletePolicyNotFound
*/
type DeletePolicyNotFound struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeletePolicyNotFound creates DeletePolicyNotFound with default headers values
func NewDeletePolicyNotFound() *DeletePolicyNotFound {

	return &DeletePolicyNotFound{}
}

// WithPayload adds the payload to the delete policy not found response
func (o *DeletePolicyNotFound) WithPayload(payload *v1.Error) *DeletePolicyNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy not found response
func (o *DeletePolicyNotFound) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePolicyInternalServerErrorCode is the HTTP code returned for type DeletePolicyInternalServerError
const DeletePolicyInternalServerErrorCode int = 500

/*DeletePolicyInternalServerError Internal error

swagger:response deletePolicyInternalServerError
*/
type DeletePolicyInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *v1.Error `json:"body,omitempty"`
}

// NewDeletePolicyInternalServerError creates DeletePolicyInternalServerError with default headers values
func NewDeletePolicyInternalServerError() *DeletePolicyInternalServerError {

	return &DeletePolicyInternalServerError{}
}

// WithPayload adds the payload to the delete policy internal server error response
func (o *DeletePolicyInternalServerError) WithPayload(payload *v1.Error) *DeletePolicyInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete policy internal server error response
func (o *DeletePolicyInternalServerError) SetPayload(payload *v1.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePolicyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
