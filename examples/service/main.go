package main

import (
	"log"
	"os"

	loads "github.com/go-openapi/loads"
	"github.com/guzenok/go-memkvdb-example/examples/service/restapi"
	"github.com/guzenok/go-memkvdb-example/examples/service/restapi/operations"
	flags "github.com/jessevdk/go-flags"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewServiceAPI(swaggerSpec)

	server := restapi.NewServer(api)
	server.Port = 8080
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "key-value db"
	parser.LongDescription = "example"

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
