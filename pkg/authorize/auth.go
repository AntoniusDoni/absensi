package authorize

import (
	"os"

	"github.com/absensi/app/services"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func New() *AuthorizeUser {
	return &AuthorizeUser{}
}

type AuthorizeUser struct {
	Ser *services.Services
}

func (authorizeUser *AuthorizeUser) Auth() fiber.Handler {
	godotenv.Load()
	key := os.Getenv("My_Secret")

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(key),
		ErrorHandler: jwtError,
	})
}

func GetNip(c *fiber.Ctx) (Nip string) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	Nip, _ = claims["nip"].(string)
	return

}
func GetRole(c *fiber.Ctx) (UserId uuid.UUID, Role string) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	Role, _ = claims["role"].(string)
	UserId, _ = uuid.Parse(claims["user_id"].(string))
	return

}

func (authorizeUser *AuthorizeUser) AuthAdmin(c *fiber.Ctx) error {
	godotenv.Load()
	key := os.Getenv("My_Secret")
	jwtware.New(jwtware.Config{
		SigningKey:   []byte(key),
		ErrorHandler: jwtError,
	})
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] == nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{
				"status":  fiber.StatusUnauthorized,
				"message": "Unauthorized",
			})
	}
	return c.Next()

}
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT"})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT"})
}
