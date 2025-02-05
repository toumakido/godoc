package main

import (
	"golang.org/x/sync/errgroup"

	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	dir := "netio/tmp"
	os.Mkdir(dir, 0750)

	urls := []string{
		"https://www.nta.go.jp/taxes/shiraberu/shinkoku/tebiki/2024/kisairei/pdf/0023010-154_05.pdf",
		"https://www.nta.go.jp/taxes/shiraberu/shinkoku/tebiki/2024/kisairei/pdf/0023010-154_05.pdf",
		"https://www.nta.go.jp/taxes/shiraberu/shinkoku/tebiki/2024/kisairei/pdf/0023010-154_05.pdf",
		"https://www.nta.go.jp/taxes/shiraberu/shinkoku/tebiki/2024/kisairei/pdf/0023010-154_05.pdf",
		"https://pkg.go.dev/os",
	}
	var eg errgroup.Group
	for i, url := range urls {
		eg.Go(func() error {
			return download(url, fmt.Sprintf("%s/%v", dir, i))
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Printf("failed to download: %s", err.Error())
	}
}

func download(url, name string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to get %s, err: %s", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		file, err := os.Create(name + getExtention(resp.Header))
		written, err := io.Copy(file, resp.Body)
		if err != nil {
			return err
		}
		fmt.Printf("file size %s: %v", name, written)
		return nil
	}
	return fmt.Errorf("http requst not sucessed :%v", resp.StatusCode)
}

func getExtention(header http.Header) string {
	switch header.Get("Content-Type") {
	case "application/pdf":
		return ".pdf"
	case "text/html; charset=utf-8":
		return ".html"
	}
	return ""
}
