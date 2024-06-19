package util

import (
	"fmt"
	"io"
	"os"
)

func WriteToFile(filePath, code string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(code)
	if err != nil {
		return fmt.Errorf("could not write to file: %w", err)
	}
	return nil
}

func ReadFromFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	content := ""
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
        if err != nil && err != io.EOF {
            return "", err
        }
        if n == 0 {
            break
        }
        content += string(buffer[:n])
	}
	return content, nil
}