run:
	go run main.go

build:
	go build -o pod-monitor main.go

fmt:
	go fmt ./...

setup-kind:
	bash scripts/setup-kind.sh

delete-kind:
	kind delete cluster --name pod-monitor