package tests

import (
	"errors"
	"github.com/graph-gophers/graphql-go"
)

var s = `
      schema {
        query: Query
        mutation: Mutation
      }
      type Query {
        hello: String!
		showError: String!
      }
      type Mutation {
       setHelloString(data: String!): String!
      }
`

type resolver struct {}

func (r *resolver) Hello() string {
	return "Hello, this is fiber-graphql-go"
}

func (r *resolver) ShowError() (string, error) {
	return "", errors.New("error in show error")
}

func (r *resolver) SetHelloString(args *struct {
	Data string
}) string {
	return args.Data
}

func GetSchema() *graphql.Schema {
	schema := graphql.MustParseSchema(s, &resolver{}, graphql.UseStringDescriptions())

	return schema
}
