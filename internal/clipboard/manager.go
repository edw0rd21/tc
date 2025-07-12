package clipboard

import (
	"fmt"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/edw0rd21/tc/internal/storage"
)

type Manager struct {
	storage *storage.Storage
}

func NewManager() (*Manager, error) {
	storage, err := storage.NewStorage(100)
	if err != nil {
		return nil, err
	}
	return &Manager{storage: storage}, nil
}

func (m *Manager) AddItem(content string) {
	m.storage.AddItem(content)
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

type FormatOptions struct {
	Index int
	Full  bool
	Raw   bool
}

func (m *Manager) FormatItem(item storage.ClipboardItem, opts FormatOptions) string {
	content := item.Content
	if !opts.Full && len(content) > 80 {
		content = content[:77] + "..."
	}

	if opts.Raw {
		return content
	}

	content = strings.ReplaceAll(content, "\n", " ")
	content = strings.ReplaceAll(content, "\r", " ")

	timeStr := item.Timestamp.Format("15:04:05")
	return fmt.Sprintf("%d➤ [%s] %s", opts.Index+1, timeStr, content)
}
