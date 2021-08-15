## Example app

A very simple example to use fiber-graphql-go to build a graphql server

to run this example, do:

```shell
go run ./examples/server.go
```

the server will be listening to port `9000`. Below is an example curl command to call the `pendingTodo` query

```shell
curl --request POST \
  --url http://localhost:9000/graphql \
  --header 'Content-Type: application/json' \
  --data '{"query":"{\n  pendingTodo {\n    id,\n    description\n    type\n    isDone\n  }\n}"}'
```
