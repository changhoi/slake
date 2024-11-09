package application

import (
	"context"
	"fmt"
	"github.com/changhoi/slake/internal/adapter/config"
	"github.com/changhoi/slake/internal/adapter/http"
	"github.com/changhoi/slake/internal/adapter/logger"
	"github.com/changhoi/slake/internal/adapter/postgresql"
	smptadapter "github.com/changhoi/slake/internal/adapter/smtp"
	postgresqlrepo "github.com/changhoi/slake/internal/repository/postgresql"
	"github.com/emersion/go-smtp"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func New() *fx.App {
	return fx.New(
		fx.Invoke(
			fx.Annotate(
				invokeHTTP,
				fx.ParamTags(`group:"router"`),
			),
		),
		fx.Module("adapter", adapterOption()),
		fx.WithLogger(func(l *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: l}
		}),
		gathererOptions(),
	)
}

type httpRouter interface {
	Register(r fiber.Router)
}

func invokeHTTP(
	routers []httpRouter,
	l *zap.Logger,
	smtpAdapter *smptadapter.Adapter,
	smtpBackend smtp.Backend,
	app *fiber.App,
	cfg *config.Config,
	lc fx.Lifecycle,
) error {
	l.Info("invoke http server")

	v1 := app.Group("/v1")

	for _, router := range routers {
		router.Register(v1)
	}

	smtpSv := smtpAdapter.NewServer(smtpBackend)

	// web server
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := app.Listen(fmt.Sprintf(":%d", cfg.HTTP.Port)); err != nil {
					l.Panic("failed to start http server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return app.ShutdownWithContext(ctx)
		},
	})

	// smtp server
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				l.Info("smtp server started", zap.Any("port", smtpSv.Addr))
				if err := smtpSv.ListenAndServe(); err != nil {
					l.Panic("failed to start smtp server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return smtpSv.Shutdown(ctx)
		},
	})
	return nil
}

func adapterOption() fx.Option {
	return fx.Provide(
		config.NewAdapter,
		http.NewAdapter,
		smptadapter.NewAdapter,
		logger.NewAdapter,
		postgresql.NewAdapter,
		postgresqlrepo.NewRepository,
	)
}
