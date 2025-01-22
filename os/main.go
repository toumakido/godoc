package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// 標準入力から相対パスを受け取って一覧表示(ディレクトリ想定で)
	dir := "../"
	entries, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("failed to change directory: %s", dir)
	}
	for _, entry := range entries {
		fmt.Printf("entry: %v\n", entry.Name())
	}
}
