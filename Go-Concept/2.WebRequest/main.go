package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://chaicode.com"

func main() {
	fmt.Println("Hello World")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Printf("Response: %T\n", res)
	DataBtypes, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(DataBtypes))
}
