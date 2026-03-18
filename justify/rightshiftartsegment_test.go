package justify

import (
	"testing"
)

func TestRightShiftArtSegment(t *testing.T) {
	const functionName = "rightShiftArtSegment"
	type record struct {
		testParameters struct {
			artSegment string
			shiftSize  int
		}
		want string
	}
	tests := []record{}
	test := record{}

	// 1st test parameter(s)
	test.testParameters.artSegment = "Test"
	test.testParameters.shiftSize = 10
	test.want = "          Test"
	tests = append(tests, test)

	// 2nd test parameter(s)
	test.testParameters.artSegment = "\n\n\n\n\n\n\nTest"
	test.testParameters.shiftSize = 5
	test.want = "\n\n\n\n\n\n\n     Test"
	tests = append(tests, test)

	for _, test := range tests {
		result := rightShiftArtSegment(
			test.testParameters.artSegment,
			test.testParameters.shiftSize,
		)
		if result != test.want {
			t.Errorf("%v(%v) = %q\nwant\n%q",
				functionName, test.testParameters, result, test.want)
		} else {
			t.Logf("Passed: %v(%v) = %q",
				functionName, test.testParameters, result)
		}
	}
}
