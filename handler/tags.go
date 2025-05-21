package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleListTags(c *gin.Context) {
	mockResponse := gin.H{
		"models": []gin.H{
			{
				"name":        defaultModel,
				"model":       defaultModel,
				"modified_at": "2025-05-21T16:04:05.677258156+08:00",
				"size":        522653526,
				"digest":      "3bae9c93586b27bedaa979979733c2b0edd1d0defc745e9638f2161192a0ccf0",
				"details": gin.H{
					"parent_model":       "",
					"format":             "gguf",
					"family":             "ollama",
					"families":           []string{"ollama"},
					"parameter_size":     "751.63M",
					"quantization_level": "Q4_K_M",
				},
			},
		},
	}

	c.JSON(http.StatusOK, mockResponse)
}
