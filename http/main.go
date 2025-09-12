package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type LogWriter struct {
}

func main() {
	resp, err := http.Get("https://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	//bs := make([]byte, 99999)
	//
	//_, err = resp.Body.Read(bs)
	//
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(1)
	//}
	//
	//fmt.Println(string(bs))

	//io.Copy(os.Stdout, resp.Body)

	lw := LogWriter{}

	io.Copy(lw, resp.Body)

	os.Exit(0)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}

func (LogWriter) Write(bs []byte) (int, error) {
	bl := len(bs)

	fmt.Println("Write bytes: ", bl)
	fmt.Println(string(bs))

	return bl, nil
}
