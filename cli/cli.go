package cli

import (
	"flag"
	"fmt"
	"github.com/mynameisdaun/squirtlecoin/explorer"
	"github.com/mynameisdaun/squirtlecoin/rest"
	"os"
)

func usage() {
	fmt.Printf("Welcome to Squirtle Coin!\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000:	Set the PORT of the server\n")
	fmt.Printf("-mode=rest: Choose between 'html' and 'rest'\n\n")
	os.Exit(1)
}

func start() {
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'HTML' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
}
