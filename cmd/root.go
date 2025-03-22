package cmd

import (
	"fmt"
	"os"

	"pod-monitor/internal/monitor"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pod-monitor",
	Short: "Monitor Kubernetes pods and alert on restarts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting pod monitor...")
		monitor.WatchPods()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
