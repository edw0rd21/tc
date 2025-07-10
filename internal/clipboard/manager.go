package clipboard

import (
	"fmt"
	"strings"

	"github.com/edw0rd21/tc/internal/storage"

	"github.com/atotto/clipboard"
)

type Manager struct {
	storage *storage.Storage
}

func (m *Manager) AddItem(content string) {
	m.storage.AddItem(content)
}

func NewManager() (*Manager, error) {
	storage, err := storage.NewStorage(100) // Keep last 100 items
	if err != nil {
		return nil, err
	}

	return &Manager{
		storage: storage,
	}, nil
}

func (m *Manager) GetLastItems(n int) ([]storage.ClipboardItem, error) {
	return m.storage.GetLastItems(n)
}

func (m *Manager) CopyToClipboard(content string) error {
	return clipboard.WriteAll(content)
}

func (m *Manager) ClearHistory() error {
	return m.storage.ClearHistory()
}

func (m *Manager) FormatItem(item storage.ClipboardItem, index int) string {
	content := item.Content

	//truncate long content
	if len(content) > 80 {
		content = content[:77] + "..."
	}

	//replace newlines with spaces
	content = strings.ReplaceAll(content, "\n", " ")
	content = strings.ReplaceAll(content, "\r", " ")

	timeStr := item.Timestamp.Format("15:04:05")

	return fmt.Sprintf("%d➤ [%s] %s", index+1, timeStr, content)
}
