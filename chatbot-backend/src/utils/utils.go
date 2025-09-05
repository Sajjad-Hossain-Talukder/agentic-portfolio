package utils

import (
	"log"

	pdf "github.com/ledongthuc/pdf"
)

func ReadPDFText(path string) string {
    f, r, err := pdf.Open(path)
    if err != nil {
        log.Println("Error opening PDF:", err)
        return ""
    }
    defer f.Close()

    var text string
    totalPages := r.NumPage()
    for pageIndex := 1; pageIndex <= totalPages; pageIndex++ {
        page := r.Page(pageIndex)
        if page.V.IsNull() {
            continue
        }
        content, err := page.GetPlainText(nil)
        if err != nil {
            log.Println("Error reading page:", err)
            continue
        }
        text += content
    }

    return text
}
