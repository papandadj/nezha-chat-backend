package service

import (
	"testing"
)

func TestSum256(t *testing.T) {
	var tests = []struct {
		Arg1   string
		Result string
	}{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"1", "6b86b273ff34fce19d6b804eff5a3f5747ada4eaa22f1d49c01e52ddb7875b4b"},
		{"secret", "2bb80d537b1da3e38bd30361aa855686bde0eacd7162fef6a25fe97bf527a25b"},
	}

	for _, test := range tests {
		result := Sum256(test.Arg1)
		if result != test.Result {
			t.Errorf("Sum256  failed to input %s, expected %s , got %s", test.Arg1, test.Result, result)
		}
	}
}
