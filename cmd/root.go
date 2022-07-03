package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tygen",
	Short: "Tygen is a tool for generating traefik yaml files.",
	Long:  `Tygen is a free and open source tool for generating traefik yaml files. You can also use it to modify your already existing traefik yaml files.`,
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized command"))
	},
}

func Execute() {
	rootCmd.Execute()
}
