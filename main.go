package main

import (
	"flag"
	"fmt"
)

func main() {
	var configFile string
	flag.StringVar(&configFile, "f", "", "config file to generate functions")

	flag.Parse()

	cmd := flag.Arg(0)
	switch cmd {
	case "init":
		cmdInit()
	case "gen":
		cmdGen(configFile)
	case "gen-test":
		cmdGenTest(configFile)
	case "template create":
		// Copy default templates to working directory
		fmt.Println("Creating customer template")
	default:
		flag.Usage()
	}

}
