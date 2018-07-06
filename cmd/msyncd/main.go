package main

import "flag"

func main() {
	// use this to sync the whole blockchain
	fullSync := flag.Bool("fullsync", false, "When set to true, the daemon will not stop when encountered a known block")

}
