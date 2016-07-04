package cmd

import "github.com/spf13/cobra"

const (
	// SockFilePath is a default path to the socket file path
	SockFilePath = "/tmp/cmpr.sock"
)

// RootCmd is a root command of cmpr
var RootCmd = &cobra.Command{
	Use:   "cmpr",
	Short: "cmpr is a simble COM Port Redirector",
	Long:  "cmpr is a simble COM Port Redirector",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
