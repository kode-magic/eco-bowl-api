package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	service "github.com/kode-magic/eco-bowl-api/services"
)

type Event struct {
	Service service.EventService
	Session *session.Store
}

// EventConstructor event constructor
func EventConstructor(appInterface service.EventService, session *session.Store) *Event {
	return &Event{
		Service: appInterface,
		Session: session,
	}
}

func (e *Event) Create(ctx *fiber.Ctx) error {
	var requestBody core.Event

	requestBodyErr := ctx.BodyParser(&requestBody)

	if requestBodyErr != nil {
		err := ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": requestBodyErr.Error(),
		})

		if err != nil {
			return err
		}

		return nil
	}

	requestBody.Institution.ID = requestBody.InstitutionID

	validateErr := requestBody.Validate()

	if len(validateErr) > 0 {
		err := ctx.Status(fiber.StatusUnprocessableEntity).JSON(validateErr)

		if err != nil {
			return err
		}
		return nil
	}

	res, resErr := e.Service.Create(&requestBody)

	if resErr != nil {
		err := ctx.Status(fiber.StatusInternalServerError).JSON(resErr)

		if err != nil {
			return err
		}
		return nil
	}

	ctxErr := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Event created successfully",
		"data":    res,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (e *Event) List(ctx *fiber.Ctx) error {
	events, err := e.Service.List()
	if err != nil {
		ctxErr := ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    events,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (e *Event) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	event, eventErr := e.Service.Get(id)
	if eventErr != nil {
		ctxErr := ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Event with " + id + " not found",
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    event,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (e *Event) Update(ctx *fiber.Ctx) error {

	id := ctx.Params("id", "")

	var requestBody core.Event

	requestBodyErr := ctx.BodyParser(&requestBody)

	requestBody.ID = id

	if requestBodyErr != nil {
		err := ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": requestBodyErr.Error(),
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

	res, err := e.Service.Update(&requestBody)

	if err != nil {
		err := ctx.Status(fiber.StatusInternalServerError).JSON(err)

		if err != nil {
			return err
		}
		return nil
	}

	ctxError := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    res,
	})

	if ctxError != nil {
		return ctxError
	}

	return nil
}

func EventRouter(router fiber.Router, services service.BaseService, session *session.Store) {
	event := EventConstructor(service.EventService{Repo: services.Event}, session)
	routes := router.Group("/event")
	routes.Post("/", event.Create)
	routes.Get("/", event.List)
	routes.Get("/:id", event.Get)
	routes.Patch("/:id", event.Update)

}
