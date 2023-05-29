package repository

import "github.com/absensi/pkg/database"

// "github.com/skp/pkg/redisclient"

type Repository struct {
	Gormdb *database.GormDB

	// RedisClient *redisclient.RedisClient
}

func New(db *database.GormDB) *Repository {
	return &Repository{
		Gormdb: db,
	}
}
