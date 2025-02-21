package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"bytes"
)

const AI_API_URL = "https://api.openai.com/v1/completions"
const API_KEY = "your-api-key" // Replace with actual API key

type AIRequest struct {
	Prompt string `json:"prompt"`
	Model  string `json:"model"`
}

type AIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

// Get AI-Powered Task Suggestion
func GetTaskSuggestion(prompt string) (string, error) {
	requestBody, _ := json.Marshal(AIRequest{
		Prompt: prompt,
		Model:  "text-davinci-003",
	})

	req, _ := http.NewRequest("POST", AI_API_URL, bytes.NewBuffer(requestBody))
	req.Header.Set("Authorization", "Bearer "+API_KEY)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response AIResponse
	json.NewDecoder(resp.Body).Decode(&response)

	if len(response.Choices) > 0 {
		return response.Choices[0].Text, nil
	}

	return "", fmt.Errorf("no suggestions found")
}
