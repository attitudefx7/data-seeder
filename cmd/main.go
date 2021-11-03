package main

import (
	"data-seeder/internal/server"
	"flag"
	"fmt"
	"log"
)

var (
	cfg = flag.String("c", "./config/config.yaml", "config file path")
)

func main() {
	flag.Parse()

	fmt.Println("hello world!")

	err := server.Init(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	server.Run()
}
