package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	service "github.com/kode-magic/eco-bowl-api/services"
)

type Entrepreneur struct {
	Service service.EntrepreneurService
	Session *session.Store
}

func EntrepreneurConstructor(service service.EntrepreneurService, session *session.Store) *Entrepreneur {
	return &Entrepreneur{Service: service, Session: session}
}

func (t *Entrepreneur) Add(ctx *fiber.Ctx) error {
	var requestBody core.Entrepreneur

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

	validateErr := requestBody.Validate()

	if len(validateErr) > 0 {
		err := ctx.Status(fiber.StatusUnprocessableEntity).JSON(validateErr)

		if err != nil {
			return err
		}

		return nil
	}

	entrepreneur, addErr := t.Service.Add(&requestBody)

	if addErr != nil {
		ctxErr := ctx.Status(fiber.StatusInternalServerError).JSON(addErr)

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": entrepreneur.Forename + " successfully added as entrepreneur",
		"data":    entrepreneur,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (t *Entrepreneur) List(ctx *fiber.Ctx) error {
	res, err := t.Service.List()

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

func (t *Entrepreneur) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	entrepreneur, traineeErr := t.Service.Get(id)
	if traineeErr != nil {
		ctxErr := ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Entrepreneur with " + id + " not found",
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    entrepreneur,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func EntrepreneurRoute(router fiber.Router, services service.BaseService, session *session.Store) {
	entrepreneur := EntrepreneurConstructor(service.EntrepreneurService{Repo: services.Entrepreneur}, session)

	userRoutes := router.Group("/entrepreneur")
	userRoutes.Post("/", entrepreneur.Add)
	userRoutes.Get("/:id", entrepreneur.Get)
	userRoutes.Get("/", entrepreneur.List)
}
