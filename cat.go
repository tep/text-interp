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
