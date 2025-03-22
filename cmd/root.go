package cmd

import (
	"fmt"
	"os"

	"pod-monitor/internal/monitor"

	"github.com/spf13/cobra"
)

var namespace string
var outputFormat string

var rootCmd = &cobra.Command{
	Use:   "pod-monitor",
	Short: "Monitor Kubernetes pods and alert on restarts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running pod-monitor with namespace=%s and output=%s\n", namespace, outputFormat)
		monitor.WatchPods(namespace, outputFormat)
	},
}

func Execute() {
	// Add flags
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Kubernetes namespace to monitor (default: all)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "json", "Output format: text | json")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
