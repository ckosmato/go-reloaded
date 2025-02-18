package main

import (
	"fmt"
	"log"
	"os"

	"reloaded/central"
)

func main() {
	// handling of original and created file
	if len(os.Args) != 3 {
		fmt.Println("Missing arguments")
	}

	name := os.Args[1]

	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	tmp, err := os.Create(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	defer tmp.Close()

	if err := central.Replace(f, tmp); err != nil {
		log.Fatal(err)
	}

	if err := tmp.Close(); err != nil {
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
