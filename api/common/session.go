package common

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	core "github.com/kode-magic/eco-bowl-api/core/entities"
	"github.com/kode-magic/eco-bowl-api/utils"
	"time"
)

func ChangePasswordSession(session *session.Store, ctx *fiber.Ctx, user *core.User) error {
	sess, sessErr := session.Get(ctx)

	if sessErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + sessErr.Error(),
		})
	}

	sess.Set(utils.KEY, true)
	sess.Set(utils.USERID, user.ID)
	sess.SetExpiry(time.Minute * 5)
	sess.Set(utils.NAME, fmt.Sprintf("%s %s", user.FirstName, user.LastName))

	err := sess.Save()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + err.Error(),
		})
	}

	return nil
}

func LoginSession(session *session.Store, ctx *fiber.Ctx, user *core.User) error {
	sess, sessErr := session.Get(ctx)

	if sessErr != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong: " + sessErr.Error(),
		})
	}

	sess.Set(utils.KEY, true)
	sess.Set(utils.USERID, user.ID)
	sess.Set(utils.NAME, fmt.Sprintf("%s %s", user.FirstName, user.LastName))

	sessSaveErr := sess.Save()

	if sessSaveErr != nil {
		ctxErr := ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": sessSaveErr.Error(),
		})

		if ctxErr != nil {
			return ctxErr
		}

		return nil
	}

	return nil
}
