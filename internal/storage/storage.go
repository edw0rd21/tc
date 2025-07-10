package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type ClipboardItem struct {
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type Storage struct {
	filePath string
	maxItems int
}

func NewStorage(maxItems int) (*Storage, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configDir := filepath.Join(homeDir, ".tc")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return nil, err
	}

	return &Storage{
		filePath: filepath.Join(configDir, "history.json"),
		maxItems: maxItems,
	}, nil
}

func (s *Storage) AddItem(content string) error {
	history, err := s.LoadHistory()
	if err != nil {
		return err
	}

	if len(history) > 0 && history[0].Content == content {
		return nil
	}

	// add new item at the beginning
	newItem := ClipboardItem{
		Content:   content,
		Timestamp: time.Now(),
	}

	history = append([]ClipboardItem{newItem}, history...)

	// keep only maxItems
	if len(history) > s.maxItems {
		history = history[:s.maxItems]
	}

	return s.saveHistory(history)
}

func (s *Storage) LoadHistory() ([]ClipboardItem, error) {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return []ClipboardItem{}, nil
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	var history []ClipboardItem
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, err
	}

	return history, nil
}

func (s *Storage) GetLastItems(n int) ([]ClipboardItem, error) {
	history, err := s.LoadHistory()
	if err != nil {
		return nil, err
	}

	if len(history) < n {
		return history, nil
	}

	return history[:n], nil
}

func (s *Storage) ClearHistory() error {
	return os.Remove(s.filePath)
}

func (s *Storage) saveHistory(history []ClipboardItem) error {
	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}
