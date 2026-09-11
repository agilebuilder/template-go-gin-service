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
	configPath := flag.String("config", "config.yaml", "配置文件路径")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	gin.SetMode(cfg.Server.Mode)
	r := router.New()

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	log.Printf("<%= projectName %> 启动，监听 %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("服务退出: %v", err)
	}
}
