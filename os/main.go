package main

import (
	"fmt"
	"os"
)

func main() {
	// 標準入力から相対パスを受け取って一覧表示(ディレクトリ想定で)
	var dir string
	for {
		n, err := fmt.Scanln(&dir)
		if err != nil {
			fmt.Printf("failed to scan standard input: %s", err.Error())
		}
		if n > 0 {
			entries, err := os.ReadDir(dir)
			if err != nil {
				fmt.Printf("failed to read directory: %s", err.Error())
			}
			for _, entry := range entries {
				fmt.Printf("%v\n", entry.Name())
			}
		}
	}
}
