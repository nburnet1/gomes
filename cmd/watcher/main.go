package main

import (
	"github.com/nburnet1/gomes/pkg/watcher"
)
var (
	cmdPath = "cmd"
	binPath = "bin"
	dependencies = map[string][]string{
		"internal/api": {"cmd/web"},
		"pkg/namespace": {"cmd/namespace"},
	}
	// Excluded directories inside `cmd/`
	pathsToExclude = []string{
		"cmd/test",
		"cmd/ignore",
		"cmd/watcher",
	}
)
func main() {
	watcher.StartWatcher(binPath, cmdPath, dependencies, pathsToExclude)
}