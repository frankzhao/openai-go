package openai

const (
	RESPONSE_FORMAT_URL    string = "url"
	RESPONSE_FORMAT_BASE64 string = "base64_json"
)

type Client struct {
	Token string
}

type ImageRequest struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format,omitempty"`
}

type ImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		Url     string `json:"url"`
		B64Data string `json:"b64_json"`
	} `json:"data"`
}

type CompletionRequest struct {
	Model       string  `json:"model,omitempty"`
	Prompt      string  `json:"prompt,omitempty"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	Temperature float32 `json:"temperature,omitempty"`
}

type CompletionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string `json:"text"`
		Index        int    `json:"index"`
		Logprobs     int    `json:"logprobs"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
