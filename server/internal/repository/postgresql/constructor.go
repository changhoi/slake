package postgresql

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewRepository(
	l *zap.Logger,
	db *gorm.DB,
) *Repository {
	return &Repository{
		logger: l.Named("repository").Named("postgresql"),
		db:     db,
	}
}
