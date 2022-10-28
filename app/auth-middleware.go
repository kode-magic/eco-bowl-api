package app

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/kode-magic/eco-bowl-api/utils"
	"strings"
)

func AuthMiddleware(app *fiber.App, store *session.Store) {
	app.Use(func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")

		if strings.Contains(ctx.Path(), "auth") {
			return ctx.Next()
		}

		sess, err := store.Get(ctx)

		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, errors.New("not authorized").Error())
		}

		if sess.Get(utils.KEY) == nil {
			return fiber.NewError(fiber.StatusUnauthorized, errors.New("not authorized").Error())
		}

		return ctx.Next()
	})
}
