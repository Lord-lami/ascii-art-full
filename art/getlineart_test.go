package art

import (
	"asciiart/banners"
	"os"
	"path/filepath"
	"testing"
)

func TestGetLineArt(t *testing.T) {
	const functionName = "GetLineArt"
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
	test.want =
		`                     
                     
__  __  _   _   ____ 
\ \/ / | | | | |_  / 
 >  <  | |_| |  / /  
/_/\_\  \__, | /___| 
        __/ /        
       |___/         
`
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
		result := GetLineArt(
			&test.testParameters.line,
		)
		if result != test.want {
			t.Errorf("%v(%v) = \n%q\n want \n\"%s\"",
				functionName, test.testParameters, result, test.want)
		} else {
			t.Logf("Passed: %v(%v) = \"%s\"",
				functionName, test.testParameters, result)
		}
	}
}
