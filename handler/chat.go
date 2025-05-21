package handler

import (
	"encoding/json"
	"net/http"
	"os"

	context "context"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

var (
	openaiClient *openai.Client
	defaultModel string
)

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	baseURL := os.Getenv("OPENAI_BASE_URL")
	defaultModel = os.Getenv("OPENAI_MODEL_NAME")

	cfg := openai.DefaultConfig(apiKey)
	if baseURL != "" {
		cfg.BaseURL = baseURL

	}
	openaiClient = openai.NewClientWithConfig(cfg)
}

// 聊天接口，支持流式和非流式
func HandleChatCompletions(c *gin.Context) {
	var req openai.ChatCompletionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// 检查并设置模型
	req.Model = defaultModel

	if req.Stream {
		stream, err := openaiClient.CreateChatCompletionStream(context.Background(), req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		defer stream.Close()
		c.Status(http.StatusOK)
		c.Header("Content-Type", "text/event-stream")
		c.Header("Cache-Control", "no-cache")
		c.Header("Connection", "keep-alive")
		for {
			resp, err := stream.Recv()
			if err != nil {
				break
			}
			data, _ := json.Marshal(resp)
			c.Writer.Write([]byte("data: "))
			c.Writer.Write(data)
			c.Writer.Write([]byte("\n\n"))
			c.Writer.Flush()
		}
		return
	}
	resp, err := openaiClient.CreateChatCompletion(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// 列出模型列表
func HandleListModels(c *gin.Context) {
	resp := gin.H{
		"object": "list",
		"data": []gin.H{
			{
				"id":       defaultModel, // 使用默认模型ID
				"object":   "model",
				"owned_by": "user",
				"created":  1686935002,
			},
		},
	}
	c.JSON(http.StatusOK, resp)
}
