package main

import (
	"fmt"
	"log"

	"github.com/TorchofFire/uRelay-guild/config"
	"github.com/TorchofFire/uRelay-guild/internal/connections"
	"github.com/TorchofFire/uRelay-guild/internal/database"
	"github.com/TorchofFire/uRelay-guild/internal/guild"
	"github.com/TorchofFire/uRelay-guild/internal/routes"
)

func main() {
	fmt.Println("Guild is starting...")
	config.LoadConfig()
	db, err := database.NewDbConnectionPool()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	guildService := guild.NewService(db)
	connectionsService := connections.NewService(guildService)
	routesService := routes.NewService(guildService, connectionsService)

	routesService.Init()
}
