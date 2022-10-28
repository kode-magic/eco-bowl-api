package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kode-magic/eco-bowl-api/api/controllers"

	//controllers "github.com/Soft-Magick/birth-death-api/api/controller"
	"github.com/kode-magic/eco-bowl-api/services"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store *session.Store
)

func App(services services.BaseService) {
	store = session.New(session.Config{
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieSameSite: "strict",
		Expiration:     time.Hour * 5,
	})

	app := fiber.New()

	CorsMiddleware(app)
	AuthMiddleware(app, store)

	controllers.BaseController(app, services, store)

	appPort := os.Getenv("APP_PORT")

	err := app.Listen(appPort)

	if err != nil {
		panic(err)
	}
}
