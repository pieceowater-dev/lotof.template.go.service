package main

import (
	"application/internal"
	"fmt"
	"log"
	"strings"
)

func init() {
	caption := "PIECEOWATER DEV"
	dash10 := strings.Repeat("-", 10)
	space9 := strings.Repeat(" ", 9)
	fmt.Println(dash10 + strings.Repeat("-", len(caption)) + dash10)
	fmt.Println("|" + space9 + caption + space9 + "|")
	fmt.Println(dash10 + strings.Repeat("-", len(caption)) + dash10)
}

func main() {
	err := internal.Run()
	if err != nil {
		log.Fatalf("Application failed during setup: %v", err)
	}
}
