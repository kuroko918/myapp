package main

import (
	"github.com/kuroko918/myapp/cmd/app/infrastructure"
)

func main() {
	infrastructure.Router.Run(":8080")
}
