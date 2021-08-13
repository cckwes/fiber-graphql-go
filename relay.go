package relay

import (
	"github.com/gofiber/fiber/v2"
	"github.com/graph-gophers/graphql-go"
)

type Handler struct {
	Schema *graphql.Schema
}

func (h *Handler) ServeHTTP(c *fiber.Ctx) error {
	type Params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	body := new(Params)

	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	response := h.Schema.Exec(c.Context(), body.Query, body.OperationName, body.Variables)
	err := c.JSON(response)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
