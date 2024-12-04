package main

import (
	"fmt"

	"github.com/TorchofFire/uRelay-guild/config"
	"github.com/TorchofFire/uRelay-guild/internal/database"
	"github.com/TorchofFire/uRelay-guild/internal/guild"
	"github.com/TorchofFire/uRelay-guild/internal/routes"
)

func main() {
	fmt.Println("Guild is starting...")
	config.LoadConfig()
	database.InitDbConnectionPool()
	defer database.DB.Close()
	guild.Init()
	routes.Init()
}
