package art

import (
	"asciiart/banners"
	"os"
	"path/filepath"
	"testing"
)

func TestDrawPaintWrapAlignTextArt(t *testing.T) {
	const functionName = "DrawPaintWrapAlignTextArt"
	type record struct {
		testParameters struct {
			text      string
			subStr    string
			brush     string
			alignment string
		}
		want string
	}
	tests := []record{}
	test := record{}

	// 1st test parameters - Plain
	test.testParameters.text = "Test"
	test.testParameters.subStr = "T"
	test.testParameters.brush = "\033[33m"
	test.testParameters.alignment = "left"
	test.want = "\x1b[33m _______  \x1b[0m              _    \n\x1b[33m|__   __| \x1b[0m             | |   \n\x1b[33m   | |    \x1b[0m  ___   ___  | |_  \n\x1b[33m   | |    \x1b[0m / _ \\ / __| | __| \n\x1b[33m   | |    \x1b[0m|  __/ \\__ \\ \\ |_  \n\x1b[33m   |_|    \x1b[0m \\___| |___/  \\__| \n\x1b[33m          \x1b[0m                   \n\x1b[33m          \x1b[0m                   \n"
	tests = append(tests, test)

	// 2nd test parameters - Newline
	test.testParameters.text = "Test\nTicket\nTraitor"
	test.testParameters.subStr = "t\n"
	test.testParameters.brush = "\033[33m"
	test.testParameters.alignment = "left"
	test.want = " _______               \x1b[33m _    \x1b[0m\n|__   __|              \x1b[33m| |   \x1b[0m\n   | |      ___   ___  \x1b[33m| |_  \x1b[0m\n   | |     / _ \\ / __| \x1b[33m| __| \x1b[0m\n   | |    |  __/ \\__ \\ \x1b[33m\\ |_  \x1b[0m\n   |_|     \\___| |___/ \x1b[33m \\__| \x1b[0m\n                       \x1b[33m      \x1b[0m\n                       \x1b[33m      \x1b[0m\n\x1b[33m\x1b[0m _______   _                       \x1b[33m _    \x1b[0m\n\x1b[33m\x1b[0m|__   __| (_)         _            \x1b[33m| |   \x1b[0m\n\x1b[33m\x1b[0m   | |     _    ___  | | _    ___  \x1b[33m| |_  \x1b[0m\n\x1b[33m\x1b[0m   | |    | |  / __| | |/ /  / _ \\ \x1b[33m| __| \x1b[0m\n\x1b[33m\x1b[0m   | |    | | | (__  |   <  |  __/ \x1b[33m\\ |_  \x1b[0m\n\x1b[33m\x1b[0m   |_|    |_|  \\___| |_|\\_\\  \\___| \x1b[33m \\__| \x1b[0m\n\x1b[33m\x1b[0m                                   \x1b[33m      \x1b[0m\n\x1b[33m\x1b[0m                                   \x1b[33m      \x1b[0m\n\x1b[33m\x1b[0m _______                  _   _                   \n\x1b[33m\x1b[0m|__   __|                (_) | |                  \n\x1b[33m\x1b[0m   | |     _ __    __ _   _  | |_    ___    _ __  \n\x1b[33m\x1b[0m   | |    | '__|  / _` | | | | __|  / _ \\  | '__| \n\x1b[33m\x1b[0m   | |    | |    | (_| | | | \\ |_  | (_) | | |    \n\x1b[33m\x1b[0m   |_|    |_|     \\__,_| |_|  \\__|  \\___/  |_|    \n\x1b[33m\x1b[0m                                                  \n\x1b[33m\x1b[0m                                                  \n"
	tests = append(tests, test)

	// 3rd test parameters - Wrapping
	test.testParameters.text = "TestTicketTraitoring Heart"
	test.testParameters.subStr = "t\n"
	test.testParameters.brush = "\033[33m"
	test.testParameters.alignment = "left"
	test.want = " _______                _     _______   _                        _     _______                  _   _                    _  \n|__   __|              | |   |__   __| (_)         _            | |   |__   __|                (_) | |                  (_) \n   | |      ___   ___  | |_     | |     _    ___  | | _    ___  | |_     | |     _ __    __ _   _  | |_    ___    _ __   _  \n   | |     / _ \\ / __| | __|    | |    | |  / __| | |/ /  / _ \\ | __|    | |    | '__|  / _` | | | | __|  / _ \\  | '__| | | \n   | |    |  __/ \\__ \\ \\ |_     | |    | | | (__  |   <  |  __/ \\ |_     | |    | |    | (_| | | | \\ |_  | (_) | | |    | | \n   |_|     \\___| |___/  \\__|    |_|    |_|  \\___| |_|\\_\\  \\___|  \\__|    |_|    |_|     \\__,_| |_|  \\__|  \\___/  |_|    |_| \n                                                                                                                            \n                                                                                                                            \n                       _    _                         _    \n                      | |  | |                       | |   \n _ __     __ _        | |__| |   ___    __ _   _ __  | |_  \n| '_ \\   / _` |       |  __  |  / _ \\  / _` | | '__| | __| \n| | | | | (_| |       | |  | | |  __/ | (_| | | |    \\ |_  \n|_| |_|  \\__, |       |_|  |_|  \\___|  \\__,_| |_|     \\__| \n          __/ |                                            \n         |___/                                             \n"
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
		result := DrawPaintWrapAlignTextArt(
			test.testParameters.text,
			test.testParameters.subStr,
			test.testParameters.brush,
			test.testParameters.alignment,
		)
		if result != test.want {
			t.Errorf("%v(%#v) = \n%q\n\nwant\n\n%q",
				functionName, test.testParameters, result, test.want)
		} else {
			t.Logf("Passed: %v(%#v) =\n%s\n",
				functionName, test.testParameters, result)
		}
	}
}
