package main

import (
	"fmt"

	
)

func ToUpper() bool {
	var s string
	fmt.Scan(&s)
	if len(s) % 2 != 0{
		return false
	}
	st := []rune{}
	open := map[rune]rune{
		'(':')',
		'[':']',
		'{':'}',
	}
	for _, r := range s {

		if closer, ok := open[r]; ok {
			fmt.Println(string(closer), ok)
		  st = append(st, closer)
		  continue
		}
	
		l := len(st) - 1
		if l < 0 || r != st[l] {
		  return false
		}
	
		st = st[:l]
	  }
	
	  return len(st) == 0
	}

// ({)}
