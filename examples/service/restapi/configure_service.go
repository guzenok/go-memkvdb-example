// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	errors "github.com/go-openapi/errors"
	//	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/guzenok/go-memkvdb-example/examples/service/restapi/operations"

	db "github.com/guzenok/go-memkvdb-example"
)

var storage *db.DB

func init() {
	var err error
	storage, err = db.NewDefault(30 * time.Second)
	if err != nil {
		log.Fatalln(err)
	}
}

//go:generate swagger generate server --target ../examples/service/gen --name service --spec ../examples/service/swagger.yml --exclude-main

func configureFlags(api *operations.ServiceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{...}
}

func configureAPI(api *operations.ServiceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	//api.JSONConsumer = runtime.JSONConsumer()

	//api.JSONProducer = runtime.JSONProducer()

	api.GetValuesKeyHandler = operations.GetValuesKeyHandlerFunc(func(params operations.GetValuesKeyParams) middleware.Responder {
		val, err := storage.Get([]byte(params.Key))
		if err == db.ErrNotFound {
			return operations.NewGetValuesKeyNotFound()
		}
		if err != nil {
			log.Println(err)
			return operations.NewGetValuesKeyDefault(500)
		}

		res := ioutil.NopCloser(bytes.NewReader(val))
		return operations.NewGetValuesKeyOK().WithPayload(res)
	})

	api.PostValuesKeyHandler = operations.PostValuesKeyHandlerFunc(func(params operations.PostValuesKeyParams) middleware.Responder {
		val, err := ioutil.ReadAll(params.Value)
		if err != nil {
			log.Println(err)
			return operations.NewGetValuesKeyDefault(400)
		}

		err = storage.Set([]byte(params.Key), val)
		if err != nil {
			log.Println(err)
			return operations.NewPostValuesKeyDefault(500)
		}
		return operations.NewPostValuesKeyOK()
	})

	// api.ServerShutdown = func() {}

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
	return
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
