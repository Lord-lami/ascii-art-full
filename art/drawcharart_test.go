package art

import (
	"asciiart/banners"
	"os"
	"path/filepath"
	"testing"
)

func TestDrawCharArt(t *testing.T) {
	const functionName = "drawCharArt"
	type record struct {
		testParameters struct {
			char byte
		}
		want string
	}
	tests := []record{}
	test := record{}

	// 1st test parameters
	test.testParameters.char = 'A'
	test.want = `           
    /\     
   /  \    
  / /\ \   
 / ____ \  
/_/    \_\ 
           
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
		result, _ := drawCharArt(
			test.testParameters.char,
		)
		if string(result) != test.want {
			t.Errorf("%v(%v) = %q want %q",
				functionName, test.testParameters, result, test.want)
		} else {
			t.Logf("Passed: %v(%v) = %q",
				functionName, test.testParameters, result)
		}
	}
}
