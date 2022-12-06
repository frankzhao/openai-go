package openai

import (
	"encoding/json"
	"errors"
	logger "github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"strings"
)

var imageApiUrl = "https://api.openai.com/v1/images/generations"
var completionApiUrl = "https://api.openai.com/v1/completions"

func New(token string) *Client {
	return &Client{Token: token}
}

func (c *Client) postImageRequest(r ImageRequest) (*ImageResponse, error) {

	reqJson, _ := json.Marshal(r)
	req, _ := http.NewRequest("POST", imageApiUrl, strings.NewReader(string(reqJson)))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		logger.Error().Msgf("Error posting to OpenAI: %v, %v", string(reqJson), err)
		return nil, errors.New("error posting to openai")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	logger.Debug().Msgf("Response from OpenAPI: %s", string(body))
	resp := ImageResponse{}
	err = json.Unmarshal(body, &resp)

	if err != nil {
		logger.Error().Msgf("Error requesting image: %v", err)
		return nil, err
	}

	return &resp, nil
}

func (c *Client) postCompletionRequest(r CompletionRequest) (*CompletionResponse, error) {
	reqJson, _ := json.Marshal(r)
	req, _ := http.NewRequest("POST", completionApiUrl, strings.NewReader(string(reqJson)))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.Token)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		logger.Error().Msgf("Error posting to OpenAI: %v, %v", string(reqJson), err)
		return nil, errors.New("error posting to openai")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	logger.Debug().Msgf("Response from OpenAPI: %s", string(body))
	resp := CompletionResponse{}
	err = json.Unmarshal(body, &resp)

	if err != nil {
		logger.Error().Msgf("Error requesting text completion: %v", err)
		return nil, err
	}

	return &resp, nil
}

// GenerateImage sends a text prompt to OpenAI's image generation API.
func (c *Client) GenerateImage(prompt string, responseFormat string, size string, n int) (*ImageResponse, error) {
	req := ImageRequest{
		Prompt:         prompt,
		N:              n,
		Size:           size,
		ResponseFormat: responseFormat,
	}

	return c.postImageRequest(req)
}

// CompleteText sends a text prompt to OpenAI's text completion generation API.
func (c *Client) CompleteText(prompt, model string, temperature float32, maxTokens int) (*CompletionResponse, error) {
	req := CompletionRequest{
		Model:       model,
		Prompt:      prompt,
		Temperature: temperature,
		MaxTokens:   maxTokens,
	}

	return c.postCompletionRequest(req)
}
