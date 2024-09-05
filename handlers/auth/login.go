package auth

import (
	"mango/utils"
	"strings"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

func Login(c *fiber.Ctx) error {
	var user LoginRequest

	// parse body
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(LoginResponse{
			Code:    500,
			Message: "Error parsing form.",
		})
	}

	// normalize email
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))

	// check if user exists by email
	query, err := utils.CheckIfUserExistsByEmail(user.Email)

	if err != nil {
		return c.JSON(RegisterResponse{
			Code:    400,
			Message: "Invalid password or email.",
		})
	}

	// compare password and hash
	compare, err := argon2id.ComparePasswordAndHash(user.Password, query.Password)

	if err != nil {
		return c.JSON(RegisterResponse{
			Code:    500,
			Message: "Unable to verify hash.",
		})
	}

	if !compare {
		return c.JSON(RegisterResponse{
			Code:    400,
			Message: "Invalid password or email.",
		})
	}

	return c.JSON(RegisterResponse{
		Code:    200,
		Message: "Logged in successfully.",
	})

}
