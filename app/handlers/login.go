package handlers

import (
	"os"
	"time"

	"github.com/absensi/app/models"
	"github.com/absensi/app/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Ser *services.Services
}

func (auth *Handler) Login(ctx *fiber.Ctx) error {
	userRequest := new(models.LoginRequest)
	err := ctx.BodyParser(userRequest)
	res := models.LoginResponse{}
	if err != nil {
		res.Code = fiber.StatusExpectationFailed
		res.Message = "Faild To Login"
		return ctx.Status(fiber.StatusAccepted).JSON(res)
	}
	errs := auth.Ser.Validator.ValidateRequest(userRequest)
	if errs != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "request body invalid",
			"error":   errs,
		})
	}
	user := &models.UserRole{}
	err = auth.Ser.Repository.GetUserNip(userRequest.Nip, user)
	if err != nil {
		res.Code = fiber.StatusNoContent
		res.Message = "User No Found"
		return ctx.Status(fiber.StatusAccepted).JSON(res)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password))

	if err != nil {
		res.Code = fiber.StatusNoContent
		res.Message = "Credential invalid"
		return ctx.Status(fiber.StatusAccepted).JSON(res)
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["nip"] = user.Nip
	claims["role"] = user.RoleName
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	// Generate encoded token and send it as response.
	godotenv.Load()
	key := os.Getenv("My_Secret")
	gettoken, err := token.SignedString([]byte(key))
	if err != nil {
		res.Code = fiber.StatusExpectationFailed
		res.Message = "Faild To Login"
		return ctx.Status(fiber.StatusAccepted).JSON(res)
	}

	res.Code = fiber.StatusOK
	res.Message = "Sucsess To Login"
	res.Token = gettoken
	return ctx.Status(fiber.StatusOK).JSON(res)
}
