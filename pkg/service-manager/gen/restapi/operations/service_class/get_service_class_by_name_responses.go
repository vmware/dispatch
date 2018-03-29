///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package service_class

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/dispatch/pkg/service-manager/gen/models"
)

// GetServiceClassByNameOKCode is the HTTP code returned for type GetServiceClassByNameOK
const GetServiceClassByNameOKCode int = 200

/*GetServiceClassByNameOK successful operation

swagger:response getServiceClassByNameOK
*/
type GetServiceClassByNameOK struct {

	/*
	  In: Body
	*/
	Payload *models.ServiceClass `json:"body,omitempty"`
}

// NewGetServiceClassByNameOK creates GetServiceClassByNameOK with default headers values
func NewGetServiceClassByNameOK() *GetServiceClassByNameOK {
	return &GetServiceClassByNameOK{}
}

// WithPayload adds the payload to the get service class by name o k response
func (o *GetServiceClassByNameOK) WithPayload(payload *models.ServiceClass) *GetServiceClassByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service class by name o k response
func (o *GetServiceClassByNameOK) SetPayload(payload *models.ServiceClass) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceClassByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServiceClassByNameBadRequestCode is the HTTP code returned for type GetServiceClassByNameBadRequest
const GetServiceClassByNameBadRequestCode int = 400

/*GetServiceClassByNameBadRequest Invalid name supplied

swagger:response getServiceClassByNameBadRequest
*/
type GetServiceClassByNameBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServiceClassByNameBadRequest creates GetServiceClassByNameBadRequest with default headers values
func NewGetServiceClassByNameBadRequest() *GetServiceClassByNameBadRequest {
	return &GetServiceClassByNameBadRequest{}
}

// WithPayload adds the payload to the get service class by name bad request response
func (o *GetServiceClassByNameBadRequest) WithPayload(payload *models.Error) *GetServiceClassByNameBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service class by name bad request response
func (o *GetServiceClassByNameBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceClassByNameBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetServiceClassByNameNotFoundCode is the HTTP code returned for type GetServiceClassByNameNotFound
const GetServiceClassByNameNotFoundCode int = 404

/*GetServiceClassByNameNotFound Service class not found

swagger:response getServiceClassByNameNotFound
*/
type GetServiceClassByNameNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServiceClassByNameNotFound creates GetServiceClassByNameNotFound with default headers values
func NewGetServiceClassByNameNotFound() *GetServiceClassByNameNotFound {
	return &GetServiceClassByNameNotFound{}
}

// WithPayload adds the payload to the get service class by name not found response
func (o *GetServiceClassByNameNotFound) WithPayload(payload *models.Error) *GetServiceClassByNameNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service class by name not found response
func (o *GetServiceClassByNameNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceClassByNameNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetServiceClassByNameDefault Generic error response

swagger:response getServiceClassByNameDefault
*/
type GetServiceClassByNameDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetServiceClassByNameDefault creates GetServiceClassByNameDefault with default headers values
func NewGetServiceClassByNameDefault(code int) *GetServiceClassByNameDefault {
	if code <= 0 {
		code = 500
	}

	return &GetServiceClassByNameDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get service class by name default response
func (o *GetServiceClassByNameDefault) WithStatusCode(code int) *GetServiceClassByNameDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get service class by name default response
func (o *GetServiceClassByNameDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get service class by name default response
func (o *GetServiceClassByNameDefault) WithPayload(payload *models.Error) *GetServiceClassByNameDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get service class by name default response
func (o *GetServiceClassByNameDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetServiceClassByNameDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
