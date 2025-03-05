package container

import (
	"tensor-graphql/infrastructure/config"
	"tensor-graphql/infrastructure/database"

	"go.uber.org/zap"
)

type SharedComponent struct {
	Conf *config.Config
	Log  *zap.Logger
	DB   *database.DB
}
