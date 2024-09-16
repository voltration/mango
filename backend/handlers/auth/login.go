package auth

import (
	"mango/utils"
	"strings"
	"time"

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
		return c.JSON(LoginResponse{
			Code:    400,
			Message: "Invalid password or email.",
		})
	}

	// compare password and hash
	compare, err := argon2id.ComparePasswordAndHash(user.Password, query.Password)

	if err != nil {
		return c.JSON(LoginResponse{
			Code:    500,
			Message: "Unable to verify hash.",
		})
	}

	if !compare {
		return c.JSON(LoginResponse{
			Code:    400,
			Message: "Invalid password or email.",
		})
	}

	// create token
	jwt, err := utils.CreateJWT(user.Email)

	if err != nil {
		return c.JSON(LoginResponse{
			Code:    500,
			Message: "Unable to create token.",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		//Secure:   true, // Set to true if using HTTPS
		SameSite: "Strict",
	})

	return c.JSON(LoginResponse{
		Code:    200,
		Message: "Logged in successfully.",
	})

}
