package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	service "github.com/kode-magic/eco-bowl-api/services"
)

type Team struct {
	Service service.TeamService
	Session *session.Store
}

func TeamConstructor(appInterface service.TeamService, session *session.Store) *Team {
	return &Team{
		Service: appInterface,
		Session: session,
	}
}

func (t *Team) Create(ctx *fiber.Ctx) error {
	eventId := ctx.Params("event_id", "")
	var requestBody service.TeamRequest

	requestBodyErr := ctx.BodyParser(&requestBody)

	if requestBodyErr != nil {
		err := ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Request error": "Cannot parse json, invalid request body",
		})

		if err != nil {
			return err
		}

		return nil
	}

	requestBody.Event = eventId

	team, addErr := t.Service.Create(&requestBody)

	if addErr != nil {
		ctxErr := ctx.Status(fiber.StatusInternalServerError).JSON(addErr)

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Team " + requestBody.Name + " has been created successfully. We wish you all the best",
		"data":    team,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (t *Team) List(ctx *fiber.Ctx) error {
	eventId := ctx.Params("event_id", "")
	res, err := t.Service.List(eventId)

	if err != nil {
		ctxErr := ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    res,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func TeamRouter(router fiber.Router, services service.BaseService, session *session.Store) {
	team := TeamConstructor(service.TeamService{Repo: services.Team, TraineeRepo: services.Trainee}, session)
	authRoutes := router.Group("/auth")
	authRoutes.Post("/:event_id/team", team.Create)

	routes := router.Group("/team")
	routes.Get("/event/:event_id", team.List)

}
