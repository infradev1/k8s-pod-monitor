package monitor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type PodRestartInfo struct {
	Namespace     string `json:"namespace"`
	PodName       string `json:"pod_name"`
	ContainerName string `json:"container_name"`
	Restarts      uint   `json:"restart_count"`
}

// Tags are metadata — they control how the struct is serialized/deserialized (e.g., to/from JSON, YAML, DB rows).

type UserInput struct {
	Namespace       string
	OutputFormat    string
	MinimumRestarts uint
	Watch           bool
	Interval        uint32
	ExitWithError   bool
}

const KubeSystem string = "kube-system"

// WatchPods checks for container restarts and prints them in the selected format
func WatchPods(args *UserInput) {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	scanPods := func() {
		pods, err := clientset.CoreV1().Pods(args.Namespace).List(context.TODO(), v1.ListOptions{})
		if err != nil {
			log.Fatalf("Error listing pods: %v", err)
		}

		if args.Namespace == KubeSystem {
			expected := []string{"kube-scheduler", "kube-controller-manager", "kube-proxy", "coredns", "etcd"}
			seen := make(map[string]bool)
			for _, name := range expected {
				seen[name] = false
			}
			for _, pod := range pods.Items {
				for _, name := range expected {
					if strings.Contains(pod.Name, name) {
						seen[name] = true
					}
				}
				for _, cs := range pod.Status.ContainerStatuses {
					if cs.State.Waiting != nil || cs.State.Terminated != nil {
						log.Printf("[CONTROL PLANE WARNING] Pod: %s - Status: %s", pod.Name, pod.Status.Phase)
						if cs.State.Waiting != nil {
							log.Printf("Waiting Reason: %s", cs.State.Waiting.Reason)
						}
						if cs.State.Terminated != nil {
							log.Printf("Terminated Reason: %s (Exit %d)", cs.State.Terminated.Reason, cs.State.Terminated.ExitCode)
						}
					}
				}
			}
			for name, found := range seen {
				if !found {
					log.Printf("[MISSING] Static pod '%s' not detected", name)
				}
			}
		}

		var results []PodRestartInfo

		for _, pod := range pods.Items {
			for _, cs := range pod.Status.ContainerStatuses {
				if cs.RestartCount > 0 {
					results = append(results, PodRestartInfo{
						Namespace:     pod.Namespace,
						PodName:       pod.Name,
						ContainerName: cs.Name,
						Restarts:      uint(cs.RestartCount),
					})
				}
			}
		}

		filteredResults := FilterRestartedPods(results, args.MinimumRestarts)

		if len(filteredResults) == 0 {
			fmt.Println("✅ No pod restarts found.")
		} else if args.OutputFormat == "json" {
			output, _ := json.MarshalIndent(filteredResults, "", "  ")
			fmt.Println(string(output))
		} else {
			for _, r := range filteredResults {
				fmt.Printf("[RESTART] %s/%s - Container: %s - Restarts: %d\n",
					r.Namespace, r.PodName, r.ContainerName, r.Restarts)
			}
		}

		if args.ExitWithError && len(filteredResults) > 0 {
			os.Exit(1)
		}
	}

	scanPods()

	for args.Watch {
		time.Sleep(time.Duration(args.Interval) * time.Second)
		scanPods()
	}
}
