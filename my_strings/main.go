package my_strings

import "strings"

type Str struct {
	s string
}

func NewStr(s string) Str {
	return Str{
		s: s,
	}
}

func (s *Str) ToUpper() *Str {
	s.s = strings.ToUpper(s.s)
	return s
}
