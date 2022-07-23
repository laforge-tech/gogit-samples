package cmd

import(
	"os"

	"github.com/spf13/cobra"

	"github.com/rs/zerolog/log"
)

var rootCmd = &cobra.Command{
	Use: "gogit-samples",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err).Msg("failed")
		os.Exit(1)
	}
  }
  