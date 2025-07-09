package daemon

import (
	"strings"
	"time"

	"github.com/atotto/clipboard"
	internalcb "github.com/edw0rd21/tc/internal/clipboard"
)

// Watcher handles background clipboard monitoring
type Watcher struct {
	manager *internalcb.Manager
}

// NewWatcher creates a new clipboard watcher
func NewWatcher() (*Watcher, error) {
	manager, err := internalcb.NewManager()
	if err != nil {
		return nil, err
	}

	return &Watcher{
		manager: manager,
	}, nil
}

// Start begins monitoring clipboard changes
func (w *Watcher) Start() error {
	lastContent := ""

	// Get current clipboard content
	current, err := clipboard.ReadAll()
	if err == nil && strings.TrimSpace(current) != "" {
		w.manager.AddItem(current)
		lastContent = current
	}

	// Monitor for changes every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		content, err := clipboard.ReadAll()
		if err != nil {
			continue
		}

		if content != lastContent && strings.TrimSpace(content) != "" {
			w.manager.AddItem(content)
			lastContent = content
		}
	}

	return nil
}
