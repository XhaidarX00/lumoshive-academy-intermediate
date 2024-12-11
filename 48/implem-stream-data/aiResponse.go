package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// UploadImage adalah fungsi untuk mengunggah gambar dan mendapatkan respons lengkap dari API.
func GetResponseAi(prompt string, listData []string) string {
	promptAwal := fmt.Sprintf("Kamu adalah asisten virtual untuk menganalisa data penjualan berikut %v\n, berikan jawaban singkat di bawah 200 karakter dari pertanyaan dibawah:", listData)

	// Encode prompt untuk URL
	encodedPrompt := url.QueryEscape(fmt.Sprintf("%s\n%s", promptAwal, prompt))
	apiURL := fmt.Sprintf("https://chatwithai.codesearch.workers.dev/?chat=%s", encodedPrompt)

	// Buat request GET
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return fmt.Sprintf("Gagal membuat request: %v", err)
	}

	// Eksekusi request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Sprintf("Gagal mengirim request: %v", err)
	}
	defer resp.Body.Close()

	// Baca body response
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Sprintf("Gagal membaca response: %v", err)
	}

	// Parsing JSON response
	var jsonResponse map[string]string
	err = json.Unmarshal(resBody, &jsonResponse)
	if err != nil {
		return fmt.Sprintf("Gagal parsing response JSON: %v", err)
	}

	// Ambil data dari JSON response
	result, ok := jsonResponse["data"]
	if !ok || result == "" {
		return "Kurang tau nih, gagal dapet response"
	}

	return result
}
