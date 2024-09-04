package auth

import (
	"mango/db"
	"mango/utils"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/nrednav/cuid2"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Register(c *fiber.Ctx) error {
	var user RegisterRequest

	// parse body
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(RegisterResponse{
			Code:    400,
			Message: "Error parsing form.",
		})
	}

	// check length requirements
	if len(user.Password) < 8 {
		return c.JSON(RegisterResponse{
			Code:    400,
			Message: "Password needs to be a minimum of 8 characters.",
		})
	}

	if len(user.Email) > 254 {
		return c.JSON(RegisterResponse{
			Code:    400,
			Message: "Email exceeds the 254 character limit.",
		})
	}

	// validate email
	if !utils.ValidateEmail(user.Email) {
		return c.JSON(RegisterResponse{
			Code:    400,
			Message: "Invalid email supplied.",
		})
	}

	// check if email is in use
	_, err := utils.CheckIfUserExistsByEmail(user.Email)

	if err == nil {
		return c.JSON(RegisterResponse{
			Code:    400,
			Message: "Email already exists.",
		})
	}

	// generate hash
	hash, err := argon2id.CreateHash(user.Password, argon2id.DefaultParams)

	if err != nil {
		return c.JSON(RegisterResponse{
			Code:    400,
			Message: "Unable to create hash.",
		})
	}

	// create user
	db.DB.Create(&db.User{
		ID:       cuid2.Generate(),
		Email:    user.Email,
		Password: hash,
	})

	return c.JSON(RegisterResponse{
		Code:    200,
		Message: "Created user successfully.",
	})
}
