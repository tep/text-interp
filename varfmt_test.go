//
// Copyright 2019 Timothy E. Peoples
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to
// deal in the Software without restriction, including without limitation the
// rights to use, copy, modify, merge, publish, distribute, sublicense, and/or
// sell copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
// IN THE SOFTWARE.
//

package interp

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestVarFormat(t *testing.T) {

	cases := []*vfmtTestcase{
		vfmtTC(``, "", 0, 0),
		vfmtTC(`foo`, "", 0, 0),
		vfmtTC(`${foo}`, "foo", 0, 5),
		vfmtTC(`${foo}bar`, "foo", 0, 5),
		vfmtTC(`bar${foo}`, "foo", 3, 8),
		vfmtTC(`${foo`, "", 0, 0),

		vfmtTC(`abc ${one${two}foo} def`, "two", 9, 14),
		vfmtTC(`abc \${one${two}foo} def`, "two", 10, 15),
		vfmtTC(`abc ${one\${two}foo} def`, `one${two}foo`, 4, 19),
		vfmtTC(`abc ${one${two\}foo} def`, "two}foo", 9, 19),
		vfmtTC(`abc ${one${two}foo\} def`, "two", 9, 14),

		vfmtTC(`abc ${one\${two}foo\} def`, "", 0, 0),
		vfmtTC(`abc ${one\${two}foo\} def}`, "one${two}foo} def", 4, 25),

		vfmtTC(`abc ${one${two${three}}} def`, "three", 14, 21),
		vfmtTC(`abc ${one${two\${three}}} def`, "two${three}", 9, 23),
		vfmtTC(`abc ${one\${two${three}}} def`, "three", 15, 22),
		vfmtTC(`abc ${one${two${three\}}} def`, "three}", 14, 23),
		vfmtTC(`abc ${one${two\${three\}}} def`, "two${three}}", 9, 25),
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("T%d", i), tc.test)
	}
}

func (tc *vfmtTestcase) test(t *testing.T) {
	if testing.Verbose() {
		debugf = fmt.Printf

		fmt.Printf("\n=====[  %-12s  ]======================\n", t.Name())
		fmt.Println("              1         2         3         4")
		fmt.Println("    0 2 4 6 8 0 2 4 6 8 0 2 4 6 8 0 2 4 6 8 0")
		fmt.Print("    ")
		os.Stdout.Write([]byte(tc.in))
		fmt.Print("\n\n")
	}

	rs := stdFormat.replString(tc.in)

	if got := rs.next(); !reflect.DeepEqual(got, tc.want) {
		t.Errorf("find(%q) == (%#v); wanted (%#v)", tc.in, got, tc.want)
	}
}

type vfmtTestcase struct {
	in   string
	want *token
}

func vfmtTC(str, want string, beg, end int) *vfmtTestcase {
	var tok *token
	if want != "" {
		tok = &token{want, str, beg, end}
	}
	return &vfmtTestcase{str, tok}
}
