package main

import (
	"context"
	"golang-email-verify/initialize"
)

var ctx context.Context

func main() {
	ctx = context.Background()
	//config load
	initialize.InitConfig(".")
}
