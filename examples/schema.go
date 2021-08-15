package examples

import (
	"errors"
	graphql "github.com/graph-gophers/graphql-go"
)

var Schema = `
	schema {
	  query: Query
	  mutation: Mutation
	}

	type Query {
	  pendingTodo: [TodoItem]!
	  todo(id: ID!): TodoItem!
	}

	type Mutation {
	  markItemDone(id: ID!): TodoItem!
	}

	type TodoItem {
	  id: ID!
	  description: String!
	  type: Type!
	  isDone: Boolean!
	}

	enum Type {
	  PERSONAL
	  WORK
	  FAMILY
	}
`

type todoItem struct {
	ID graphql.ID
	Description string
	Type string
	IsDone bool
}

var todos = []*todoItem{
	{
		ID:     "0b6afc2b-f554-418d-b6b8-6c45e712177c",
		Description: "Prepare proposal for project A",
		Type:        "WORK",
		IsDone:      false,
	},
	{
		ID:     "5085f9be-f272-421b-8644-87b56d4df0d2",
		Description: "Pay for invoice 94456712",
		Type:        "WORK",
		IsDone:      false,
	},
	{
		ID:     "2acfe612-41e1-4a62-88f0-5a6abd932e79",
		Description: "Learn setting up GraphQL server with golang",
		Type:        "PERSONAL",
		IsDone:      false,
	},
	{
		ID:     "cce3f263-9c60-4dc9-88d5-ef052724b6ab",
		Description: "Buy toilet paper",
		Type:        "FAMILY",
		IsDone:      false,
	},
}

type todoItemResolver struct {
	t *todoItem
}

func (r *todoItemResolver) ID() graphql.ID {
	return r.t.ID
}

func (r *todoItemResolver) Description() string {
	return r.t.Description
}

func (r *todoItemResolver) Type() string {
	return r.t.Type
}

func (r *todoItemResolver) IsDone() bool {
	return r.t.IsDone
}

type Resolver struct {}

func (r *Resolver) PendingTodo() []*todoItemResolver {
	var result []*todoItemResolver

	for _, item := range todos {
		if !item.IsDone {
			result = append(result, &todoItemResolver{item})
		}
	}

	return result
}

func (r* Resolver) Todo(args struct { ID graphql.ID }) (*todoItemResolver, error) {
	for _, item := range todos {
		if item.ID == args.ID {
			return &todoItemResolver{item}, nil
		}
	}

	return nil, errors.New("not found")
}

func (r *Resolver) MarkItemDone(args struct { ID graphql.ID }) (*todoItemResolver, error) {
	for _, item := range todos {
		if item.ID == args.ID {
			item.IsDone = true
			return &todoItemResolver{item}, nil
		}
	}

	return nil, errors.New("not found")
}
