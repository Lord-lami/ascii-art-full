package formatforfile

import (
	"testing"
)

func TestFormatForFile(t *testing.T) {
	const functionName = "FormatForFile"
	type record struct {
		testParameters struct {
			art string
		}
		want string
	}
	tests := []record{}
	test := record{}

	// 1st test parameters
	test.testParameters.art = "\n\n\n\n\n\n\n\x1b[7A\x1b[39m ____   \x1b[B\x1b[8D|  _ \\  \x1b[B\x1b[8D| |_) | \x1b[B\x1b[8D|  _ <  \x1b[B\x1b[8D| |_) | \x1b[B\x1b[8D|____/  \x1b[B\x1b[8D        \x1b[B\x1b[8D        \x1b[0m"
	test.want = ""
	tests = append(tests, test)

	for _, test := range tests {
		FormatForFile(test.testParameters.art)
	}
}
