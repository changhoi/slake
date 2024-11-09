package application

import (
	"github.com/changhoi/slake/internal/gatherer"
	"github.com/changhoi/slake/internal/gatherer/controller/mail"
	"github.com/emersion/go-smtp"
	"go.uber.org/fx"
)

func gathererOptions() fx.Option {
	return fx.Provide(
		fx.Annotate(mail.NewGathererController, fx.As(new(smtp.Backend))),
		gatherer.New,
	)
}
