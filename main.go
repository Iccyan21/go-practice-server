package main

import (
	"fmt"
	"net/http"
)

var count int

// go run main.go でmainの処理が開始
func main() {
	// URL指定処理
	http.HandleFunc("/", handler)
	http.HandleFunc("/get", countHandler)
	// Serverを立ち上げている
	http.ListenAndServe(":8080", nil)
}

// リクエストを受ける処理
func handler(w http.ResponseWriter, r *http.Request) {
	count++
	//数値型はそのままでは無理なので変換
	fmt.Fprintf(w, "<html><h1>Hello World</h1></html>")
}
func countHandler(w http.ResponseWriter, r *http.Request) {
	count++
	//数値型はそのままでは無理なので変換
	fmt.Fprintf(w, "count: %d", count)
}
