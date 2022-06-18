package main

import (
	"github.com/mynameisdaun/squirtlecoin/cli"
	"github.com/mynameisdaun/squirtlecoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
