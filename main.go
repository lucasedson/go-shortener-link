package main

import (
	"github.com/lucasedson/go-shortener-link/config"
	router "github.com/lucasedson/go-shortener-link/router"
)

func main() {
	err := config.Init()
	config.Init()
	if err != nil {
		panic(err)
	}

	router.Initialize()
}
