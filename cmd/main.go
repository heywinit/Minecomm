package main

import (
	"fmt"
	"github.com/heywinit/minecomm/config"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	fmt.Printf("%+v\n", cfg)
}
