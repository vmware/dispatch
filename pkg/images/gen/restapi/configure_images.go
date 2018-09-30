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

	"github.com/vmware/dispatch/pkg/images/gen/restapi/operations"
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

	// Applies when the "Authorization" header is set
	api.BearerAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (bearer) Authorization from header param [Authorization] has not yet been implemented")
	}
	// Applies when the "Cookie" header is set
	api.CookieAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (cookie) Cookie from header param [Cookie] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.ImageAddImageHandler = image.AddImageHandlerFunc(func(params image.AddImageParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation image.AddImage has not yet been implemented")
	})
	api.ImageDeleteImageByNameHandler = image.DeleteImageByNameHandlerFunc(func(params image.DeleteImageByNameParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation image.DeleteImageByName has not yet been implemented")
	})
	api.ImageGetImageByNameHandler = image.GetImageByNameHandlerFunc(func(params image.GetImageByNameParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation image.GetImageByName has not yet been implemented")
	})
	api.ImageGetImagesHandler = image.GetImagesHandlerFunc(func(params image.GetImagesParams, principal interface{}) middleware.Responder {
		return middleware.NotImplemented("operation image.GetImages has not yet been implemented")
	})
	api.ImageUpdateImageByNameHandler = image.UpdateImageByNameHandlerFunc(func(params image.UpdateImageByNameParams, principal interface{}) middleware.Responder {
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
func configureServer(s *http.Server, scheme, addr string) {
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
