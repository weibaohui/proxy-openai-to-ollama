package main

import (
	"log"

	"proxy-openai-to-ollama/handler"
	"proxy-openai-to-ollama/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(middleware.LogRequestResponse(), gin.Recovery())

	r.POST("/v1/chat/completions", handler.HandleChatCompletions)
	r.GET("/v1/models", handler.HandleListModels)
	r.GET("/api/tags", handler.HandleListTags) // 添加新的路由
	r.POST("/api/show", handler.HandleShow)    // 添加新路由

	log.Println("[INFO] Proxy server listening on :11434 ...")
	if err := r.Run(":11434"); err != nil {
		log.Fatalf("[FATAL] %v", err)
	}
}
