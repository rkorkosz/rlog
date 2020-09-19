package slug

import "strings"

var sub = map[rune]rune{
	'ą': 'a',
	'ć': 'c',
	'ę': 'e',
	'ł': 'l',
	'ń': 'n',
	'ó': 'o',
	'ś': 's',
	'ź': 'z',
	'ż': 'z',
	'&': 'i',
	' ': '-',
}

// Slug returns a slug from a string
func Slug(s string) string {
	s = strings.ToLower(s)
	f := func(r rune) rune {
		v, ok := sub[r]
		if ok {
			return v
		}
		return r
	}
	return strings.Map(f, s)
}
