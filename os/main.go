package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// 標準入力から相対パスを受け取って一覧表示(ディレクトリ想定で)
	if err := setEnv(); err != nil {
		log.Fatalf("failed to set env: %s", err.Error())
	}
	if err := chHomeDir(); err != nil {
		log.Fatalf("failed to change home directory: %s", err.Error())
	}

	var path string
	for {
		n, err := fmt.Scanln(&path)
		if err != nil {
			fmt.Printf("failed to scan standard input: %s", err.Error())
		}
		if n > 0 {
			entries, err := os.ReadDir(path)
			if err != nil {
				fmt.Printf("failed to read directory: %s", err.Error())
			}
			for _, entry := range entries {
				fmt.Printf("%v\n", entry.Name())
			}
		}
	}
}

func setEnv() error {
	envFile, err := os.ReadFile("os/.env")
	if err != nil {
		return fmt.Errorf("failed to read file %w", err)
	}
	rows := strings.Split(string(envFile), "\n")
	for _, row := range rows {
		r := strings.Split(row, "=")
		if len(r) != 2 {
			return fmt.Errorf("falied to parse environment variable %v", row)
		}
		os.Setenv(r[0], r[1])
	}
	return nil
}

func chHomeDir() error {
	homeDir := os.Getenv("HOME_DIR")
	if homeDir == "" {
		homeDir = "."
	}
	if err := os.Chdir(homeDir); err != nil {
		return fmt.Errorf("failed to change directory: %w", err)
	}
	return nil
}
