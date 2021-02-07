package main

import (
	"github.com/kuroko918/myapp/cmd/grpc-app/infrastructure/router"
)

func main() {
	go router.Server()
	router.ReverseProxyServer()
}
