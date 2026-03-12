package art

import (
	"asciiart/banners"
	"os"
	"path/filepath"
	"testing"
)

func TestGetLineArt(t *testing.T) {
	const functionName = "getLineArt"
	type record struct {
		testParameters struct {
			line string
		}
		want string
	}
	tests := []record{}
	test := record{}

	// 1st test parameters
	test.testParameters.line = "xyz"
	test.want = "       \n       \n__  __ \n\\ \\/ / \n >  <  \n/_/\\_\\ \n       \n       \n\x1b[8A\x1b[7C        \x1b[B\x1b[8D        \x1b[B\x1b[8D _   _  \x1b[B\x1b[8D| | | | \x1b[B\x1b[8D| |_| | \x1b[B\x1b[8D \\__, | \x1b[B\x1b[8D __/ /  \x1b[B\x1b[8D|___/   \n\x1b[8A\x1b[15C      \x1b[B\x1b[6D      \x1b[B\x1b[6D ____ \x1b[B\x1b[6D|_  / \x1b[B\x1b[6D / /  \x1b[B\x1b[6D/___| \x1b[B\x1b[6D      \x1b[B\x1b[6D      \n"
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
		first, rest := getLineArt(
			test.testParameters.line,
		)
		result := first + rest("")
		if result != test.want {
			t.Errorf("%v(%v) = \n\"%#v\"\n want \n\"%#v\"",
				functionName, test.testParameters, result, test.want)
		} else {
			t.Logf("Passed: %v(%v) = \n\"%s\"",
				functionName, test.testParameters, result)
		}
	}
}
