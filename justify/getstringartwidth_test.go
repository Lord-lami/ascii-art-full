package justify

import (
	"asciiart/banners"
	"os"
	"path/filepath"
	"testing"
)

func TestGetStringArtWidth(t *testing.T) {
	const functionName = "getStringArtWidth"
	type record struct {
		testParameters struct {
			text string
		}
		want int
	}
	tests := []record{}
	test := record{}

	// 1st test parameters
	test.testParameters.text = "Test"
	test.want = 29
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
		result := getStringArtWidth(
			test.testParameters.text,
		)
		if result != test.want {
			t.Errorf("%v(%v) = %d want %d",
				functionName, test.testParameters, result, test.want)
		} else {
			t.Logf("Passed: %v(%v) = %d",
				functionName, test.testParameters, result)
		}
	}
}
