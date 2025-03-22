package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type PodRestartInfo struct {
	Namespace     string `json:"namespace"`
	PodName       string `json:"pod_name"`
	ContainerName string `json:"container_name"`
	Restarts      int32  `json:"restart_count"`
}

// WatchPods checks for container restarts and prints them in the selected format
func WatchPods(namespace, outputFormat string) {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pods: %v", err)
	}

	var results []PodRestartInfo

	for _, pod := range pods.Items {
		for _, cs := range pod.Status.ContainerStatuses {
			if cs.RestartCount > 0 {
				results = append(results, PodRestartInfo{
					Namespace:     pod.Namespace,
					PodName:       pod.Name,
					ContainerName: cs.Name,
					Restarts:      cs.RestartCount,
				})
			}
		}
	}

	if outputFormat == "json" {
		output, _ := json.MarshalIndent(results, "", "  ")
		fmt.Println(string(output))
	} else {
		if len(results) == 0 {
			fmt.Println("✅ No pod restarts found.")
		} else {
			for _, r := range results {
				fmt.Printf("[RESTART] %s/%s - Container: %s - Restarts: %d\n",
					r.Namespace, r.PodName, r.ContainerName, r.Restarts)
			}
		}
	}
}
