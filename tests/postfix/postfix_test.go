package postfix_test

import (
	"calgo/internals/postfix"
	"testing"
)

func TestPostfix(t *testing.T) {
	testcases := []struct {
		input    string
		expected int
	}{
		{
			input:    "1+2",
			expected: 3,
		},
		{
			input:    "2-1",
			expected: 1,
		},
	}

	for _, testcase := range testcases {
		ret := postfix.Postfix(testcase.input)
		if ret != testcase.expected {
			t.Fatalf("expected %d, got %d, some random bs lmfaooo, u suck", testcase.expected, ret)
		}
	}
}
