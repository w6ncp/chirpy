package main

import (
	"fmt"
	"os"
)

func readTemplate(filePath string) string {
	b, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filePath, err)
		return ""
	}
	return string(b)
}
