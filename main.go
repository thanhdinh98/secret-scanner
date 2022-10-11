package main

import (
	"guardian/common/cmd"
	"guardian/config"
	scannerCmd "guardian/guardian/cmd"
)

func main() {
	config.LoadEnv()
	// initialize additional flag options
	cmd.Initialize()
	scannerCmd.Execute()
}
