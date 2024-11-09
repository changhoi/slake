package postgresql

import (
	"fmt"
	"github.com/changhoi/slake/internal/adapter/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsnFormat = "host=%s user=%s password=%s dbname=%s port=%d TimeZone=Asia/Seoul"

func NewAdapter(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		dsnFormat,
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Username,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.DBName,
		cfg.PostgreSQL.Port,
	)

	return gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
}
