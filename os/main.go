package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if err := setEnv(); err != nil {
		log.Fatalf("failed to set env: %s", err.Error())
	}
	if err := chHomeDir(); err != nil {
		log.Fatalf("failed to change home directory: %s", err.Error())
	}

	var cmd, value string
	for {
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
			}
		} else {
			fmt.Println("unknown input")
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
