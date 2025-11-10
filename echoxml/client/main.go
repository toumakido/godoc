package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/toumakido/godoc/echoxml/def"
)

func main() {
	var res def.Response
	httpRes, err := http.Post("http://localhost:8000", "text/xml", nil)
	if err != nil {
		log.Fatalf("post err: %s", err.Error())
	}
	defer httpRes.Body.Close()
	if errUnmarshal := xml.NewDecoder(httpRes.Body).Decode(&res); errUnmarshal != nil {
		body, _ := io.ReadAll(httpRes.Body)
		fmt.Println(errUnmarshal)
		log.Fatalf("xml err: %s", body)
	}
	if res.IsError() {
		fmt.Println("OK")
		fmt.Println(res.Error())
	}
}
