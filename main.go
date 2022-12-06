package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting aplication2...")
	cfg, err := getConfig()
	if err != nil {
		panic(err)
	}
	StartServer(cfg)
}
