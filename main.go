package main

import (
	"flag"
	"log"
)

const (
	// Version defines the gophernotes version.
	Version = "1.0.0"

	// ProtocolVersion defines the Jupyter protocol version.
	ProtocolVersion = "5.1"
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
