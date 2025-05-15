package main

import (
    "log"
    "net/http"
    "go-awbScrapping/awb"
)

func main() {
    http.HandleFunc("/awb/", awb.Handler)
    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
