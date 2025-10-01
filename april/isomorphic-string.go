package april

import "fmt"

func isIsomorphic(s string, t string) bool {
	as := [150]byte{}
	at := [150]byte{}

	for i := 0; i < len(s); i++ {
		if as[s[i]] == 0 {
			as[s[i]] = t[i]
		}
		if at[t[i]] == 0 {
			at[t[i]] = s[i]
		}
		if as[s[i]] != t[i] || at[t[i]] != s[i] {
			return false
		}
	}

	return true
}

func Day2() {
	qs := [][]string{
		{"egg", "add"},
		{"foo", "bar"},
		{"paper", "title"},
		{"bbbaaaba", "aaabbbba"},
		{"badc", "baba"},
	}

	for _, v := range qs {
		fmt.Println(isIsomorphic(v[0], v[1]))
	}
}
