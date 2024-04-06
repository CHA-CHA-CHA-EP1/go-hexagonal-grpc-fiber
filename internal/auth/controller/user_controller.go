package controller

import (
	"errors"
	"strconv"

	"github.com/CHA-CHA-CHA-EP1/go-hexagonal-grpc-fiber/internal/auth/service"
	"github.com/CHA-CHA-CHA-EP1/go-hexagonal-grpc-fiber/internal/auth/domain"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
    GetUserById(c *fiber.Ctx) error
    Register(c *fiber.Ctx) error
}

type userController struct {
    userService service.UserService
}

func NewUserController(
    userService service.UserService,
) UserController {
    return &userController{
        userService: userService,
    }
}

// @summary Get user by id
// @param id path int true "User ID"
// @success 200 {object} Response
// @Failure 400 {object} Response
// @router /users/{id} [get]
func (ctl *userController) GetUserById(c *fiber.Ctx) error {
    // return success response
    id, err := ctl.getUserId(c.Params("id"))
    if err != nil {
        c.Status(fiber.StatusBadRequest)
        c.JSON(fiber.Map{
            "error": err.Error(),
        })
        return err
    }

    user, err := ctl.userService.GetUserById(id)
    if err != nil {
        c.Status(fiber.StatusBadRequest)
        return c.JSON(fiber.Map{
            "error_errror": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "data": user,
    })
}

func (ctl *userController) Register(c *fiber.Ctx) error {
    var newUser domain.UserRegistration

    if err := c.BodyParser(&newUser); err != nil {
        c.Status(fiber.StatusBadRequest)
        return c.JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    user, err := ctl.userService.RegisterUser(newUser)
    if err != nil {
        c.Status(fiber.StatusBadRequest)
        return c.JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "data": user,
    })
}

// private method //
func (ctl *userController) getUserId(userIDParam string) (uint, error) {
    userID, err := strconv.Atoi(userIDParam)
    if err != nil {
        return 0, errors.New("user id should be a number")
    }
    return uint(userID), nil
}
