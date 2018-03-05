///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/vmware/dispatch/pkg/image-manager/gen/restapi/operations/base_image"
	"github.com/vmware/dispatch/pkg/image-manager/gen/restapi/operations/image"
)

// NewImageManagerAPI creates a new ImageManager instance
func NewImageManagerAPI(spec *loads.Document) *ImageManagerAPI {
	return &ImageManagerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		BaseImageAddBaseImageHandler: base_image.AddBaseImageHandlerFunc(func(params base_image.AddBaseImageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation BaseImageAddBaseImage has not yet been implemented")
		}),
		ImageAddImageHandler: image.AddImageHandlerFunc(func(params image.AddImageParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ImageAddImage has not yet been implemented")
		}),
		BaseImageDeleteBaseImageByNameHandler: base_image.DeleteBaseImageByNameHandlerFunc(func(params base_image.DeleteBaseImageByNameParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation BaseImageDeleteBaseImageByName has not yet been implemented")
		}),
		ImageDeleteImageByNameHandler: image.DeleteImageByNameHandlerFunc(func(params image.DeleteImageByNameParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ImageDeleteImageByName has not yet been implemented")
		}),
		BaseImageGetBaseImageByNameHandler: base_image.GetBaseImageByNameHandlerFunc(func(params base_image.GetBaseImageByNameParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation BaseImageGetBaseImageByName has not yet been implemented")
		}),
		BaseImageGetBaseImagesHandler: base_image.GetBaseImagesHandlerFunc(func(params base_image.GetBaseImagesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation BaseImageGetBaseImages has not yet been implemented")
		}),
		ImageGetImageByNameHandler: image.GetImageByNameHandlerFunc(func(params image.GetImageByNameParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ImageGetImageByName has not yet been implemented")
		}),
		ImageGetImagesHandler: image.GetImagesHandlerFunc(func(params image.GetImagesParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ImageGetImages has not yet been implemented")
		}),
		BaseImageUpdateBaseImageByNameHandler: base_image.UpdateBaseImageByNameHandlerFunc(func(params base_image.UpdateBaseImageByNameParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation BaseImageUpdateBaseImageByName has not yet been implemented")
		}),
		ImageUpdateImageByNameHandler: image.UpdateImageByNameHandlerFunc(func(params image.UpdateImageByNameParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation ImageUpdateImageByName has not yet been implemented")
		}),

		// Applies when the "Cookie" header is set
		CookieAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (cookie) Cookie from header param [Cookie] has not yet been implemented")
		},

		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*ImageManagerAPI VMware Dispatch Image Manager
 */
type ImageManagerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// CookieAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Cookie provided in the header
	CookieAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// BaseImageAddBaseImageHandler sets the operation handler for the add base image operation
	BaseImageAddBaseImageHandler base_image.AddBaseImageHandler
	// ImageAddImageHandler sets the operation handler for the add image operation
	ImageAddImageHandler image.AddImageHandler
	// BaseImageDeleteBaseImageByNameHandler sets the operation handler for the delete base image by name operation
	BaseImageDeleteBaseImageByNameHandler base_image.DeleteBaseImageByNameHandler
	// ImageDeleteImageByNameHandler sets the operation handler for the delete image by name operation
	ImageDeleteImageByNameHandler image.DeleteImageByNameHandler
	// BaseImageGetBaseImageByNameHandler sets the operation handler for the get base image by name operation
	BaseImageGetBaseImageByNameHandler base_image.GetBaseImageByNameHandler
	// BaseImageGetBaseImagesHandler sets the operation handler for the get base images operation
	BaseImageGetBaseImagesHandler base_image.GetBaseImagesHandler
	// ImageGetImageByNameHandler sets the operation handler for the get image by name operation
	ImageGetImageByNameHandler image.GetImageByNameHandler
	// ImageGetImagesHandler sets the operation handler for the get images operation
	ImageGetImagesHandler image.GetImagesHandler
	// BaseImageUpdateBaseImageByNameHandler sets the operation handler for the update base image by name operation
	BaseImageUpdateBaseImageByNameHandler base_image.UpdateBaseImageByNameHandler
	// ImageUpdateImageByNameHandler sets the operation handler for the update image by name operation
	ImageUpdateImageByNameHandler image.UpdateImageByNameHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ImageManagerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ImageManagerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ImageManagerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ImageManagerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ImageManagerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ImageManagerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ImageManagerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ImageManagerAPI
func (o *ImageManagerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CookieAuth == nil {
		unregistered = append(unregistered, "CookieAuth")
	}

	if o.BaseImageAddBaseImageHandler == nil {
		unregistered = append(unregistered, "base_image.AddBaseImageHandler")
	}

	if o.ImageAddImageHandler == nil {
		unregistered = append(unregistered, "image.AddImageHandler")
	}

	if o.BaseImageDeleteBaseImageByNameHandler == nil {
		unregistered = append(unregistered, "base_image.DeleteBaseImageByNameHandler")
	}

	if o.ImageDeleteImageByNameHandler == nil {
		unregistered = append(unregistered, "image.DeleteImageByNameHandler")
	}

	if o.BaseImageGetBaseImageByNameHandler == nil {
		unregistered = append(unregistered, "base_image.GetBaseImageByNameHandler")
	}

	if o.BaseImageGetBaseImagesHandler == nil {
		unregistered = append(unregistered, "base_image.GetBaseImagesHandler")
	}

	if o.ImageGetImageByNameHandler == nil {
		unregistered = append(unregistered, "image.GetImageByNameHandler")
	}

	if o.ImageGetImagesHandler == nil {
		unregistered = append(unregistered, "image.GetImagesHandler")
	}

	if o.BaseImageUpdateBaseImageByNameHandler == nil {
		unregistered = append(unregistered, "base_image.UpdateBaseImageByNameHandler")
	}

	if o.ImageUpdateImageByNameHandler == nil {
		unregistered = append(unregistered, "image.UpdateImageByNameHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ImageManagerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ImageManagerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	result := make(map[string]runtime.Authenticator)
	for name, scheme := range schemes {
		switch name {

		case "cookie":

			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.CookieAuth)

		}
	}
	return result

}

// Authorizer returns the registered authorizer
func (o *ImageManagerAPI) Authorizer() runtime.Authorizer {

	return o.APIAuthorizer

}

// ConsumersFor gets the consumers for the specified media types
func (o *ImageManagerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ImageManagerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ImageManagerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the image manager API
func (o *ImageManagerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ImageManagerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/baseimage"] = base_image.NewAddBaseImage(o.context, o.BaseImageAddBaseImageHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/image"] = image.NewAddImage(o.context, o.ImageAddImageHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/baseimage/{baseImageName}"] = base_image.NewDeleteBaseImageByName(o.context, o.BaseImageDeleteBaseImageByNameHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/image/{imageName}"] = image.NewDeleteImageByName(o.context, o.ImageDeleteImageByNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/baseimage/{baseImageName}"] = base_image.NewGetBaseImageByName(o.context, o.BaseImageGetBaseImageByNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/baseimage"] = base_image.NewGetBaseImages(o.context, o.BaseImageGetBaseImagesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/image/{imageName}"] = image.NewGetImageByName(o.context, o.ImageGetImageByNameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/image"] = image.NewGetImages(o.context, o.ImageGetImagesHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/baseimage/{baseImageName}"] = base_image.NewUpdateBaseImageByName(o.context, o.BaseImageUpdateBaseImageByNameHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/image/{imageName}"] = image.NewUpdateImageByName(o.context, o.ImageUpdateImageByNameHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ImageManagerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *ImageManagerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
