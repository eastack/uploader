package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	setupRoutes()
}

func download(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPatch:
		fmt.Println("Receive a file...")

		// create temp file
		tempFile, err := ioutil.TempFile("temp-chunk", "chunk-*")
		if err != nil {
			fmt.Println("Create temp file failed")
			fmt.Println(err)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Read file failed")
			fmt.Println(err)
			return
		}
		tempFile.Write(fileBytes)

		fmt.Fprintf(w, "Successfully upload file\n")
	default:
		fmt.Fprintf(w, "Sorry, only support patch")
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPatch:
		fmt.Println("Receive a file...")
		fmt.Println(r.Header)

		tempFile, err := ioutil.TempFile("", "chunk-*")
		if err != nil {
			panic(err.Error())
		}
		defer tempFile.Close()

		chunk, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err.Error())
		}
		tempFile.Write(chunk)

		w.WriteHeader(201)
		fmt.Fprintf(w, "Successfully upload file\n")
	default:
		fmt.Fprintf(w, "Sorry, only support patch")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

func setupRoutes() {
	http.HandleFunc("/upload", upload)
	http.HandleFunc("/download", download)
	http.HandleFunc("/hi", index)
	http.ListenAndServe(":8080", nil)
}
