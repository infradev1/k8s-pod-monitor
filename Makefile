setup-kind:
	bash scripts/setup-kind.sh

run:
	go run main.go

test:
	go test ./... -v

build:
	go build -o pod-monitor main.go

fmt:
	go fmt ./...

delete-kind:
	kind delete cluster --name pod-monitor