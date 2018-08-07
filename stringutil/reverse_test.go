package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		actual, expect string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases {
		got := Reverse(c.actual)
		if got != c.expect {
			t.Errorf("Reverse(%q) == %q, expect %q", c.actual, got, c.expect)
		}
	}
}
