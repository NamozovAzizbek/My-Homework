// 2- topshiriq Namozov Azizbek
// 1-masala. son kiritiladi sonni teskari tartibda chiqarish kerak.
// package main

// import "fmt"

// func rec(n int){
// 	if n == 0 {
// 		return
// 	}
// 	fmt.Print(n%10)
// 	n /= 10
// 	rec(n)
// }
// func main(){
// 	var n int
// 	fmt.Scan(&n)
// 	rec(n)
// }

//2-masal son kiritiladi son 3 va 6 ga bo'lisa good aks holda bad
// package main

// import "fmt"

// func main(){
// 	var n int
// 	fmt.Scan(&n)
// 	if n % 3 == 0 && n % 6 != 0{
// 		fmt.Print("Good")
// 	}else{
// 		fmt.Print("Bad")
// 	}
// }
// 3-masal foydalanuvchi son kiritadi son 3, 5, 7 sonlarga bo'linsa good aks holda bad chiqarsin
// package main

// import "fmt"

// func main(){
// 	var n int
// 	fmt.Scan(&n)
// 	if n % 3 == 0 && n % 5 == 0 && n % 7 == 0{
// 		fmt.Print("Good")
// 	}else{
// 		fmt.Print("Bad")
// 	}
// }
// 4-masala son kiritiladi shu son 10 ga bo'linib 20 ga bo'linmasa good aks holda bad
// package main

// import "fmt"

// func main(){
// 	var n int
// 	fmt.Scan(&n)
// 	if n % 10 == 0 && n % 20 != 0{
// 		fmt.Println("Good")
// 	}else {
// 		fmt.Println("Bad")
// 	}
// }
// 5-masala. son kiritiladi shu sonda qatnashgan eng kichik raqamni ekranga chiqaring
// package main

// import "fmt"

// func main(){
// 	var n int
// 	fmt.Scan(&n)
// 	mn := n % 10
// 	n /= 10
// 	for n != 0 {
// 		if mn > n % 10{
// 			mn  = n % 10
// 		}
// 		n /= 10
// 	}
// 	fmt.Println(mn)
// }
// 6-masal. hafta kuni kritiladi ish kuni yoki yo'qligini aniqlansin.
// package main

// import "fmt"

// func main(){
// 	var n int
// 	fmt.Scan(&n)
// 	switch n{
// 	case 1: fmt.Print("Ish kuni")
// 	case 2: fmt.Print("Ish kuni")
// 	case 3: fmt.Print("Ish kuni")
// 	case 4: fmt.Print("Ish kuni")
// 	case 5: fmt.Print("Ish kuni")
// 	case 6: fmt.Print("Dam olish kuni")
// 	case 7: fmt.Print("Dam olish kuni")
// 	default : fmt.Println("Bunday kun mavjud emas")
// 	}
// }
// 7-masal. oy kiritiladi shu oyda necha kun borligin chiqarish kerak.
// package main

// import "fmt"

// func main(){
// 	var n int
// 	fmt.Scan(&n)
// 	switch n{
// 	case 1: fmt.Print(31)
// 	case 2: fmt.Print(28)
// 	case 3: fmt.Print(31)
// 	case 4: fmt.Print(30)
// 	case 5: fmt.Print(31)
// 	case 6: fmt.Print(30)
// 	case 7: fmt.Print(31)
// 	case 8: fmt.Print(31)
// 	case 9: fmt.Print(30)
// 	case 10: fmt.Print(31)
// 	case 11: fmt.Print(30)
// 	case 12: fmt.Print(31)
// 	default : fmt.Println("Bunday oy mavjud emas")
// 	}
// }
// 8-masal. A pangram is a sentence where every letter of the English alphabet appears at least once.

//Given a string sentence containing only lowercase English letters, return true if sentence is a pangram, or false otherwise.

// func checkIfPangram(sentence string) bool {
//     var n [123]bool
//     for _, v := range sentence{
// 		if v <= 96 || v >= 123{
// 			return false
// 		}else {
//             n[v] = true
//         }
// 	}
//     for i := 97; i <= 122; i ++{
// 		if n[i] == false{
// 			return false
// 		}
// 	}
// 	return true
// }
// 9-masal.Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

//Return the running sum of nums.
// func runningSum(nums []int) []int {

//	    for i := 1; i < len(nums); i ++{
//	      nums[i] = nums[i]+nums[i-1]
//	    }
//	    return nums
//	}
// 10-masala. You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.

//A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.


// func maximumWealth(accounts [][]int) int {
//     mx := 0
//     for _, v := range accounts{
//        if  mx < sum(v){
//            mx = sum(v)
//        }
// 	}
//     return mx
// }
// func sum (n []int) int{
// 	s := 0
// 	for _, v := range n{
// 		s += v
// 	}
// 	return s
// }