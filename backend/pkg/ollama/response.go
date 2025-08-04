package ollama

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Kocannn/self-dunking-ai/config"
	"github.com/Kocannn/self-dunking-ai/domain"
	"github.com/sirupsen/logrus"
)

func PostPrompt(messages []*domain.Message) (*domain.Ollama, error) {
	cfg := config.GetConfig()

	// Use the model from config
	requestBody := domain.OllamaRequest{
		Model:    cfg.OLLAMA_MODEL,
		Messages: messages,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		logrus.Errorf("Error marshaling request: %v", err)
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/api/chat", cfg.OLLAMA_HOST), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorf("Error sending request to Ollama: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Error reading response body: %v", err)
		return nil, err
	}

	// Debug the response
	logrus.Debugf("Raw response body: %s", string(body))

	// For Ollama streaming responses, we need to collect the content from multiple events
	// Each line is a separate JSON object
	var fullContent string
	var model string

	// Split the response by newlines and process each JSON object
	parts := bytes.Split(body, []byte("\n"))

	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		var streamResp domain.OllamaStreamResponse
		if err := json.Unmarshal(part, &streamResp); err != nil {
			logrus.Warnf("Error unmarshaling part of response: %v", err)
			continue
		}

		if streamResp.Model != "" {
			model = streamResp.Model
		}

		if streamResp.Content != "" {
			fullContent += streamResp.Content
		} else if streamResp.Message.Content != "" {
			fullContent += streamResp.Message.Content
		}
	}

	// Create a complete response
	assistantMessage := domain.Message{
		Role:    "assistant",
		Content: fullContent,
	}

	ollamaResponse := &domain.Ollama{
		Model:    model,
		Done:     true,
		Messages: []domain.Message{assistantMessage},
	}

	return ollamaResponse, nil
}

// New streaming function
func StreamPrompt(w http.ResponseWriter, messages []*domain.Message) error {
	cfg := config.GetConfig()

	requestBody := domain.OllamaRequest{
		Model:    cfg.OLLAMA_MODEL,
		Messages: messages,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		logrus.Errorf("Error marshaling request: %v", err)
		return err
	}

	// Set up SSE headers
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Create HTTP request to Ollama with stream flag
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/chat", cfg.OLLAMA_HOST), bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorf("Error creating request: %v", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logrus.Errorf("Error sending request to Ollama: %v", err)
		return err
	}
	defer resp.Body.Close()

	// Create a scanner to read line by line from response
	scanner := bufio.NewScanner(resp.Body)

	// Stream each chunk as it arrives
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		var streamResp domain.OllamaStreamResponse
		if err := json.Unmarshal([]byte(line), &streamResp); err != nil {
			logrus.Warnf("Error unmarshaling part of response: %v", err)
			continue
		}

		// Extract content from the response
		content := ""
		if streamResp.Content != "" {
			content = streamResp.Content
		} else if streamResp.Message.Content != "" {
			content = streamResp.Message.Content
		}

		// Send each chunk to client
		fmt.Fprintf(w, "data: %s\n\n", content)
		w.(http.Flusher).Flush() // Flush to ensure data is sent immediately
	}

	// Send end-of-stream signal
	fmt.Fprintf(w, "data: [DONE]\n\n")
	w.(http.Flusher).Flush()

	return nil
}
