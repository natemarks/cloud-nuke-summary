package main

import (
	"log"
	"os"

	"github.com/natemarks/cloud-nuke-summary/summary"
	"github.com/natemarks/cloud-nuke-summary/version"
	"github.com/rs/zerolog"
)

func main() {
	argsWithoutProg := os.Args[1:]
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stderr).With().Str("version", version.Version).Timestamp().Logger()
	logger.Info().Msg("starting")
	contents, err := summary.GetContentsFromFile(argsWithoutProg[0])
	if err != nil {
		log.Fatal(err)
	}
	logger.Info().Msgf("opened %v: SHA256SUM: %v  with %v lines",
		contents.Filepath,
		contents.Sha256sum,
		len(contents.AllLines),
	)
	summary.PrintReport(contents)
}
