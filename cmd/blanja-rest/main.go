package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/maulidihsan/interop-commerce/config"
	"github.com/maulidihsan/interop-commerce/cmd/blanja-rest/server"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
