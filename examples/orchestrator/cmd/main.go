package main

import "flag"

func main() {
	flag.Parse()

	// wait for cntrl-c...
	select {}
}
