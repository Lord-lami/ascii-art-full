package art

import (
	"asciiart/banners"
	"os"
	"path/filepath"
	"testing"
)

func TestGetLineRowArt(t *testing.T) {
	const functionName = "getLineRowArt"
	type record struct {
		testParameters struct {
			line string
			y    byte
		}
		want string
	}
	tests := []record{}
	test := record{}

	// 1st test parameters
	test.testParameters.line = "BAC"
	test.testParameters.y = 1
	test.want = `|  _ \      /\      / ____| `
	tests = append(tests, test)

	// Setup
	path := filepath.Join("..", "banners", "standard.txt")
	var err error
	banners.BannerFile, err = os.Open(path)
	if err != nil {
		panic(err)
	}
	defer banners.BannerFile.Close()
	banners.SetBannerLineIndex()

	for _, test := range tests {
		result := getLineRowArt(
			&test.testParameters.line,
			test.testParameters.y,
		)
		if string(result) != test.want {
			t.Errorf("%v(%v) = \n%s\n want \n%q",
				functionName, test.testParameters, result, test.want)
		} else {
			t.Logf("Passed: %v(%v) = %s",
				functionName, test.testParameters, result)
		}
	}
}
