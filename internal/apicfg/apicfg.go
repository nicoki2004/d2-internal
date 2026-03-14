package apicfg

import (
	"github.com/nicoki2004/d2-internal/internal/database"
	"github.com/nicoki2004/d2-internal/internal/destiny"
	"github.com/nicoki2004/d2-internal/internal/logger"
	"github.com/nicoki2004/d2-internal/internal/repository"
)

type APIConfig struct {
	DB     *database.Queries
	Log    logger.Logger
	Repo   repository.DestinyRepository
	Client destiny.Client
}
