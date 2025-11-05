package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "hello     world     ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "pretty big command to test \"asdfbasdfb\"",
			expected: []string{"pretty", "big", "command", "to", "test", "\"asdfbasdfb\""},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf(
				"cleanInput(\"%v\"): \ngot = %v\nexpected = %v",
				c.input, actual, c.expected,
			)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("fail")
			}
		}
	}
}
