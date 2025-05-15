package utils

import (
    "fmt"
    "time"
    "golang.org/x/text/language"
    "golang.org/x/text/message"
)

func FormatTimeIndo(t time.Time) string {
    p := message.NewPrinter(language.Indonesian)
    return fmt.Sprintf("%02d %s %d, %02d:%02d WIB",
        t.Day(),
        p.Sprintf("%s", t.Month().String()), 
        t.Year(),
        t.Hour(),
        t.Minute())
}


