package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	go makeReqest("https://takahashi-pao.github.io/oretachi-omaetachi/member.html")

	go count("goroutine")
	count("main")

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

	GetRequest()
}

func count(label string) {
	for i := 1; i <= 5; i++ {
		for n := 1; n <= 10000000; n++ {

		}
		fmt.Println(label, i)
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

func GetRequest() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := "https://takahashi-pao.github.io/oretachi-omaetachi/member.html"

		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, "Error making request", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading response body", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, string(body)) // HTMLコンテンツをレスポンスとして返す
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
