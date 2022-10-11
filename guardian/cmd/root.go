package cmd

import (
	"guardian/common/cmd"

	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{
	Short: "Guardian",
}

func Execute() {
	cmd.ExecuteRootCmd(&rootCmd)
}
