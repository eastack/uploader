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

func uploadFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPatch:
		fmt.Println("Receive a file...")

		// parse multipart form
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			fmt.Println("Parse multipart form failed")
			fmt.Println(err)
			return
		}

		// get chunk
		file, header, err := r.FormFile("chunk")
		if err != nil {
			fmt.Println("Upload failed")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Upload file size: %+v\n", header.Header)

		// create temp file
		tempFile, err := ioutil.TempFile("temp-chunk", "chunk-*")
		if err != nil {
			fmt.Println("Create temp file failed")
			fmt.Println(err)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
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

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}
