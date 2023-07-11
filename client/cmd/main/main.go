package main

import (
	"go.uber.org/fx"
	"whydah.xyz/grpc-go-playground/client/internal/app"
)

func main() {
	fx.New(app.PackageOptions, app.InternalOptions).Run()
}
