package tests

import (
	"github.com/gofiber/fiber/v2/utils"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	fgg "github.com/cckwes/fiber-graphql-go"
	fiber "github.com/gofiber/fiber/v2"
)

func TestGraphQLQuery(t *testing.T) {
	app := fiber.New()
	h := fgg.Handler{Schema: GetSchema()}
	app.Post("/graphql", h.ServeHTTP)

	req := httptest.NewRequest(
		"POST",
		"/graphql",
		strings.NewReader(`
						{
						  "query": "{ hello }"
						}
					`),
	)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, `{"data":{"hello":"Hello, this is fiber-graphql-go"}}`, string(body))
}

func TestGraphQLMutation(t *testing.T) {
	app := fiber.New()
	h := fgg.Handler{Schema: GetSchema()}
	app.Post("/graphql", h.ServeHTTP)

	req := httptest.NewRequest(
		"POST",
		"/graphql",
		strings.NewReader(`
						{
						  "query": "mutation SetHello($data: String!) { setHelloString(data: $data) }",
						  "variables": { "data": "Updated hello" }
						}
					`),
	)
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, 200, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, `{"data":{"setHelloString":"Updated hello"}}`, string(body))
}