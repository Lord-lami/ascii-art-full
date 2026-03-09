package art

import (
	"asciiart/banners"
	"os"
	"path/filepath"
	"testing"
)

func TestFetchCharRowArtConc(t *testing.T) {
	const functionName = "FetchCharRowArtConc"
	type record struct {
		testParameters struct {
			char     byte
			position [2]int
			ch       chan charRowArt
		}
		want string
	}
	tests := []record{}
	test := record{}

	// 1st test parameters
	test.testParameters.char = 'A'
	test.testParameters.position = [2]int{0, 1}
	test.testParameters.ch = make(chan charRowArt)
	test.want = `    /\     `
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
		go fetchCharRowArtConc(
			test.testParameters.char,
			test.testParameters.position,
			test.testParameters.ch,
		)
		result := <-test.testParameters.ch
		close(test.testParameters.ch)
		resultStr := string(result.RowArt)
		if resultStr != test.want {
			t.Errorf("%v(%v) = %q want %q",
				functionName, test.testParameters, resultStr, test.want)
		} else {
			t.Logf("Passed: %v(%v) = %q",
				functionName, test.testParameters, result.RowArt)
		}
	}
}
