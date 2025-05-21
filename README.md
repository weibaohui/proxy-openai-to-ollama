# proxy-openai-to-ollama

## 项目简介
`proxy-openai-to-ollama` 是一个用于将 OpenAI API 请求代理到 Ollama API 的代理工具。它可以让你使用 Ollama 的接口格式来调用 OpenAI 的大模型服务。

## 安装
1. 克隆项目代码：
   ```sh
   git clone https://github.com/your-repo/proxy-openai-to-ollama.git
   cd proxy-openai-to-ollama
   ```

2. 安装依赖：
   ```sh
   go mod tidy
   ```

## 配置
1. 设置环境变量：
   ```sh
   export OPENAI_API_KEY=your_openai_api_key
   export OPENAI_BASE_URL=https://api.openai.com
   export OPENAI_MODEL_NAME=qwen3:0.6b
   ```

2. 运行服务：
   ```sh
   go build
   ./proxy-openai-to-ollama
   ```

## API 使用说明

### 1. 聊天接口
将 OpenAI 的聊天接口伪装成 Ollama 格式：

```sh
# 非流式请求
curl http://localhost:8080/api/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "qwen3:0.6b",
    "messages": [
      {
        "role": "user",
        "content": "你好，请介绍一下自己"
      }
    ]
  }'

# 流式请求
curl http://localhost:8080/api/chat/completions \
  -H "Content-Type: application/json" \
  -d '{
    "model": "qwen3:0.6b",
    "messages": [
      {
        "role": "user",
        "content": "你好，请介绍一下自己"
      }
    ],
    "stream": true
  }'
```

### 2. 模型接口
查看可用模型列表：
```sh
curl http://localhost:8080/api/models
```

获取模型详细信息：
```sh
curl http://localhost:8080/api/show/qwen3:0.6b
```

## 注意事项
1. 所有请求会自动转发到配置的OpenAI端点
2. 响应格式与Ollama API保持一致
3. 支持流式和非流式响应
4. 默认使用环境变量中配置的模型名称

