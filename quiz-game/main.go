package main

import (
	"flag"
	"os"
	"fmt"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "CSV file in the format 'problem, answer'")
	flag.Parse()

	file, err := os.Open(*csvFilename);
	
	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s\n", *csvFilename))
	}

	_ = file
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
