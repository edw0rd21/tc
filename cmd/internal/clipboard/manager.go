package clipboard

import (
	"fmt"
	"strings"
	"tc/internal/storage"
	"time"

	"github.com/atotto/clipboard"
)

// Manager handles clipboard operations
type Manager struct {
	storage *storage.Storage
}

// NewManager creates a new clipboard manager
func NewManager() (*Manager, error) {
	storage, err := storage.NewStorage(100) // Keep last 100 items
	if err != nil {
		return nil, err
	}
	
	return &Manager{
		storage: storage,
	}, nil
}

// StartWatcher starts monitoring clipboard changes
func (m *Manager) StartWatcher() error {
	lastContent := ""
	
	// Get current clipboard content
	current, err := clipboard.ReadAll()
	if err == nil {
		lastContent = current
		if strings.TrimSpace(current) != "" {
			m.storage.AddItem(current)
		}
	}
	
	// Monitor for changes
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	
	for range ticker.C {
		content, err := clipboard.ReadAll()
		if err != nil {
			continue
		}
		
		// Only add if content changed and is not empty
		if content != lastContent && strings.TrimSpace(content) != "" {
			m.storage.AddItem(content)
			lastContent = content
		}
	}
	
	return nil
}

// GetLastItems returns the last n clipboard items
func (m *Manager) GetLastItems(n int) ([]storage.ClipboardItem, error) {
	return m.storage.GetLastItems(n)
}

// CopyToClipboard copies content to system clipboard
func (m *Manager) CopyToClipboard(content string) error {
	return clipboard.WriteAll(content)
}

// ClearHistory clears clipboard history
func (m *Manager) ClearHistory() error {
	return m.storage.ClearHistory()
}

// FormatItem formats a clipboard item for display
func (m *Manager) FormatItem(item storage.ClipboardItem, index int) string {
	content := item.Content
	
	// Truncate long content
	if len(content) > 80 {
		content = content[:77] + "..."
	}
	
	// Replace newlines with spaces for display
	content = strings.ReplaceAll(content, "\n", " ")
	content = strings.ReplaceAll(content, "\r", " ")
	
	// Format timestamp
	timeStr := item.Timestamp.Format("15:04:05")
	
	return fmt.Sprintf("%d: [%s] %s", index+1, timeStr, content)
}