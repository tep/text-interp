package interp

type token struct {
	name  string
	str   string
	start int
	end   int
}

func (t *token) replace(val string) string {
	return t.str[:t.start] + val + t.str[t.end+1:]
}
