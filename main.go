package main

import (
    "flag"
    "io"
    "io/ioutil"
    "log"
    "os"
)

func main() {
    debug := flag.Bool("debug", false, "Log extra info to stderr")
    flag.Parse()
    if flag.NArg() < 1 {
        log.Fatalln("Need a command line argument for the connection file.")
    }
    var logwriter io.Writer = os.Stderr
    if !*debug {
        logwriter = ioutil.Discard
    }
    RunKernel(flag.Arg(0), logwriter)
}
