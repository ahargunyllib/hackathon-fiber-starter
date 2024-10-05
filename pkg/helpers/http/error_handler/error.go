package errorhandler

import (
	"errors"

	"github.com/ahargunyllib/hackathon-fiber-starter/domain"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/helpers/http/response"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var valErr *validator.ValidationErrorsResponse
	if errors.As(err, &valErr) {
		response.SendResponse(c, fiber.StatusUnprocessableEntity, valErr.Serialize())
	}

	switch e := err.(type) {
	case *domain.RequestError:
		response.SendResponse(c, e.StatusCode, e.Error())
	default:
		response.SendResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
