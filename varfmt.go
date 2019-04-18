package interp

import "strings"

var stdFormat = &VarFormat{
	Begin:  "${",
	End:    "}",
	Escape: '\\',
}

// VarFormat defines what a variable expression should look like.
type VarFormat struct {
	Begin  string // The string token that preceeds a variable name
	End    string // The string token that follows a variable name
	Escape byte   // The escale character used to skip one of the above tokens
}

func (vf *VarFormat) lenOf(c cat) int {
	switch c {
	case cNorm:
		return 1
	case cBegin:
		return len(vf.Begin)
	case cEnd:
		return len(vf.End)
	case cEscNorm:
		return 2
	case cEscBegin:
		return len(vf.Begin) + 1
	case cEscEnd:
		return len(vf.End) + 1
	default:
		return 0
	}
}

func (vf *VarFormat) replString(s string) *replString {
	return &replString{s, len(s), vf}
}

type replString struct {
	str    string
	length int
	*VarFormat
}

func (rs *replString) next() *token {
	var i, x int
	p := -1

	for i < rs.length {
		c := rs.posCat(i)

		debugf("        i=%-2d  c=%c  %v  p=%-2d  x=%-2d\n", i, rs.str[i], c, p, x)

		switch c {
		case cBegin:
			p = i

			if x > 0 {
				x--
			}
		case cEscBegin:
			x++
		case cEnd:
			if x == 0 && p >= 0 {
				return rs.token(p, i)
			}

			if x > 0 {
				x--
			}
		}
		i += rs.lenOf(c)
	}
	return nil
}

func (rs *replString) token(beg, end int) *token {
	vb := beg + rs.lenOf(cBegin)
	n := strings.Replace(rs.str[vb:end], "\\", "", -1)
	return &token{n, rs.str, beg, end}
}

func (rs *replString) posCat(p int) cat {
	var esc bool

	if rs.str[p] == rs.Escape {
		esc = true
		p++
	}

	var (
		be = p + len(rs.Begin)
		ee = p + len(rs.End)
	)

	switch {
	case rs.length >= be && rs.str[p:be] == rs.Begin:
		return cBegin.escaped(esc)
	case rs.length >= ee && rs.str[p:ee] == rs.End:
		return cEnd.escaped(esc)
	default:
		return cNorm.escaped(esc)
	}
}
