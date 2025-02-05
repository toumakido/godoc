package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var length int64

func main() {
	var input string
	n, _ := fmt.Scanln(&input)
	length = int64(len(input))
	if n > 0 {
		file1, _ := os.OpenFile("stdio/file1.txt", os.O_RDWR|os.O_TRUNC, 0644)
		file2, _ := os.OpenFile("stdio/file2.txt", os.O_RDWR|os.O_TRUNC, 0644)
		file3, _ := os.OpenFile("stdio/file3.txt", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)
		defer file1.Close()
		defer file2.Close()

		myFile1 := &myFile{File: file1}
		myFile2 := &myFile{File: file2}
		myFile3 := &myFile{File: file3}

		if _, err := myFile1.Write([]byte(input)); err != nil {
			fmt.Printf("write err: %s\n", err.Error())
		}
		if _, err := myFile2.Write([]byte("ohayou")); err != nil {
			fmt.Printf("write err: %s\n", err.Error())
		}

		r := io.MultiReader(myFile1, myFile2)
		copy(myFile3, r)

		// content1 := make([]byte, 1028)
		// content2 := make([]byte, 1028)
		// myFile1.Read(content1)
		// myFile2.Read(content2)
		// fmt.Printf("content1: %s\n", string(content1))
		// fmt.Printf("content2: %s\n", string(content2))
	}
}

type myFile struct {
	File *os.File
}

func (m *myFile) Write(b []byte) (n int, err error) {
	fmt.Println(string(b))
	time.Sleep(1 * time.Second)
	n, err = m.File.Write(b)
	return
}

func (m *myFile) Read(b []byte) (n int, err error) {
	_, err = m.File.Seek(-length, io.SeekCurrent)
	n, err = m.File.Read(b)
	return
}

func copy(dst io.Writer, src io.Reader) (n int, err error) {
	b := make([]byte, 1028)
	n, err = src.Read(b)
	if err != nil {
		return
	}
	fmt.Printf("copy: %s\n", string(b))
	n, err = dst.Write(b[:n])
	return
}
