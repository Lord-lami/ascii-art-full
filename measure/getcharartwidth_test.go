package measure

import (
	"asciiart/banners"
	"os"
	"path/filepath"
	"testing"
)

func TestGetCharArtWidth(t *testing.T) {
	const functionName = "GetCharArtWidth"
	type record struct {
		testParameters struct {
			char byte
		}
		want int
	}
	tests := []record{}
	test := record{}

	// 1st test parameters
	test.testParameters.char = 'T'
	test.want = 10
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
		result := GetCharArtWidth(
			test.testParameters.char,
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
