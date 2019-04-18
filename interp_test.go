package interp

import (
	"fmt"
	"testing"
)

func TestInterpolator(t *testing.T) {
	cases := []*intTestcase{
		{"foo${one}bar", "fooabcbar"},
		{"foo${${one}bar}", "fooshizzl"},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("T%d", i), tc.test)
	}
}

type intTestcase struct {
	in   string
	want string
}

func (tc *intTestcase) test(t *testing.T) {
	i := New(mkVars())

	if got, err := i.Interpolate(tc.in); err != nil || got != tc.want {
		t.Errorf("Interpolate(%q) == (%q, %v); wanted(%q, %v)", tc.in, got, err, tc.want, nil)
	}
}

func mkVars() testResolver {
	return testResolver(map[string]string{
		"one":    "abc",
		"abcbar": "shizzl",
	})
}

type testResolver map[string]string

func (tg testResolver) Resolve(s string) (Value, error) {
	if v, ok := tg[s]; ok {
		return v, nil
	}
	return "", fmt.Errorf("unknown var: %s", s)
}
