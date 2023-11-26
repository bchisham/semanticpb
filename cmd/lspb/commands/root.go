package commands

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

const (
	appName = "semanticpb"
)

type AppFunc func() error

type Service interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

func LoggerFunc() *zerolog.Logger {
	l := log.With().Str("application-name", appName).Logger()
	return &l
}

var AtExitFuncs []func()
var rootCmd = &cobra.Command{
	Use:   "semanticpb",
	Short: "semanticpb provides tools for using semantic protobuf annotation extensions",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("unable to start application")
	}
}

func RunApp(newService func() Service) {
	service := newService()
	app := fx.New(
		fx.Provide(
			LoggerFunc,
		),
		fx.Invoke(func(lifecycle fx.Lifecycle) {
			lifecycle.Append(
				fx.Hook{
					OnStart: func(ctx context.Context) error {
						service.Start(ctx)
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return service.Stop(ctx)
					},
				})
		}),
		fx.NopLogger,
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal().Err(err)
	}
	defer app.Stop(context.Background())
	select {}
}
