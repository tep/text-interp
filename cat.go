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

type cat int

const (
	cNorm cat = iota
	cBegin
	cEnd
	cEscNorm
	cEscBegin
	cEscEnd
)

func (c cat) String() string {
	switch c {
	case cBegin:
		return "beg"
	case cEnd:
		return "end"
	case cEscNorm:
		return "NRM"
	case cEscBegin:
		return "BEG"
	case cEscEnd:
		return "END"
	default:
		return "nrm"
	}
}

func (c cat) escaped(e bool) cat {
	if !e {
		return c
	}

	switch c {
	case cBegin:
		return cEscBegin
	case cEnd:
		return cEscEnd
	default:
		return cEscNorm
	}
}
