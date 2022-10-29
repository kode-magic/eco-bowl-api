package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	service "github.com/kode-magic/eco-bowl-api/services"
)

type Trainee struct {
	Service service.TraineeService
	Session *session.Store
}

func TraineeConstructor(service service.TraineeService, session *session.Store) *Trainee {
	return &Trainee{Service: service, Session: session}
}

func (t *Trainee) Add(ctx *fiber.Ctx) error {
	eventId := ctx.Params("event_id", "")
	var requestBody core.Trainee

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

	requestBody.EventID = eventId
	requestBody.Event.ID = eventId

	validateErr := requestBody.Validate()

	if len(validateErr) > 0 {
		err := ctx.Status(fiber.StatusUnprocessableEntity).JSON(validateErr)

		if err != nil {
			return err
		}

		return nil
	}

	trainee, addErr := t.Service.Add(&requestBody)

	if addErr != nil {
		ctxErr := ctx.Status(fiber.StatusInternalServerError).JSON(addErr)

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Thank you " + trainee.Forename + " for registering for the event.\nWe will keep you updated on everything.",
		"data":    trainee,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (t *Trainee) List(ctx *fiber.Ctx) error {
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

func (t *Trainee) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	trainee, traineeErr := t.Service.Get(id)
	if traineeErr != nil {
		ctxErr := ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Trainee with " + id + " not found",
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    trainee,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func TraineeRoute(router fiber.Router, services service.BaseService, session *session.Store) {
	trainee := TraineeConstructor(service.TraineeService{Repo: services.Trainee}, session)

	userRoutes := router.Group("/trainee")
	userRoutes.Get("/trainee/:id", trainee.Get)

	authRoutes := router.Group("/auth")
	authRoutes.Post("/:event_id/trainee", trainee.Add)
	authRoutes.Get("/:event_id/trainee", trainee.List)
}
