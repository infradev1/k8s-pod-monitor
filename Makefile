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

delete-kind:
	kind delete cluster --name pod-monitor