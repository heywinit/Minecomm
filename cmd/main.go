package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/heywinit/minecomm"
	"github.com/heywinit/minecomm/internal/models"
)

func main() {
	client := minecomm.NewClient()
	playerUUID, _ := uuid.NewUUID()

	player := models.Player{
		Name: "heywinit",
		UUID: playerUUID,
	}

	err := client.Connect("mc.hypixel.net", 25565, player, 754)
	if err != nil {
		return
	}

	fmt.Printf("Connected to %s\n", client.GetAddr())
}
