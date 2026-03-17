package art

import (
	"asciiart/banners"
	"os"
	"path/filepath"
	"testing"
)

func TestCommissionArtist(t *testing.T) {
	const functionName = "CommissionArtist"
	type record struct {
		testParameters struct {
			text  string
			start int
			stop  int
		}
		want string
	}
	tests := []record{}
	test := record{}

	// 1st test parameters
	test.testParameters.text = "Test\n\n\nExam"
	test.testParameters.start = 0
	test.testParameters.stop = len(test.testParameters.text)
	test.want = "\n\n\n\n\n\n\n\x1b[7A _______  \x1b[B\x1b[10D|__   __| \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   |_|    \x1b[B\x1b[10D          \x1b[B\x1b[10D          \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D  ___  \x1b[B\x1b[7D / _ \\ \x1b[B\x1b[7D|  __/ \x1b[B\x1b[7D \\___| \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A      \x1b[B\x1b[6D      \x1b[B\x1b[6D ___  \x1b[B\x1b[6D/ __| \x1b[B\x1b[6D\\__ \\ \x1b[B\x1b[6D|___/ \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[7A _    \x1b[B\x1b[6D| |   \x1b[B\x1b[6D| |_  \x1b[B\x1b[6D| __| \x1b[B\x1b[6D\\ |_  \x1b[B\x1b[6D \\__| \x1b[B\x1b[6D      \x1b[B\x1b[6D      \n\n\n\n\n\n\n\n\n\n\x1b[7A ______  \x1b[B\x1b[9D|  ____| \x1b[B\x1b[9D| |__    \x1b[B\x1b[9D|  __|   \x1b[B\x1b[9D| |____  \x1b[B\x1b[9D|______| \x1b[B\x1b[9D         \x1b[B\x1b[9D         \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D__  __ \x1b[B\x1b[7D\\ \\/ / \x1b[B\x1b[7D >  <  \x1b[B\x1b[7D/_/\\_\\ \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A        \x1b[B\x1b[8D        \x1b[B\x1b[8D  __ _  \x1b[B\x1b[8D / _` | \x1b[B\x1b[8D| (_| | \x1b[B\x1b[8D \\__,_| \x1b[B\x1b[8D        \x1b[B\x1b[8D        \x1b[7A            \x1b[B\x1b[12D            \x1b[B\x1b[12D _ __ ___   \x1b[B\x1b[12D| '_ ` _ \\  \x1b[B\x1b[12D| | | | | | \x1b[B\x1b[12D|_| |_| |_| \x1b[B\x1b[12D            \x1b[B\x1b[12D            "
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
		drawTestArt := CommissionArtist(
			test.testParameters.text,
		)
		result := drawTestArt(test.testParameters.start, test.testParameters.stop)
		if string(result) != test.want {
			t.Errorf("%v(%v) = %q want %q",
				functionName, test.testParameters, result, test.want)
		} else {
			t.Logf("Passed: %v(%v) = %q",
				functionName, test.testParameters, result)
		}
	}
}
