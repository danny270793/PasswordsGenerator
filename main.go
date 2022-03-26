package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/danny270793/passwords-generator/src/passwordsgenerator"
)

const VERSION = "1.0.0"

func main() {
	lengthFlag := flag.Int("length", 16, "number of characters")
	useSymbols := flag.Bool("symbols", true, "use symbols")
	useNumbers := flag.Bool("numbers", true, "use numbers")
	showVersion := flag.Bool("version", false, "show version")
	flag.Parse()

	if *showVersion {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	password := passwordsgenerator.Generate(*lengthFlag, *useSymbols, *useNumbers)
	fmt.Println(password)
}
