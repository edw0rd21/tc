package daemon

import (
	"fmt"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	myclip "github.com/edw0rd21/tc/internal/clipboard"
)

type Watcher struct {
	manager *myclip.Manager
}

func NewWatcher() (*Watcher, error) {
	manager, err := myclip.NewManager()
	if err != nil {
		return nil, err
	}
	return &Watcher{manager: manager}, nil
}

func (w *Watcher) Start() {
	lastContent := ""

	//attempt to capture current clipboard on start
	current, err := clipboard.ReadAll()
	if err == nil && strings.TrimSpace(current) != "" {
		w.manager.AddItem(current)
		lastContent = current
		fmt.Println("Initial clipboard captured.")
	}

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
}

// trim shortens clipboard log output
// unused
func trim(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.TrimSpace(s)
	if len(s) > 60 {
		return s[:57] + "..."
	}
	return s
}
