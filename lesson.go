package main

import "fmt"

/*
var (
	i int = 1
	f64 float64 = 1.2
	s string = "test"
	t, f bool = true,false
)

func foo() {
短縮記法は関数外では使えなくなる
	xi := 1
	xi = 2
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)

}
*/

const Pi = 3.14
const (
	Username = "test_user"
	Password = "test_pass"
)
/* varの場合はoverfrowしてしまう
var big int = 9223372036854775807 + 1
constは変数を認識しているが関数内で処理出来る。
*/
const big = 9223372036854775807 + 1
func main() {
	fmt.Println(Pi, Username, Password)
	fmt.Println(big-1)
}