package gatherer

import (
	"github.com/changhoi/slake/internal/repository/postgresql"
	"go.uber.org/zap"
)

type Repository interface {
}

type Gatherer struct {
	logger     *zap.Logger
	repository Repository
}

func New(
	l *zap.Logger,
	pgRepo *postgresql.Repository,
) *Gatherer {
	return &Gatherer{
		logger:     l.Named("gatherer"),
		repository: pgRepo,
	}
}
