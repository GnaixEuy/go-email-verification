package main

import (
	"context"
	"fmt"
	"golang-email-verify/initialize"
)

var ctx context.Context

func main() {
	ctx = context.Background()
	//config load
	initialize.InitConfig(".")
	fmt.Println(initialize.GetConfig())
	// database load
	initialize.InitDBClient(ctx)
	// close database
	defer initialize.CloseDBClient(ctx)
}
