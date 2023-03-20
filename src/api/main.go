package main

import (
	"github.com/zenthangplus/call-billing-example/src/api/bootstrap"
	"go.uber.org/fx"
)

// @title Sample Public API
// @version 1.0.0
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	fx.New(bootstrap.All()).Run()
}
