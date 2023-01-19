//Namozov Azizbek 1-topshiriq

// 1-masala
// package main

// import "fmt"

// func main(){
// 	var a, b, c int
// 	fmt.Scan(&a, &b, &c)
// 	fmt.Println(max(max(a, b), c))
// }
// func max(a, b int) int{
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// 2-masala
// package main

// import "fmt"

// func main(){
// 	var n int
// 	fmt.Scan(&n)
// 	if (n % 4 == 0 && n % 100 != 0) || n % 400 == 0{
// 		fmt.Println("Yes")
// 	}else{
// 		fmt.Println("No")
// 	}
// }

// 3-masala
// package main

// import "fmt"

// func main(){
// 	var n int
// 	fmt.Scan(&n)
// 	fmt.Println(n/100 + n%10 + n%100/10)
// }

// 4-masala
// package main

// import "fmt"

// func main(){
// 	var n , mx int
// 	fmt.Scan(&n)
// 	mx = n
// 	for i := 0; i < 4; i ++{
// 		fmt.Scan(&n)
// 		if mx < n{
// 			mx = n
// 		}
// 	}
// 	fmt.Println(mx)
// }

// 5-masala
// package main

// import "fmt"

// func main(){
// 	var a, b int
// 	var q string
// 	fmt.Scan(&a, &q, &b)
// 	if q == "*"{
// 		print(a*b)
// 	}
// 	if q == "/"{
// 		a := float32(a)
// 		b := float32(b)
// 		fmt.Print(a/b)
// 	}
// 	if q == "+"{
// 		print(a+b)
// 	}
// 	if q == "-"{
// 		print(a-b)
// 	}
// }

// 6-masala
// package main

// import "fmt"

// func main() {
// 	var n int
// 	fmt.Scan(&n)
// 	switch n {
// 	case 1:
// 		print("Dushanba")
// 	case 2:
// 		print("Seshanba")
// 	case 3:
// 		print("Chorshanba")
// 	case 4:
// 		print("Payshanba")
// 	case 5:
// 		print("Juma")
// 	case 6:
// 		print("Shanba")
// 	case 7:
// 		print("Yakshanba")
// 	default :
// 		print("Bunday kun yo'q")
// 	}
// }

//7-masala
// package main

// import "fmt"

// func main(){
// 	var n int
// 	fmt.Scan(&n)
// 	if(n <= 100 && n >= 80){
// 		print(5)
// 	}else if(n >= 60 && n < 80){
// 		print(4)
// 	}else if(n >= 40 && n < 60){
// 		print(3)
// 	}else if(n >= 20 && n < 40){
// 		print(2)
// 	}else {
// 		print("Siz bebahosiz")
// 	}
// }
