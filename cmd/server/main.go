package main

import (
	"flag"
	"fmt"
	"log"

	"<%= modulePath %>/internal/config"
	"<%= modulePath %>/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to the config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	gin.SetMode(cfg.Server.Mode)
	r := router.New()

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("<%= projectName %> started, listening on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("server exited: %v", err)
	}
}
