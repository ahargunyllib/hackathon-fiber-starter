package middlewares

import (
	"context"
	"time"

	"github.com/ahargunyllib/hackathon-fiber-starter/pkg/log"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
)

type contextKey string

const correlationIDKey contextKey = "correlation_id"

func LoggerConfig() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		start := time.Now()

		correlationID := xid.New().String()

		c := context.WithValue(ctx.Context(), correlationIDKey, correlationID)

		ctx.SetUserContext(c)

		log.UpdateContext(string(correlationIDKey), correlationID)

		ctx.Request().Header.Add(string(correlationIDKey), correlationID)

		defer func() {
			log.Info(
				log.LogInfo{
					"method":     ctx.Method(),
					"path":       ctx.Path(),
					"status":     ctx.Response().StatusCode(),
					"user_agent": ctx.Get("User-Agent"),
					"latency":    time.Since(start).String(),
				},
				"[LOGGER MIDDLEWARE]",
			)
		}()

		return ctx.Next()
	}
}
