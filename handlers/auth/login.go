package auth

import (
	"mango/db"
	"regexp"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
	"github.com/nrednav/cuid2"
)

type user struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data,omitempty"`
}

func Login(c *fiber.Ctx) error {
	var user user

	// parse body
	if err := c.BodyParser(&user); err != nil {
		return c.JSON(Response{
			Code:    400,
			Message: "Error parsing form.",
		})
	}

	// check length requirements
	if len(user.Password) < 8 {
		return c.JSON(Response{
			Code:    400,
			Message: "Password needs to be a minimum of 8 characters.",
		})
	}

	if len(user.Email) > 254 {
		return c.JSON(Response{
			Code:    400,
			Message: "Email exceeds the 254 character limit.",
		})
	}

	// validate email
	if !validateEmail(user.Email) {
		return c.JSON(Response{
			Code:    400,
			Message: "Invalid email supplied.",
		})
	}

	// check if email is in use
	var exUser db.User
	result := db.DB.Where("email = ?", user.Email).First(&exUser)

	if result.Error == nil {
		return c.JSON(Response{
			Code:    400,
			Message: "Email already exists.",
		})
	}

	// generate hash
	hash, err := argon2id.CreateHash(user.Password, argon2id.DefaultParams)

	if err != nil {
		return c.JSON(Response{
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

	return c.JSON(Response{
		Code:    200,
		Message: "Created user successfully.",
	})
}

func validateEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(email)
}
