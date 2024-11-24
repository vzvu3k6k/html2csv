package html2csv

import (
	"fmt"
	"os"
)

func Run() error {
	csv, err := getCSVFromHTML(os.Stdin)
	if err != nil {
		return err
	}
	fmt.Println(csv)
	return nil
}
