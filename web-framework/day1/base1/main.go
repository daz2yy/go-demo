package main

// http 库使用：http.HandleFunc, http.ListenAndServe, http.ResponseWriter, http.Request
// fmt 库使用：fmt.Fprintf
// log 库使用：log.Println

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHander)
	http.HandleFunc("/header", headerHander)

	log.Println("Serving....")
	log.Fatal(http.ListenAndServe(":8899", nil))
}

func indexHander(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "URL.Path = %q", req.URL.Path)
}

func headerHander(rw http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(rw, "Header[%q] = [%q]", k, v)
	}
}
