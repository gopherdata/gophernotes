package main

import (
	"flag"
	"log"
)

func main() {

	// Parse the connection file.
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatalln("Need a command line argument specifying the connection file.")
	}

	// Run the kernel.
	runKernel(flag.Arg(0))
}
