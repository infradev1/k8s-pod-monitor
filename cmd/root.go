package cmd

import (
	"fmt"
	"os"

	"pod-monitor/internal/monitor"

	"github.com/spf13/cobra"
)

var namespace string
var outputFormat string
var minRestarts uint

var rootCmd = &cobra.Command{
	Use:   "pod-monitor",
	Short: "Monitor Kubernetes pods and alert on restarts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running pod-monitor with namespace=%s and output=%s\n", namespace, outputFormat)
		monitor.WatchPods(&monitor.UserInput{
			Namespace:       namespace,
			OutputFormat:    outputFormat,
			MinimumRestarts: minRestarts,
		})
	},
}

func Execute() {
	// Add flags
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "Kubernetes namespace to monitor (default: all)")
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "json", "Output format: text | json")
	rootCmd.PersistentFlags().UintVar(&minRestarts, "min-restarts", 1, "Minimum number of pod restarts to monitor (default: 0)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
