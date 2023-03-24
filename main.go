package main

import (
	"fmt"
	utils "go-yt/core/utils"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/get", utils.DownloadHandler)

	fmt.Println("Starting GO-YT server on :7070...")
	log.Fatal(http.ListenAndServe(":7070", nil))
}
