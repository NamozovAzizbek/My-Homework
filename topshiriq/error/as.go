package main

import (
	"errors"
	"fmt"
)

type MyError_As struct{s string}

func (e MyError_As)Error_As()string{
	return 	e.s
}

func doFoo() error{
	return fmt.Errorf("Wrapped: %w",MyError_As{"mine"})
}
func As(){
	var MyErr MyError_As
	caught := errors.As(doFoo(), &MyErr)
	fmt.Printf("Caught: %v, MyErr: %v",caught, MyErr)//3:47
}

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"io/fs"
// 	"os"
// )

// func main() {
// 	if _, err := os.Open("non-existing1"); err != nil {
// 		var pathError *fs.PathError
// 		if errors.As(err, &pathError) {
// 			fmt.Println("Failed at path:", pathError.Path)
// 		} else {
// 			fmt.Println(err)
// 		}
// 	}

// }