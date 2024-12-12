package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

func UploadToCDN(filePath string) (string, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %w", err)
	}

	if fileInfo.Size() > 1*1024*1024 {
		return "", fmt.Errorf("file size exceeds limit: %d bytes", fileInfo.Size())
	}

	cdnURL := os.Getenv("CDN_API_URL")
	if cdnURL == "" {
		return "", errors.New("CDN_API_URL is not set")
	}

	// Inisialisasi Resty client dengan timeout dan mekanisme retry
	client := resty.New().
		SetTimeout(10 * time.Second). // Batas waktu 10 detik
		SetRetryCount(3).             // Maksimal 3 percobaan ulang
		SetRetryWaitTime(2 * time.Second).
		SetRetryMaxWaitTime(10 * time.Second)


	// Kirim file ke CDN API
	resp, err := client.R().
		SetFile("image", filePath).
		Post(cdnURL) // Gunakan CDN URL langsung


	if err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	// Validasi status HTTP response
	if resp.StatusCode() != 200 {
		return "", fmt.Errorf("upload failed with status: %d, response: %s", resp.StatusCode(), resp.String())
	}

	// Parsing JSON respons sesuai struktur
	var result struct {
		Data    struct {
			URL string `json:"url"`
		} `json:"data"`
	}
	
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", fmt.Errorf("failed to parse CDN response: %w", err)
	}
	fmt.Println("result: ", result)

	// Validasi apakah URL tersedia di dalam data
	if result.Data.URL == "" {
		return "", errors.New("CDN response missing 'url' field")
	}

	return result.Data.URL, nil
}
