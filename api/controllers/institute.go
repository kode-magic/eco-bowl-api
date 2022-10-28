package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	service "github.com/kode-magic/eco-bowl-api/services"
)

type Centre struct {
	Service service.InstituteService
	Session *session.Store
}

// InstituteConstructor Supply constructor
func InstituteConstructor(appInterface service.InstituteService, session *session.Store) *Centre {
	return &Centre{
		Service: appInterface,
		Session: session,
	}
}

func (i *Centre) Create(ctx *fiber.Ctx) error {
	var requestBody core.Institution

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
		"message": "Institution created successfully",
		"data":    res,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (i *Centre) List(ctx *fiber.Ctx) error {
	departments, err := i.Service.List()
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
		"data":    departments,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (i *Centre) Get(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	department, departErr := i.Service.Get(id)
	if departErr != nil {
		ctxErr := ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
			"message": "Institution with " + id + " not found",
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    department,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (i *Centre) Update(ctx *fiber.Ctx) error {

	id := ctx.Params("id", "")

	var requestBody core.Institution

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

func InstituteRouter(router fiber.Router, services service.BaseService, session *session.Store) {
	institute := InstituteConstructor(service.InstituteService{Repo: services.Institute}, session)
	routes := router.Group("/institute")
	routes.Post("/", institute.Create)
	routes.Get("/", institute.List)
	routes.Get("/:id", institute.Get)
	routes.Patch("/:id", institute.Update)

}
