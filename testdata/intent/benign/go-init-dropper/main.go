package main

import "flag"

var verbose bool

func init() {
	flag.BoolVar(&verbose, "v", false, "verbose output")
}

func main() { flag.Parse() }
