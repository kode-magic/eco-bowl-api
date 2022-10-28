package controllers

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/kode-magic/eco-bowl-api/api/common"
	enum "github.com/kode-magic/eco-bowl-api/core/commons"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	service "github.com/kode-magic/eco-bowl-api/services"
	"github.com/kode-magic/eco-bowl-api/utils"
	"html"
	"strings"
)

type User struct {
	Service service.User
	Session *session.Store
}

func initUser(service service.User, session *session.Store) *User {
	return &User{Service: service, Session: session}
}

func (u *User) Add(ctx *fiber.Ctx) error {
	var requestBody core.User

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

	validateErr := requestBody.Validate("")

	if len(validateErr) > 0 {
		err := ctx.Status(fiber.StatusUnprocessableEntity).JSON(validateErr)

		if err != nil {
			return err
		}

		return nil
	}

	user, addErr := u.Service.Add(&requestBody)

	if addErr != nil {
		ctxErr := ctx.Status(fiber.StatusInternalServerError).JSON(addErr)

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User created successfully",
		"data":    user,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (u *User) Login(ctx *fiber.Ctx) error {
	var requestBody core.User

	requestBodyErr := ctx.BodyParser(&requestBody)

	if requestBodyErr != nil {
		ctxErr := ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Request error": "Cannot parse json, invalid request body",
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	requestBody.Phone = html.EscapeString(strings.TrimSpace(requestBody.Phone))

	validateErr := requestBody.Validate("login")

	if len(validateErr) > 0 {
		ctxErr := ctx.Status(fiber.StatusUnprocessableEntity).JSON(validateErr)

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	user, loginErr := u.Service.Login(requestBody.Phone, requestBody.Password)

	if loginErr != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": loginErr.Error(),
		})
	}

	sessErr := common.LoginSession(u.Session, ctx, user)

	if sessErr != nil {
		ctxErr := ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "something went wrong",
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"user":    user.PublicUser(),
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (u *User) Logout(ctx *fiber.Ctx) error {

	sess, sessErr := u.Session.Get(ctx)

	if sessErr != nil {
		ctxErr := ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": sessErr.Error(),
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	err := sess.Destroy()

	if err != nil {
		ctxErr := ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": sessErr.Error(),
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Logout successfully",
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (u *User) Edit(ctx *fiber.Ctx) error {

	id := ctx.Params("id", "")

	var requestBody core.User

	requestBodyErr := ctx.BodyParser(&requestBody)

	requestBody.ID = id

	if requestBodyErr != nil {
		ctxErr := ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Request error": "Cannot parse json, invalid request body",
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	validateErr := requestBody.Validate("update")

	if len(validateErr) > 0 {
		ctxErr := ctx.Status(fiber.StatusUnprocessableEntity).JSON(validateErr)

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	_, err := u.Service.Edit(&requestBody)

	if err != nil {
		ctxErr := ctx.Status(fiber.StatusInternalServerError).JSON(err)

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User updated successfully",
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (u *User) User(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	user, err := u.Service.User(id)

	if err != nil {
		ctxErr := ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (u *User) Users(ctx *fiber.Ctx) error {
	res, err := u.Service.Users()

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

func (u *User) ChangePassword(ctx *fiber.Ctx) error {
	sess, err := u.Session.Get(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}

	id := fmt.Sprintf("%s", sess.Get(utils.USERID))

	var requestBody core.User

	requestBody.ID = id

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

	_, err = u.Service.ChangePassword(requestBody)

	if err != nil {
		err := ctx.Status(fiber.StatusNotFound).JSON(err.Error())

		if err != nil {
			return err
		}
		return nil
	}

	ctxError := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Password change successfully",
	})

	if ctxError != nil {
		return ctxError
	}

	return nil
}

func (u *User) Me(ctx *fiber.Ctx) error {
	sess, err := u.Session.Get(ctx)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}

	id := fmt.Sprintf("%s", sess.Get(utils.USERID))

	user, err := u.Service.User(id)

	if err != nil {
		ctxErr := ctx.Status(fiber.StatusUnauthorized).JSON(errors.New("unauthorized").Error())

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxErr := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"data":    user,
	})

	if ctxErr != nil {
		return ctxErr
	}

	return nil
}

func (u *User) Remove(ctx *fiber.Ctx) error {

	id := ctx.Params("id", "")

	_, err := u.Service.Remove(id)

	if err != nil {
		err := ctx.Status(fiber.StatusInternalServerError).JSON(err)

		if err != nil {
			return err
		}
		return nil
	}

	ctxError := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "User removed successfully",
	})

	if ctxError != nil {
		return ctxError
	}

	return nil
}

func (u *User) GetPassword(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	user, err := u.Service.GetPassword(id)

	if err != nil {
		ctxErr := ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	if user.Status != string(enum.Pending) {
		ctxErr := ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	ctxResponse := ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success":  true,
		"password": user.BasePassword,
	})

	if ctxResponse != nil {
		return ctxResponse
	}

	return nil

}

func UserRoute(router fiber.Router, services service.BaseService, session *session.Store) {
	user := initUser(service.User{Repo: services.User}, session)

	userRoutes := router.Group("/user")
	userRoutes.Post("/", user.Add)
	userRoutes.Get("/me", user.Me)
	userRoutes.Get("/logout", user.Logout)
	userRoutes.Get("/", user.Users)
	userRoutes.Patch("/change/password", user.ChangePassword)
	userRoutes.Get("/get-password/:id", user.GetPassword)
	userRoutes.Get("/:id", user.User)
	userRoutes.Patch("/:id", user.Edit)
	userRoutes.Delete("/:id", user.Remove)

	authRoutes := router.Group("/auth")
	authRoutes.Post("/login", user.Login)
}
