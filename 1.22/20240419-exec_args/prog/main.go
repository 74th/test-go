package main

import (
	"flag"
	"log"
)

func main() {
	var flag1 string
	var flag2 string

	flag.StringVar(&flag1, "f1", "", "flag 1")
	flag.StringVar(&flag2, "f2", "", "flag 2")
	flag.Parse()

	log.Printf("FLAG1 %s\n", flag1)
	log.Printf("FLAG2 %s\n", flag2)

	for i, arg := range flag.Args() {
		log.Printf("ARG[%d] %s\n", i, arg)
	}
}
