package main

import (
	server_grpc "github.com/st-vasyl/echo-server/pkg/v1/server_grpc"
	server_http "github.com/st-vasyl/echo-server/pkg/v1/server_http"
)

func main() {
	go server_grpc.Run()
	server_http.Run()
}
