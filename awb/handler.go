package awb

import (
    "encoding/json"
    "net/http"
    "strings"
    "go-awbScrapping/scraper"
)

type Status struct {
    Code    string `json:"code"`
    Message string `json:"message"`
}

type Response struct {
    Status Status      `json:"status"`
    Data   interface{} `json:"data"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
    awbNumber := strings.TrimPrefix(r.URL.Path, "/awb/")
    w.Header().Set("Content-Type", "application/json")

    data, err := scraper.ScrapeAWB(awbNumber)
    if err != nil {
        json.NewEncoder(w).Encode(Response{
            Status: Status{
                Code:    "060102",
                Message: "AWB number not found",
            },
            Data: nil,
        })
        return
    }

    json.NewEncoder(w).Encode(Response{
        Status: Status{
            Code:    "060101",
            Message: "Delivery tracking detail fetched successfully",
        },
        Data: data,
    })
}
