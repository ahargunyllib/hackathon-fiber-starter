package errorhandler

import (
	"errors"

	"github.com/ahargunyllib/hackathon-fiber-starter/domain"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/helpers/http/response"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/log"
	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	var e validator.ValidationErrors
	if errors.As(err, &e) {
		log.Info(log.LogInfo{
			"error": e,
		}, "[ERROR_HANDLER] Validation error occurred")

		return response.SendResponse(c, fiber.StatusUnprocessableEntity, err)
	}

	switch e := err.(type) {
	case *domain.RequestError:
		response.SendResponse(c, e.StatusCode, e.Error())
	default:
		response.SendResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return nil
}
