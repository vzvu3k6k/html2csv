package html2csv

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_getCSVFromHTML(t *testing.T) {
	type testcase struct {
		name string
		html string
		csv  string
	}

	// Build testcases from files in testdata directory.
	// The testdata directory contains pairs of HTML and CSV files.
	// The HTML file serves as the input for the getCSVFromHTML function, and the CSV file is the expected output.
	testcases := func() []*testcase {
		testcases := []*testcase{}

		htmlFilepaths, err := filepath.Glob("testdata/*.html")
		if err != nil {
			t.Fatal(err)
		}
		for _, htmlFilepath := range htmlFilepaths {
			html, err := os.ReadFile(htmlFilepath)
			if err != nil {
				t.Fatal(err)
			}

			basename := htmlFilepath[:len(htmlFilepath)-len(".html")]
			csvFilepath := basename + ".csv"
			csv, err := os.ReadFile(csvFilepath)
			if err != nil {
				t.Fatal(err)
			}

			testcases = append(testcases, &testcase{
				name: basename,
				html: string(html),
				csv:  string(csv),
			})
		}
		return testcases
	}

	for _, tc := range testcases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			gotCSV, err := getCSVFromHTML(strings.NewReader(tc.html))
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tc.csv, gotCSV); diff != "" {
				t.Errorf("(-want +got):\n%s", diff)
			}
		})
	}
}
