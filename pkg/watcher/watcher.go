package watcher

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	watcher       *fsnotify.Watcher
	processes     = make(map[string]*exec.Cmd)
	mu            sync.Mutex
	debounceTimes = make(map[string]time.Time)
	debounceLock  sync.Mutex
	debounceDelay = 500 * time.Millisecond
)

type watcherParams struct {
	cmdPath        string
	binPath        string
	dependencies   map[string][]string
	pathsToExclude []string
}

func StartWatcher(binPath string, cmdPath string, dependencies map[string][]string, pathsToExclude []string) {
	watcherParams := watcherParams{
		cmdPath:        cmdPath,
		binPath:        binPath,
		dependencies:   dependencies,
		pathsToExclude: pathsToExclude,
	}
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Handle graceful shutdown
	go handleShutdown()

	// Ensure bin directory exists
	if _, err := os.Stat(watcherParams.binPath); os.IsNotExist(err) {
		os.Mkdir(watcherParams.binPath, os.ModePerm)
	}

	// Watch `cmd/` directories while excluding specified paths
	err = filepath.Walk(watcherParams.cmdPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			relativePath := strings.TrimPrefix(path, watcherParams.cmdPath+"/")
			if watcherParams.isExcludedPath("cmd/" + relativePath) {
				log.Printf("🚫 Skipping: %s\n", path)
				return filepath.SkipDir
			}
			log.Printf("🔍 Watching: %s\n", path)
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Watch additional non-cmd paths from dependencies
	for dep := range watcherParams.dependencies {
		if err := watcher.Add(dep); err != nil {
			fmt.Printf("⚠️ Failed to watch %s: %v\n", dep, err)
		} else {
			fmt.Printf("🔍 Watching dependency: %s\n", dep)
		}
	}

	// Process events
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if !strings.HasSuffix(event.Name, ".go") {
				continue
			}

			// Determine which directory changed
			dir := watcherParams.getWatchedDir(event.Name)
			if dir == "" {
				continue
			}

			// Skip excluded paths
			if watcherParams.isExcludedPath(dir) {
				continue
			}

			// Handle debouncing per binary
			debounceLock.Lock()
			lastBuildTime, exists := debounceTimes[dir]
			if exists && time.Since(lastBuildTime) < debounceDelay {
				debounceLock.Unlock()
				continue
			}
			debounceTimes[dir] = time.Now()
			debounceLock.Unlock()

			// Restart the service and its dependencies
			go func(dir string) {
				_, isDependency := watcherParams.dependencies[dir]
				if isDependency {
					for _, dep := range watcherParams.getDependencies(dir) {
						watcherParams.buildAndRestart(dep)
					}
				}else {
					watcherParams.buildAndRestart(dir)
				}
			}(dir)

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Printf("⚠️ Watcher error: %v\n", err)
		}
	}
}

// getWatchedDir extracts the watched directory from a file path
func (w *watcherParams) getWatchedDir(filePath string) string {
	parts := strings.Split(filePath, string(os.PathSeparator))

	// Ensure we have a valid path
	if len(parts) < 2 {
		return ""
	}

	// Handle cases like `admin/api/main.go`
	for dep := range w.dependencies {
		if strings.HasPrefix(filePath, dep) {
			return dep
		}
	}

	// Handle `cmd/` paths
	if parts[0] == "cmd" {
		return "cmd/" + parts[1]
	}

	return ""
}

// handleShutdown catches SIGINT and SIGTERM and stops all running processes
func handleShutdown() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c // Block until a signal is received

	fmt.Println("\n🛑 Stopping all running processes before exit...")
	stopAllProcesses()
	os.Exit(0)
}

// buildAndRestart compiles and restarts the binary
func (w *watcherParams) buildAndRestart(dir string) {
	fmt.Printf("🔧 Change detected in %s → Rebuilding...\n", dir)

	// Extract last part of dir for binary name
	binaryName := filepath.Base(dir)
	binaryPath := filepath.Join(w.binPath, binaryName)
	sourcePath := dir

	// Ensure we are only building values from the dependencies map
	// If this `dir` is a dependency *key*, do NOT build it
	for key, depList := range w.dependencies {
		if key == dir {
			fmt.Printf("ℹ️ Skipping build for %s (dependencies will be restarted instead)\n", dir)
			for _, dep := range depList {
				w.buildAndRestart(dep)
			}
			return
		}
	}

	// Build the binary
	cmd := exec.Command("go", "build", "-o", binaryPath, "./"+sourcePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Build failed for %s: %v\n", dir, err)
		return
	}

	fmt.Printf("✅ Built %s successfully!\n", dir)

	// Set executable permission
	err := os.Chmod(binaryPath, 0755)
	if err != nil {
		fmt.Printf("⚠️ Failed to set executable permission for %s: %v\n", binaryPath, err)
		return
	}

	// Restart the binary
	restartBinary(binaryName, binaryPath)
}

// restartBinary stops the old process and starts a new one
func restartBinary(dir, binaryPath string) {
	mu.Lock()
	defer mu.Unlock()

	// Stop old process if running
	if proc, exists := processes[dir]; exists && proc.Process != nil {
		fmt.Printf("🛑 Stopping old instance of %s...\n", dir)
		_ = proc.Process.Kill()
		delete(processes, dir)
	}

	// Start new process
	fmt.Printf("🚀 Starting new instance of %s...\n", dir)
	newProc := exec.Command(binaryPath)
	newProc.Stdout = os.Stdout
	newProc.Stderr = os.Stderr

	if err := newProc.Start(); err != nil {
		fmt.Printf("❌ Failed to start %s: %v\n", dir, err)
		return
	}

	processes[dir] = newProc
}

// stopAllProcesses stops all running binaries before exiting
func stopAllProcesses() {
	mu.Lock()
	defer mu.Unlock()

	for dir, proc := range processes {
		if proc.Process != nil {
			fmt.Printf("🛑 Stopping %s...\n", dir)
			_ = proc.Process.Kill()
		}
	}

	processes = make(map[string]*exec.Cmd) // Clear process map
}

// Checks if a cmd subdirectory should be excluded
func (w *watcherParams) isExcludedPath(path string) bool {
	for _, exclude := range w.pathsToExclude {
		if strings.HasPrefix(path, exclude) {
			return true
		}
	}
	return false
}

// Retrieves dependencies for a given directory
func (w *watcherParams) getDependencies(dir string) []string {
	if deps, exists := w.dependencies[dir]; exists {
		return deps
	}
	return nil
}
