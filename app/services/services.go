package services

import (
	"github.com/absensi/app/repository"
	"github.com/absensi/pkg/database"
	"github.com/absensi/pkg/validator"
)

type Services struct {
	Db *database.GormDB
	// RedisClient *redisclient.RedisClient
	Validator  *validator.Validator
	Repository *repository.Repository
	// Auth       *authorize.AuthorizeUser
}

func New() *Services {

	validator := validator.New()
	gormdb := database.New()
	// redisclient := redisclient.New()
	// Auth := authorize.New()
	repository := repository.New(gormdb)

	return &Services{
		Db: gormdb,
		// RedisClient: redisclient,
		Validator:  validator,
		Repository: repository,
		// Auth:       Auth,
	}
}
