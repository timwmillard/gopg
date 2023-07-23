package main

import "flag"

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
	default:
		flag.Usage()
	}

}
