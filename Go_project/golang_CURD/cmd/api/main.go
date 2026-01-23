// config->db->router-> run server
package main

import (
	"fmt"
	"log"
	"notes-api/internal/config"
	"notes-api/internal/db"
	"notes-api/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error")
	}
	client, _, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("DB Error")
	}

	defer func() {
		if err := db.Disconnet(client); err != nil {
			log.Printf("mongo Disconnected: %v", err)
		}
	}()

	router := server.NewRouter()

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	if err := router.Run(addr); err != nil {
		log.Fatalf("server falied")
	}

}
