package main

import (
	"github.com/alexanderkarlis/naivechain/block"
	server "github.com/alexanderkarlis/naivechain/socket"
)

var blockchain chan []block.Block

func main() {
	server.Serve()
}
