package main

import (
	"fmt"
	fgg "github.com/cckwes/fiber-graphql-go"
	"github.com/cckwes/fiber-graphql-go/examples"
	"github.com/gofiber/fiber/v2"
	"github.com/graph-gophers/graphql-go"
)

func main()  {
	app := fiber.New()

	rawSchema, err := examples.ReadSchemaFile()
	if err != nil {
		fmt.Printf("error reading schema file %v", err.Error())
		return
	}

	schema := graphql.MustParseSchema(rawSchema, &examples.Resolver{})
	handler := fgg.Handler{Schema: schema}
	app.Post("/graphql", handler.ServeHTTP)

	err = app.Listen(":9000")
	if err != nil {
		fmt.Println("Failed to listen to port 9000, terminating...")
		fmt.Println(err.Error())
		return
	}
}
