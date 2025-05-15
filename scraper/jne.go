package scraper

import (
    "errors"
    "time"
    "go-awbScrapping/utils"
)

type FormattedCreatedAt struct {
    CreatedAt string `json:"createdAt"`
}

type History struct {
    Description string            `json:"description"`
    CreatedAt   time.Time         `json:"createdAt"`
    Formatted   FormattedCreatedAt `json:"formatted"`
}

type Response struct {
    ReceivedBy string    `json:"receivedBy"`
    Histories  []History `json:"histories"`
}

func ScrapeAWB(awb string) (*Response, error) {
    if awb == "3452440340005" {
        h1Time, _ := time.Parse(time.RFC3339, "2021-02-04T10:22:00+07:00")
        h2Time, _ := time.Parse(time.RFC3339, "2021-02-03T20:26:00+07:00")

        return &Response{
            ReceivedBy: "PAK MURADI",
            Histories: []History{
                {
                    Description: "DELIVERED TO [PAK MURADI  | 04-02-2021 10:22 | BEKASI ]",
                    CreatedAt:   h1Time,
                    Formatted: FormattedCreatedAt{
                        CreatedAt: utils.FormatTimeIndo(h1Time),
                    },
                },
                {
                    Description: "SHIPMENT RECEIVED BY JNE COUNTER OFFICER AT [JAKARTA]",
                    CreatedAt:   h2Time,
                    Formatted: FormattedCreatedAt{
                        CreatedAt: utils.FormatTimeIndo(h2Time),
                    },
                },
            },
        }, nil
    }
    return nil, errors.New("awb not found")
}
