package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if err := setEnv(); err != nil {
		log.Fatalf("failed to set env: %s", err.Error())
	}
	homeDir, err := chHomeDir()
	if err != nil {
		log.Fatalf("failed to change home directory: %s", err.Error())
	}
	var cmd, value string
	for {
		currentDir, err := filepath.Abs(".")
		if err != nil {
			log.Fatalf("failed to abs current directory: %s", err.Error())
		}
		// 相対パスを入力を受け付ける左側につける
		path, err := filepath.Rel(homeDir, currentDir)
		if err != nil {
			log.Fatalf("failed to calcurate relative path: %s", err.Error())
		}
		fmt.Printf("%s:", path)
		n, err := fmt.Scanln(&cmd, &value)
		if err != nil {
			fmt.Printf("failed to scan standard input: %s\n", err.Error())
		}
		if n == 2 {
			switch cmd {
			case "ls":
				if err := ls(value); err != nil {
					fmt.Printf("failed to ls: %s\n", err.Error())
				}
			case "touch":
				if err := touch(value); err != nil {
					fmt.Printf("failed to touch: %s\n", err.Error())
				}
			case "cat":
				if err := cat(value); err != nil {
					fmt.Printf("failed to cat: %s\n", err.Error())
				}
			case "cd":
				if err := cd(value); err != nil {
					fmt.Printf("failed to cd: %s\n", err.Error())
				}
			default:
				fmt.Printf("unknown cmd: %s\n", cmd)
			}
		}
	}
}

func ls(path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}
	for _, entry := range entries {
		fmt.Printf("%v\n", entry.Name())
	}
	return nil
}

func touch(path string) error {
	_, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create file %w", err)
	}
	fmt.Printf("file created: %s\n", path)
	return nil
}

func cat(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file(%s): %w", path, err)
	}
	fmt.Println(string(file))
	return nil
}

func cd(path string) error {
	return os.Chdir(path)
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

func chHomeDir() (string, error) {
	homeDir := os.Getenv("HOME_DIR")
	if homeDir == "" {
		homeDir = "."
	}
	if err := os.Chdir(homeDir); err != nil {
		return "", fmt.Errorf("failed to change directory: %w", err)
	}
	return filepath.Abs(homeDir)
}
