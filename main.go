package main

import "pod-monitor/cmd"

func main() {
	cmd.Execute()
}

/*
go mod init pod-monitor
go get k8s.io/client-go@v0.29.0
go get github.com/spf13/cobra@v1.7.0
*/
