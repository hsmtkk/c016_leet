package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hsmtkk/c016_leet/pkg/leet"
)

func main() {
	input := os.Args[1]
	output, err := leet.Leet(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
