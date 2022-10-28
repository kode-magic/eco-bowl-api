package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	service "github.com/kode-magic/eco-bowl-api/services"
)

type Reward struct {
	Service service.RewardService
	Session *session.Store
}

// RewardConstructor reward constructor
func RewardConstructor(appInterface service.RewardService, session *session.Store) *Reward {
	return &Reward{
		Service: appInterface,
		Session: session,
	}
}

func (i *Reward) Create(ctx *fiber.Ctx) error {
	var requestBody core.Reward

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

	requestBody.Event.ID = requestBody.EventID

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
		"message": "Reward created successfully",
		"data":    res,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (i *Reward) List(ctx *fiber.Ctx) error {
	rewards, err := i.Service.List()
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
		"data":    rewards,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (i *Reward) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	reward, rewardErr := i.Service.Get(id)
	if rewardErr != nil {
		ctxErr := ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Reward with " + id + " not found",
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    reward,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (i *Reward) Update(ctx *fiber.Ctx) error {

	id := ctx.Params("id", "")

	var requestBody core.Reward

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

	res, err := i.Service.Update(&requestBody)

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

func RewardRouter(router fiber.Router, services service.BaseService, session *session.Store) {
	reward := RewardConstructor(service.RewardService{Repo: services.Reward}, session)
	routes := router.Group("/reward")
	routes.Post("/", reward.Create)
	routes.Get("/", reward.List)
	routes.Get("/:id", reward.Get)
	routes.Patch("/:id", reward.Update)

}
