package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {

	cfg := config.MustLoad()
	fmt.Println(cfg)
	// TODO: инициировать объект конфига

	// TODO: логгер

	// TODO: инит приложения (app)

	//TODO: run grpc
}
