package main

import (
	"dewkit/cmd"
	"dewkit/config"
	"log/slog"
	"os"
)

func initLogger() {
	env := config.GetEnvDefault("ENV", "production")

	var handler slog.Handler

	if env == "development" {
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	} else {
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	slog.SetDefault(slog.New(handler))
}

func main() {
	initLogger()
	cmd.Execute()
}
