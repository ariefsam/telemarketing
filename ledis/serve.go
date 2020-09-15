package main

import (
	"log"

	lediscfg "github.com/ledisdb/ledisdb/config"
	"github.com/ledisdb/ledisdb/server"
)

func main() {
	cfg := lediscfg.NewConfigDefault()
	cfg.Addr = ":6379"
	cfg.HttpAddr = ":8111"
	app, err := server.NewApp(cfg)

	if err != nil {
		log.Fatal(err)
	}
	app.Run()
}
