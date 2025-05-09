package main

import (
	"flag"
	"fmt"
	"github.com/chenmingyong0423/mcp-ip-geo/internal/server"
)

func main() {
	addr := flag.String("address", "", "The host and port to run the sse server")
	flag.Parse()

	if err := server.Run(*addr); err != nil {
		panic(err)
	}
	fmt.Println("mcp server is running...")
}
