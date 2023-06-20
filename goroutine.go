package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	go makeReqest("https://qiita.com/ohbashunsuke/items/e7c673db606a6dced8a6")

	go count("goroutine")
	count("main")

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}

func count(label string) {
	for i := 1; i <= 5; i++ {
		for n := 1; n <= 10000000; n++ {

		}
		fmt.Println(label, i)
		// time.Sleep(time.Microsecond * 500)
	}
}

func makeReqest(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
