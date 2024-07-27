package main

import (
	"fmt"
	"github.com/heywinit/minecomm"
)

func main() {
	client := minecomm.NewClient()
	err := client.Connect("mc.hypixel.net", "")
	if err != nil {
		return
	}

	fmt.Printf("Connected to %s\n", client.GetAddr())
}
