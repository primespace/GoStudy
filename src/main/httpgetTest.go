package main

import "net/http"

func main() {

	resp, err := http.Get("http://www.naver.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

}
