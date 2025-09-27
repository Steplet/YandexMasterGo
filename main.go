package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func main() {
	for {
		resp, err := http.Get("http://srv.msk01.gigacorp.local/_stats")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		paramsList := strings.Split(string(body), ",")
		fmt.Println(paramsList)

		time.Sleep(10 * time.Millisecond)
	}
}
