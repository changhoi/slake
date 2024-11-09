package mail

import (
	"github.com/emersion/go-smtp"
	"go.uber.org/zap"
)

type Controller struct {
	logger *zap.Logger
}

func (c *Controller) NewSession(conn *smtp.Conn) (smtp.Session, error) {
	return &Session{
		host: conn.Hostname(),
	}, nil
}

func NewGathererController(l *zap.Logger) *Controller {
	return &Controller{
		logger: l.Named("gatherer").Named("controller").Named("mail"),
	}
}
