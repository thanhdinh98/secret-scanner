package cmd

import (
	"log"
	"net"

	"guardian/guardian/services/webhook"

	"github.com/spf13/cobra"
)

func webhookAppPreRun(cmd *cobra.Command, args []string) {
}

var webhookCmd = cobra.Command{
	Use:   "webhook",
	Short: "Webhook",
}

var webhookHttpCmd = cobra.Command{
	Use:              "http",
	Short:            "Serve HTTP API service",
	PersistentPreRun: webhookAppPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetInt("port")
		webhook.StartServer(host, port)
	},
}

func init() {
	rootCmd.AddCommand(&webhookCmd)

	localIP := GetOutboundIP().String()
	webhookCmd.AddCommand(&webhookHttpCmd)
	webhookHttpCmd.Flags().String("host", localIP, "Host to bind")
	webhookHttpCmd.Flags().Int("port", 8090, "Port to bind")
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}
