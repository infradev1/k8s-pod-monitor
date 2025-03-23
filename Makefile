setup-kind:
	bash scripts/setup-kind.sh

clean:
	go mod tidy

run:
	go run main.go

test:
	go test ./... -v

build:
	go build -o pod-monitor main.go

fmt:
	go fmt ./...

break-scheduler:
	docker exec pod-monitor-control-plane \
	bash -c "sed -i 's|--kubeconfig=/etc/kubernetes/scheduler.conf|--kubeconfig=/etc/kubernetes/invalid.conf|' /etc/kubernetes/manifests/kube-scheduler.yaml"

fix-scheduler:
	docker exec pod-monitor-control-plane \
	bash -c "sed -i 's|--kubeconfig=/etc/kubernetes/invalid.conf|--kubeconfig=/etc/kubernetes/scheduler.conf|' /etc/kubernetes/manifests/kube-scheduler.yaml"

delete-kind:
	kind delete cluster --name pod-monitor