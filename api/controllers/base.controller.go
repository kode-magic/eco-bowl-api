package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/kode-magic/eco-bowl-api/services"
)

func BaseController(app *fiber.App, appServices services.BaseService, session *session.Store) {
	api := app.Group("/api")
	UserRoute(api, appServices, session)
	InstituteRouter(api, appServices, session)
	EventRouter(api, appServices, session)
}
