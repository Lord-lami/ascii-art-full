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
	test.want = "\n\n\n\n\n\n\n\x1b[7A\x1b[33m _______  \x1b[B\x1b[10D|__   __| \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   |_|    \x1b[B\x1b[10D          \x1b[B\x1b[10D          \x1b[7A\x1b[0m       \x1b[B\x1b[7D       \x1b[B\x1b[7D  ___  \x1b[B\x1b[7D / _ \\ \x1b[B\x1b[7D|  __/ \x1b[B\x1b[7D \\___| \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A      \x1b[B\x1b[6D      \x1b[B\x1b[6D ___  \x1b[B\x1b[6D/ __| \x1b[B\x1b[6D\\__ \\ \x1b[B\x1b[6D|___/ \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[7A _    \x1b[B\x1b[6D| |   \x1b[B\x1b[6D| |_  \x1b[B\x1b[6D| __| \x1b[B\x1b[6D\\ |_  \x1b[B\x1b[6D \\__| \x1b[B\x1b[6D      \x1b[B\x1b[6D      "
	tests = append(tests, test)

	// 2nd test parameters - Newline
	test.testParameters.text = "Test\nTicket\nTraitor"
	test.testParameters.subStr = "t\n"
	test.testParameters.brush = "\033[33m"
	test.testParameters.alignment = "left"
	test.want = "\n\n\n\n\n\n\n\x1b[7A _______  \x1b[B\x1b[10D|__   __| \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   |_|    \x1b[B\x1b[10D          \x1b[B\x1b[10D          \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D  ___  \x1b[B\x1b[7D / _ \\ \x1b[B\x1b[7D|  __/ \x1b[B\x1b[7D \\___| \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A      \x1b[B\x1b[6D      \x1b[B\x1b[6D ___  \x1b[B\x1b[6D/ __| \x1b[B\x1b[6D\\__ \\ \x1b[B\x1b[6D|___/ \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[7A\x1b[33m _    \x1b[B\x1b[6D| |   \x1b[B\x1b[6D| |_  \x1b[B\x1b[6D| __| \x1b[B\x1b[6D\\ |_  \x1b[B\x1b[6D \\__| \x1b[B\x1b[6D      \x1b[B\x1b[6D      \n\n\n\n\n\n\n\n\x1b[7A\x1b[0m _______  \x1b[B\x1b[10D|__   __| \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   |_|    \x1b[B\x1b[10D          \x1b[B\x1b[10D          \x1b[7A _  \x1b[B\x1b[4D(_) \x1b[B\x1b[4D _  \x1b[B\x1b[4D| | \x1b[B\x1b[4D| | \x1b[B\x1b[4D|_| \x1b[B\x1b[4D    \x1b[B\x1b[4D    \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D  ___  \x1b[B\x1b[7D / __| \x1b[B\x1b[7D| (__  \x1b[B\x1b[7D \\___| \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A       \x1b[B\x1b[7D _     \x1b[B\x1b[7D| | _  \x1b[B\x1b[7D| |/ / \x1b[B\x1b[7D|   <  \x1b[B\x1b[7D|_|\\_\\ \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D  ___  \x1b[B\x1b[7D / _ \\ \x1b[B\x1b[7D|  __/ \x1b[B\x1b[7D \\___| \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A\x1b[33m _    \x1b[B\x1b[6D| |   \x1b[B\x1b[6D| |_  \x1b[B\x1b[6D| __| \x1b[B\x1b[6D\\ |_  \x1b[B\x1b[6D \\__| \x1b[B\x1b[6D      \x1b[B\x1b[6D      \n\n\n\n\n\n\n\n\x1b[7A\x1b[0m _______  \x1b[B\x1b[10D|__   __| \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   |_|    \x1b[B\x1b[10D          \x1b[B\x1b[10D          \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D _ __  \x1b[B\x1b[7D| '__| \x1b[B\x1b[7D| |    \x1b[B\x1b[7D|_|    \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A        \x1b[B\x1b[8D        \x1b[B\x1b[8D  __ _  \x1b[B\x1b[8D / _` | \x1b[B\x1b[8D| (_| | \x1b[B\x1b[8D \\__,_| \x1b[B\x1b[8D        \x1b[B\x1b[8D        \x1b[7A _  \x1b[B\x1b[4D(_) \x1b[B\x1b[4D _  \x1b[B\x1b[4D| | \x1b[B\x1b[4D| | \x1b[B\x1b[4D|_| \x1b[B\x1b[4D    \x1b[B\x1b[4D    \x1b[7A _    \x1b[B\x1b[6D| |   \x1b[B\x1b[6D| |_  \x1b[B\x1b[6D| __| \x1b[B\x1b[6D\\ |_  \x1b[B\x1b[6D \\__| \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[7A        \x1b[B\x1b[8D        \x1b[B\x1b[8D  ___   \x1b[B\x1b[8D / _ \\  \x1b[B\x1b[8D| (_) | \x1b[B\x1b[8D \\___/  \x1b[B\x1b[8D        \x1b[B\x1b[8D        \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D _ __  \x1b[B\x1b[7D| '__| \x1b[B\x1b[7D| |    \x1b[B\x1b[7D|_|    \x1b[B\x1b[7D       \x1b[B\x1b[7D       "
	tests = append(tests, test)

	// 3rd test parameters - Wrapping
	test.testParameters.text = "TestTicketTraitoring Heart"
	test.testParameters.subStr = "t\n"
	test.testParameters.brush = "\033[33m"
	test.testParameters.alignment = "left"
	test.want = "\n\n\n\n\n\n\n\x1b[7A _______  \x1b[B\x1b[10D|__   __| \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   |_|    \x1b[B\x1b[10D          \x1b[B\x1b[10D          \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D  ___  \x1b[B\x1b[7D / _ \\ \x1b[B\x1b[7D|  __/ \x1b[B\x1b[7D \\___| \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A      \x1b[B\x1b[6D      \x1b[B\x1b[6D ___  \x1b[B\x1b[6D/ __| \x1b[B\x1b[6D\\__ \\ \x1b[B\x1b[6D|___/ \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[7A _    \x1b[B\x1b[6D| |   \x1b[B\x1b[6D| |_  \x1b[B\x1b[6D| __| \x1b[B\x1b[6D\\ |_  \x1b[B\x1b[6D \\__| \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[7A _______  \x1b[B\x1b[10D|__   __| \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   |_|    \x1b[B\x1b[10D          \x1b[B\x1b[10D          \x1b[7A _  \x1b[B\x1b[4D(_) \x1b[B\x1b[4D _  \x1b[B\x1b[4D| | \x1b[B\x1b[4D| | \x1b[B\x1b[4D|_| \x1b[B\x1b[4D    \x1b[B\x1b[4D    \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D  ___  \x1b[B\x1b[7D / __| \x1b[B\x1b[7D| (__  \x1b[B\x1b[7D \\___| \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A       \x1b[B\x1b[7D _     \x1b[B\x1b[7D| | _  \x1b[B\x1b[7D| |/ / \x1b[B\x1b[7D|   <  \x1b[B\x1b[7D|_|\\_\\ \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D  ___  \x1b[B\x1b[7D / _ \\ \x1b[B\x1b[7D|  __/ \x1b[B\x1b[7D \\___| \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A _    \x1b[B\x1b[6D| |   \x1b[B\x1b[6D| |_  \x1b[B\x1b[6D| __| \x1b[B\x1b[6D\\ |_  \x1b[B\x1b[6D \\__| \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[7A _______  \x1b[B\x1b[10D|__   __| \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   | |    \x1b[B\x1b[10D   |_|    \x1b[B\x1b[10D          \x1b[B\x1b[10D          \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D _ __  \x1b[B\x1b[7D| '__| \x1b[B\x1b[7D| |    \x1b[B\x1b[7D|_|    \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A        \x1b[B\x1b[8D        \x1b[B\x1b[8D  __ _  \x1b[B\x1b[8D / _` | \x1b[B\x1b[8D| (_| | \x1b[B\x1b[8D \\__,_| \x1b[B\x1b[8D        \x1b[B\x1b[8D        \x1b[7A _  \x1b[B\x1b[4D(_) \x1b[B\x1b[4D _  \x1b[B\x1b[4D| | \x1b[B\x1b[4D| | \x1b[B\x1b[4D|_| \x1b[B\x1b[4D    \x1b[B\x1b[4D    \x1b[7A _    \x1b[B\x1b[6D| |   \x1b[B\x1b[6D| |_  \x1b[B\x1b[6D| __| \x1b[B\x1b[6D\\ |_  \x1b[B\x1b[6D \\__| \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[7A        \x1b[B\x1b[8D        \x1b[B\x1b[8D  ___   \x1b[B\x1b[8D / _ \\  \x1b[B\x1b[8D| (_) | \x1b[B\x1b[8D \\___/  \x1b[B\x1b[8D        \x1b[B\x1b[8D        \n\n\n\n\n\n\n\x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D _ __  \x1b[B\x1b[7D| '__| \x1b[B\x1b[7D| |    \x1b[B\x1b[7D|_|    \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A _  \x1b[B\x1b[4D(_) \x1b[B\x1b[4D _  \x1b[B\x1b[4D| | \x1b[B\x1b[4D| | \x1b[B\x1b[4D|_| \x1b[B\x1b[4D    \x1b[B\x1b[4D    \x1b[7A        \x1b[B\x1b[8D        \x1b[B\x1b[8D _ __   \x1b[B\x1b[8D| '_ \\  \x1b[B\x1b[8D| | | | \x1b[B\x1b[8D|_| |_| \x1b[B\x1b[8D        \x1b[B\x1b[8D        \x1b[7A        \x1b[B\x1b[8D        \x1b[B\x1b[8D  __ _  \x1b[B\x1b[8D / _` | \x1b[B\x1b[8D| (_| | \x1b[B\x1b[8D \\__, | \x1b[B\x1b[8D  __/ | \x1b[B\x1b[8D |___/  \x1b[7A      \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[B\x1b[6D      \x1b[7A _    _  \x1b[B\x1b[9D| |  | | \x1b[B\x1b[9D| |__| | \x1b[B\x1b[9D|  __  | \x1b[B\x1b[9D| |  | | \x1b[B\x1b[9D|_|  |_| \x1b[B\x1b[9D         \x1b[B\x1b[9D         \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D  ___  \x1b[B\x1b[7D / _ \\ \x1b[B\x1b[7D|  __/ \x1b[B\x1b[7D \\___| \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A        \x1b[B\x1b[8D        \x1b[B\x1b[8D  __ _  \x1b[B\x1b[8D / _` | \x1b[B\x1b[8D| (_| | \x1b[B\x1b[8D \\__,_| \x1b[B\x1b[8D        \x1b[B\x1b[8D        \x1b[7A       \x1b[B\x1b[7D       \x1b[B\x1b[7D _ __  \x1b[B\x1b[7D| '__| \x1b[B\x1b[7D| |    \x1b[B\x1b[7D|_|    \x1b[B\x1b[7D       \x1b[B\x1b[7D       \x1b[7A _    \x1b[B\x1b[6D| |   \x1b[B\x1b[6D| |_  \x1b[B\x1b[6D| __| \x1b[B\x1b[6D\\ |_  \x1b[B\x1b[6D \\__| \x1b[B\x1b[6D      \x1b[B\x1b[6D      "
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
