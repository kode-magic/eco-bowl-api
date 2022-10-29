package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	service "github.com/kode-magic/eco-bowl-api/services"
)

type Solution struct {
	Service service.SolutionService
	Session *session.Store
}

// SolutionConstructor solution constructor
func SolutionConstructor(appInterface service.SolutionService, session *session.Store) *Solution {
	return &Solution{
		Service: appInterface,
		Session: session,
	}
}

func (i *Solution) Create(ctx *fiber.Ctx) error {
	eventId := ctx.Params("event_id", "")
	var requestBody core.Solution

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

	requestBody.EventID = eventId
	requestBody.Event.ID = eventId
	requestBody.Team.ID = requestBody.TeamID

	validateErr := requestBody.Validate()

	if len(validateErr) > 0 {
		err := ctx.Status(fiber.StatusUnprocessableEntity).JSON(validateErr)

		if err != nil {
			return err
		}
		return nil
	}

	res, resErr := i.Service.Create(&requestBody)

	if resErr != nil {
		err := ctx.Status(fiber.StatusInternalServerError).JSON(resErr)

		if err != nil {
			return err
		}
		return nil
	}

	ctxErr := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Solution created successfully",
		"data":    res,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (i *Solution) List(ctx *fiber.Ctx) error {
	eventId := ctx.Params("event_id", "")
	solutions, err := i.Service.List(eventId)
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
		"data":    solutions,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (i *Solution) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	solution, rewardErr := i.Service.Get(id)
	if rewardErr != nil {
		ctxErr := ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Solution with " + id + " not found",
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    solution,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (i *Solution) AddedReward(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	var requestBody service.SolutionRewardRequest

	requestBodyErr := ctx.BodyParser(&requestBody)

	requestBody.Solution = id

	if requestBodyErr != nil {
		err := ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": requestBodyErr.Error(),
		})

		if err != nil {
			return err
		}

		return nil
	}

	res, err := i.Service.AddedReward(&requestBody)

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

func SolutionRouter(router fiber.Router, services service.BaseService, session *session.Store) {
	solution := SolutionConstructor(service.SolutionService{Repo: services.Solution}, session)
	routes := router.Group("/solution")
	routes.Post("/", solution.Create)
	routes.Get("/", solution.List)
	routes.Get("/:id", solution.Get)
	routes.Patch("/:id", solution.AddedReward)

	authRoutes := router.Group("/auth")
	authRoutes.Post("/:event_id/solution", solution.Create)
}
