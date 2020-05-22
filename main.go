package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	help     bool
	file     string
	instance string
)

func init() {
	flag.BoolVar(&help, "h", false, "show this page")
	flag.BoolVar(&help, "help", false, "show this page")
	flag.StringVar(&file, "c", "", "Specify the config path")
	flag.StringVar(&file, "config", "", "Specify the config path")
	flag.StringVar(&instance, "i", "default", "Select a command to be run")
	flag.StringVar(&instance, "instance", "default", "Select a command to be run")
	flag.Parse()
	flag.Usage = usage
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, `l3s version: l3s/0.0.1
Usage: l3s [OPTION] {Args}...

Options:
`)
	flag.PrintDefaults()
}

func main() {
	if help {
		flag.Usage()
	}

	var config Config
	if notExist(file) {
		usage()
		panic("[ERROR] must specify the config")
	}

	err := config.Read(file)
	if err != nil {
		panic(err)
	}
	preStartCheck(&config)
	startup(&config, instance)
}
