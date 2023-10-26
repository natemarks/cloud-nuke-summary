package main

import (
	"log"
	"os"

	"github.com/natemarks/cloud-nuke-summary/file"
	"github.com/natemarks/cloud-nuke-summary/version"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stderr).With().Str("version", version.Version).Timestamp().Logger()
	logger.Info().Msg("starting")
	filename := "example.txt" // Change this to the path of your text file
	lines, err := file.ReadFileLines(filename)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	logger.Info().Msgf("Read %d lines from %s", len(lines), filename)
}
