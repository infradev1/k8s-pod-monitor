package monitor

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// WatchPods scans all pods and logs any with restart counts > 0
func WatchPods() {
	// Load kubeconfig from default location
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error loading kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	namespace := "" // all namespaces; you can parameterize this later

	fmt.Printf("Scanning for pod restarts in all namespaces...\n\n")

	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing pods: %v", err)
	}

	found := false

	for _, pod := range pods.Items {
		for _, cs := range pod.Status.ContainerStatuses {
			if cs.RestartCount > 0 {
				found = true
				fmt.Printf("[RESTART] %s/%s - Container: %s - Restarts: %d\n",
					pod.Namespace, pod.Name, cs.Name, cs.RestartCount)
			}
		}
	}

	if !found {
		fmt.Println("✅ No restarts found.")
	}
}
