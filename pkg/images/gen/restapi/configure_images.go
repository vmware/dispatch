///////////////////////////////////////////////////////////////////////
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0
///////////////////////////////////////////////////////////////////////

// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/vmware/dispatch/pkg/images/gen/restapi/operations"
	"github.com/vmware/dispatch/pkg/images/gen/restapi/operations/base_image"
	"github.com/vmware/dispatch/pkg/images/gen/restapi/operations/image"
)

//go:generate swagger generate server --target ../pkg/images/gen --name Images --spec ../swagger/images.yaml --model-package v1 --skip-models --exclude-main

func configureFlags(api *operations.ImagesAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ImagesAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BaseImageAddBaseImageHandler = base_image.AddBaseImageHandlerFunc(func(params base_image.AddBaseImageParams) middleware.Responder {
		return middleware.NotImplemented("operation base_image.AddBaseImage has not yet been implemented")
	})
	api.ImageAddImageHandler = image.AddImageHandlerFunc(func(params image.AddImageParams) middleware.Responder {
		return middleware.NotImplemented("operation image.AddImage has not yet been implemented")
	})
	api.BaseImageDeleteBaseImageByNameHandler = base_image.DeleteBaseImageByNameHandlerFunc(func(params base_image.DeleteBaseImageByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation base_image.DeleteBaseImageByName has not yet been implemented")
	})
	api.ImageDeleteImageByNameHandler = image.DeleteImageByNameHandlerFunc(func(params image.DeleteImageByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation image.DeleteImageByName has not yet been implemented")
	})
	api.BaseImageGetBaseImageByNameHandler = base_image.GetBaseImageByNameHandlerFunc(func(params base_image.GetBaseImageByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation base_image.GetBaseImageByName has not yet been implemented")
	})
	api.BaseImageGetBaseImagesHandler = base_image.GetBaseImagesHandlerFunc(func(params base_image.GetBaseImagesParams) middleware.Responder {
		return middleware.NotImplemented("operation base_image.GetBaseImages has not yet been implemented")
	})
	api.ImageGetImageByNameHandler = image.GetImageByNameHandlerFunc(func(params image.GetImageByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation image.GetImageByName has not yet been implemented")
	})
	api.ImageGetImagesHandler = image.GetImagesHandlerFunc(func(params image.GetImagesParams) middleware.Responder {
		return middleware.NotImplemented("operation image.GetImages has not yet been implemented")
	})
	api.BaseImageUpdateBaseImageByNameHandler = base_image.UpdateBaseImageByNameHandlerFunc(func(params base_image.UpdateBaseImageByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation base_image.UpdateBaseImageByName has not yet been implemented")
	})
	api.ImageUpdateImageByNameHandler = image.UpdateImageByNameHandlerFunc(func(params image.UpdateImageByNameParams) middleware.Responder {
		return middleware.NotImplemented("operation image.UpdateImageByName has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
