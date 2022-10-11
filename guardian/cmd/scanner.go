package cmd

import (
	"guardian/guardian/services/scanner"

	"github.com/spf13/cobra"
)

func scannerAppPreRun(cmd *cobra.Command, args []string) {
	// database.Initialize()
}

var scannerCmd = cobra.Command{
	Use:   "scanner",
	Short: "Scanner",
}

var scannerHttpCmd = cobra.Command{
	Use:              "http",
	Short:            "Serve HTTP API service",
	PersistentPreRun: scannerAppPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		scanner.StartServer(host, port)
	},
}

var scannerWorkerCmd = cobra.Command{
	Use:              "worker",
	Short:            "Start worker service",
	PersistentPreRun: scannerAppPreRun,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(&scannerCmd)

	scannerCmd.AddCommand(&scannerHttpCmd)
	scannerHttpCmd.Flags().String("host", "127.0.0.1", "Host to bind")
	scannerHttpCmd.Flags().Int("port", 8080, "Port to bind")

	scannerCmd.AddCommand(&scannerWorkerCmd)
}
