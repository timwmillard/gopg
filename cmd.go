package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed init_gopg.yaml

var initConfigFile []byte

func cmdInit() {
	fmt.Println("GoPG init")

	_, err := os.Stat("gopg.yaml")
	if os.IsNotExist(err) {

		err := os.WriteFile("gopg.yaml", initConfigFile, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to create gopg.yaml file")
			os.Exit(1)
		}
		fmt.Println(" - Created gopg.yaml")
	} else {
		fmt.Println(" - gopg.yaml already exists, nothing to do")
	}
}

func cmdGen(configFile string) {
	config, err := readConfig(configFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	generate(config)
}
