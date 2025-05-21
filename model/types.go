package model

// Ollama chat completion 请求体结构体（简化版）
type OllamaChatRequest struct {
	Model    string        `json:"model"`
	Messages []interface{} `json:"messages"`
}

// OpenAI chat completion 请求体结构体（简化版）
type OpenAIChatRequest struct {
	Model    string        `json:"model"`
	Messages []interface{} `json:"messages"`
}

// OpenAI chat completion 响应体结构体（简化版）
type OpenAIChatResponse struct {
	ID      string      `json:"id"`
	Object  string      `json:"object"`
	Choices interface{} `json:"choices"`
}
