package main

import (
	"dima/opt/app"
	"github.com/prometheus/common/version"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger.Info().Str("build_context", version.BuildContext()).Str("version", version.Info()).Str("action", "starting").Msg("")

	app.Initialize()
	app.Prepare()
	app.Run()
}
