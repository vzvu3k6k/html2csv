package html2csv

import (
	"errors"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var ErrNoTableFound = errors.New("no table found")

func getCSVFromHTML(r io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}

	table := doc.Find("table")
	if table.Length() == 0 {
		return "", ErrNoTableFound
	}

	var b strings.Builder
	for _, tr := range table.Find("tr").EachIter() {
		tds := tr.Find("td")
		for i, td := range tds.EachIter() {
			b.WriteString(strings.TrimSpace(td.Text()))
			if i < tds.Length()-1 {
				b.WriteString(",")
			}
		}
		b.WriteString("\n")
	}

	return b.String(), nil
}
