package main

import "net/http"

func main() {

	network := http.NewServeMux()

	network.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Пирювет Мама!"))
	})

	http.ListenAndServe("localhost:8080", network)
}
