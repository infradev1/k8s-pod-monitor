# Kubernetes Pod Monitor CLI

A lightweight Go-based CLI tool that watches all pods in a Kubernetes cluster and logs any containers that have restarted.
If --namespace=kube-system, it monitors the health of the control plane static pods.

Ideal for DevOps engineers, SREs, or platform teams who want real-time visibility into pod health for troubleshooting or automation.

---

## 🚀 Features

- Watches all pods across namespaces (or a specific namespace)
- Detects and logs container restarts
- Supports `~/.kube/config` (works out of the box with KinD, Minikube, or EKS)
- Built with `client-go` and `cobra`

---

 ## Roadmap

- Add --interval flag
- Add subcommands: watch, summary, slack-alert, etc.
- Slack or webhook alerts
- CrashLoopBackOff filter
- Helm chart for K8s deployment

---

## 🛠️ Installation

### Clone the repo

```bash
git clone https://github.com/yourusername/k8s-pod-monitor.git
cd k8s-pod-monitor
make setup-kind      # creates KinD cluster + crashloop pod
make run             # runs CLI against it
make delete-kind     # cleanup