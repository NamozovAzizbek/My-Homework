package main

import "fmt"

func Qovs() bool {
	// (([])), ({}())
	var s string
	fmt.Scan(&s)
	if len(s)%2 != 0{
		return false
	}
	st := []rune{}
	open :=  map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, v := range s{

		if close, ok := open[v]; ok{
			st = append(st, close)
			continue
		}
		//({}())
		l := len(st)-1
		if l < 0 || st[l] != v{
			return false
		}
		st = st[:l]
	}
	return len(st) == 0
}