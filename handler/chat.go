package handler

import (
	"encoding/json"
	"net/http"
	"os"

	context "context"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

var openaiClient *openai.Client

func init() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	baseURL := os.Getenv("OPENAI_BASE_URL")
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
	modelName := os.Getenv("OPENAI_MODEL_NAME")
	if modelName == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Missing OPENAI_MODEL_NAME env var"})
		return
	}
	resp := gin.H{
		"object": "list",
		"data": []gin.H{
			{
				"id":     modelName,
				"object": "model",
			},
		},
	}
	c.JSON(http.StatusOK, resp)
}

// 获取模型详情
func HandleRetrieveModel(c *gin.Context) {
	modelID := c.Param("model")
	resp, err := openaiClient.GetModel(context.Background(), modelID)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
