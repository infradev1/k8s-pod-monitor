# Kubernetes Pod Monitor CLI

[![CI](https://github.com/CarlosLaraFP/k8s-pod-monitor/actions/workflows/ci.yml/badge.svg)](https://github.com/CarlosLaraFP/k8s-pod-monitor/actions)

A lightweight Go CLI tool that detects restarting or failed pods in your Kubernetes cluster — including control plane components.

Ideal for DevOps engineers, SREs, or platform teams who want real-time visibility into pod health for troubleshooting or automation.

---

## 🚀 Features

- ✅ Detect pods with restart counts over a threshold (`--min-restarts`)
- ✅ Watch for control plane issues in `kube-system`
- ✅ Support for JSON or plain text output
- ✅ Optional `--watch` mode with polling interval
- ✅ Exit with non-zero status for use in CI/CD
- ✅ Works with KinD out of the box
- ✅ GitHub Actions tested (with live KinD cluster)

---

## 🔧 Flags

* --namespace, -n	    Namespace to scan ("" = all namespaces)
* --output, -o	    json or text
* --min-restarts	    Minimum restarts to report (default 1)
* --watch, -w	        Continuously monitor pods
* --interval, -i	    Polling interval in seconds
* --exit, -e	        Exit with 1 if restarts are found

---

## Roadmap

- Add subcommands: summary, alert, etc.
- Slack or webhook alerts
- Amazon SNS email notifications
- Additional filters
- Self-healing
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

# Detect kube-scheduler error (optional):
make setup-kind
make break-scheduler
kubectl run nginx --image=nginx  # Pending in the default namespace
make build
./pod-monitor -n kube-system -e
make fix-scheduler # wait for the scheduler pod to restart and confirm nginx is now running
make delete-kind