package auth

import (
	"mango/db"
	"mango/utils"
	"strings"

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
			Code:    500,
			Message: "Error parsing form.",
		})
	}

	// normalize email
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))

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
			Code:    500,
			Message: "Unable to create hash.",
		})
	}

	// create user
	newUser := &db.User{
		ID:       cuid2.Generate(),
		Email:    user.Email,
		Password: hash,
	}

	query := db.DB.Create(newUser)

	if query.Error != nil {
		return c.JSON(RegisterResponse{
			Code:    500,
			Message: "Unable to create new user.",
		})
	}

	return c.JSON(RegisterResponse{
		Code:    200,
		Message: "Created user successfully.",
	})
}
