package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

// ClipboardItem represents a single clipboard entry
type ClipboardItem struct {
	Content   string    `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

// Storage handles clipboard history persistence
type Storage struct {
	filePath string
	maxItems int
}

// NewStorage creates a new storage instance
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

// AddItem adds a new clipboard item to history
func (s *Storage) AddItem(content string) error {
	history, err := s.LoadHistory()
	if err != nil {
		return err
	}
	
	// Don't add if it's the same as the last item
	if len(history) > 0 && history[0].Content == content {
		return nil
	}
	
	// Add new item at the beginning
	newItem := ClipboardItem{
		Content:   content,
		Timestamp: time.Now(),
	}
	
	history = append([]ClipboardItem{newItem}, history...)
	
	// Keep only maxItems
	if len(history) > s.maxItems {
		history = history[:s.maxItems]
	}
	
	return s.saveHistory(history)
}

// LoadHistory loads clipboard history from file
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

// GetLastItems returns the last n items from history
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

// ClearHistory clears all clipboard history
func (s *Storage) ClearHistory() error {
	return os.Remove(s.filePath)
}

// saveHistory saves history to file
func (s *Storage) saveHistory(history []ClipboardItem) error {
	data, err := json.MarshalIndent(history, "", "  ")
	if err != nil {
		return err
	}
	
	return os.WriteFile(s.filePath, data, 0644)
}	