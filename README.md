# Kubernetes Pod Monitor CLI

A lightweight Go-based CLI tool that watches all pods in a Kubernetes cluster and logs any containers that have restarted.

Ideal for DevOps engineers, SREs, or platform teams who want real-time visibility into pod restarts for troubleshooting or automation.

---

## 🚀 Features

- Watches all pods across namespaces (or a specific namespace)
- Detects and logs container restarts
- Supports `~/.kube/config` (works out of the box with KinD, Minikube, or EKS)
- Built with `client-go` and `cobra`

---

 ## Roadmap

- Add --namespace and --interval flags
- Slack or webhook alerts
- CrashLoopBackOff filter
- Unit tests
- Dockerfile + Helm chart for K8s deployment

---

## 🛠️ Installation

### Clone the repo

```bash
git clone https://github.com/yourusername/k8s-pod-monitor.git
cd k8s-pod-monitor
