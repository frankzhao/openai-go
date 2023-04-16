package openai

const (
	RESPONSE_FORMAT_URL    string = "url"
	RESPONSE_FORMAT_BASE64 string = "b64_json"

	MODEL_GPT_DAVINCI   string = "text-davinci-003"
	MODEL_CODEX_DAVINCI string = "code-davinci-002"
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

// EmbeddingsRequest The Request message for the Embeddings API.
type EmbeddingsRequest struct {
	Model string `json:"model"`
	Inputs []string `json:"inputs"`
	User string `json:"user,omitempty"` 
}

// EmbeddingsResponse The Response message for the Embeddings API
type EmbeddingsResponse struct {
	Object  string `json:"object"`
	Data []struct {
		Object  string `json:"object"`
		Embedding []float32 `json:"embedding"`
		Index int `json:"index"`
	}
	Model   string `json:"model"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}