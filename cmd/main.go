package main

import (
	"fmt"

	"github.com/heywinit/minecomm"
)

func main() {
	client := minecomm.NewClient()	
	client.Connect("mc.hypixel.net", "")

	fmt.Printf("Connected to %s\n", client.GetAddr())
}
