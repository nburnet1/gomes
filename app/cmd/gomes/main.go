package main

import (
	"gomes/pkg/err"
	"gomes/server"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

func main() {
	config := zap.Config{
        Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
        Development:      true,
        Encoding:         "console", // Use "json" for JSON format
        EncoderConfig:    zap.NewDevelopmentEncoderConfig(),
        OutputPaths:      []string{"stdout"},
        ErrorOutputPaths: []string{"stderr"},
    }

	logger, _ := config.Build()
	defer logger.Sync()
	

	mainError := err.NewError(logger)

	err := setWorkingDirectoryToRoot(os.Args[0])

	mainError.SetErrorAndFatalIfNotNil(err)

	val, _ := os.Getwd()
	logger.Info(val)

	server.Start(mainError)
}

func setWorkingDirectoryToRoot(mainFile string) error {
	mainDir := filepath.Dir(mainFile)
	return os.Chdir(mainDir + "/../..")
}


