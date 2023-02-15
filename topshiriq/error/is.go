// package main

// import (
// 	"errors"
// 	"fmt"
// )

// var ErrFoo = errors.New("foo")

//	func isErrFoo(e error) bool {
//		return errors.Is(e, ErrFoo)
//	}
//
//	func main() {
//		var err = errors.New("f")
//		fmt.Println(isErrFoo(err), err, ErrFoo)
//	}
package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main_is() {
	if _, err := os.Open("non-existing1"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println("open file")
	}
}
