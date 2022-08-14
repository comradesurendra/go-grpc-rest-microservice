package main

import (
	jdconfig "GRPC-SERVE/config"
	"GRPC-SERVE/rpcserver"
	"fmt"
)

func main() {

	fmt.Printf("%+v", jdconfig.Cfg)
	jdconfig.Cfg.DbConnect()
	defer jdconfig.CloseDb()
	rpcserver.StartServer()
}
