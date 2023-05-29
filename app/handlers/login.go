package handlers

import (
	"github.com/absensi/app/services"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Ser *services.Services
}

func (auth *Handler) Login(ctx *fiber.Ctx) error {
	// userRequest := new(models.UserLoginReq)
	// _ = ctx.BodyParser(userRequest)
	// auth.Ser.Validator.ValidateRequest(userRequest)
	// user := &models.User{}
	// userrole := &models.LoginResponse{}
	// err := auth.Ser.Repository.GetUserEmail(userRequest.Usermail, user)
	// if err != nil || user.ID == 0 {
	// 	return ctx.Status(fiber.StatusAccepted).JSON(
	// 		utils.CustomesResponses{
	// 			Code:    fiber.StatusNoContent,
	// 			Message: "Your Account not Exist",
	// 		})
	// }
	// auth.Ser.Repository.GetUserRole(user.ID, userrole)
	// utils.GenerateEncodeStructs(userrole)
	// b, err := json.Marshal(userrole)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// encryptdata, _ := cryto.Encrypt(b)
	// // Create token
	// token := jwt.New(jwt.SigningMethodHS256)
	// // Set claims
	// claims := token.Claims.(jwt.MapClaims)
	// claims["user"] = base64.StdEncoding.EncodeToString(encryptdata)
	// claims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	// // Generate encoded token and send it as response.
	// godotenv.Load()
	// key := os.Getenv("My_Secret")
	// gettoken, err := token.SignedString([]byte(key))
	// if err != nil {
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(utils.CustomesResponses{
	// 		Code:    fiber.StatusExpectationFailed,
	// 		Message: "Faild To Login",
	// 	})
	// }
	// res := &utils.CustomesResponses{
	// 	Code:    fiber.StatusOK,
	// 	Message: "Login",
	// 	Token:   gettoken,
	// }

	// return ctx.Status(fiber.StatusOK).JSON(map)
	return nil
}
