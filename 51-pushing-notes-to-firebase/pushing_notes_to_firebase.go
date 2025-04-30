package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Note struct {
	Content string `json:"content"`
	Date    string `json:"date"`
}

func loadEnv() (string, string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", "", err
	}
	baseURL := os.Getenv("FIREBASE_URL")
	apiKey := os.Getenv("FIREBASE_API_KEY")
	if baseURL == "" || apiKey == "" {
		return "", "", fmt.Errorf("FIREBASE_URL or FIREBASE_API_KEY not set in .env")
	}
	return baseURL, apiKey, nil
}

func saveNote(content string) error {
	baseURL, apiKey, err := loadEnv()
	if err != nil {
		return err
	}

	note := Note{
		Content: content,
		Date:    time.Now().Format("2006-01-02"),
	}

	jsonData, _ := json.Marshal(note)

	url := fmt.Sprintf("%s/notes.json?auth=%s", baseURL, apiKey)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Failed to save note: %s", body)
	}

	fmt.Println("Your note was saved.")
	return nil
}

func showNotes() error {
	baseURL, apiKey, err := loadEnv()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/notes.json?auth=%s", baseURL, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("Failed to fetch notes: %s", body)
	}

	var data map[string]Note
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return err
	}

	for _, note := range data {
		fmt.Printf("%s - %s\n", note.Date, note.Content)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n  mynotes new <your note>\n  mynotes show")
		return
	}

	command := os.Args[1]

	switch command {
	case "new":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a note.")
			return
		}
		note := strings.Join(os.Args[2:], " ")
		if err := saveNote(note); err != nil {
			fmt.Println("Error:", err)
		}
	case "show":
		if err := showNotes(); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Unknown command. Use 'new' or 'show'.")
	}
}
