package main

import "fmt"

/*
<<golang_lesson 変数宣言>>
----------------------------------
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
----------------------------------


<<golang_lesson const>>
----------------------------------
const Pi = 3.14
const (
	Username = "test_user"
	Password = "test_pass"
)
varの場合はoverfrowしてしまう
var big int = 9223372036854775807 + 1
constは変数を認識しているが関数内で処理出来る。

const big = 9223372036854775807 + 1
----------------------------------
*/


func main() {
	//golngのお作法で”=”のところを揃える
	/*
		var(
			u8 uint8 	  = 255
			i8 int8 	  = 127
			f32 float32   = 0.2
			c64 complex64 = -5 + 12i
		)
		fmt.Println(u8, i8, f32, c64)

		fmt.Printf("%T %v", u8, u8)
	*/
	/*
	x := 1 + 1
	fmt.Println(x)
	fmt.Println(1+1, 2+2)
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1 =", 10-1)
	fmt.Println("10 / 2 =", 10/2)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10.0 / 3 =", 10.0/3)
	fmt.Println("10 / 3.0 =", 10/3.0)
	fmt.Println("10 % 2 =", 10%2)
	fmt.Println("10 % 3 =", 10%3)
	 */
	/*
	x := 0
	fmt.Println(x)
	// x = x + 1
	x++
	fmt.Println(x)
	// x = x - 1
	x--
	fmt.Println(x)
	 */
	fmt.Println(1 << 0) //0001 0001
	fmt.Println(1 << 1) //0001 0010
	fmt.Println(1 << 2) //0001 0100
	fmt.Println(1 << 3) //0001 1000
}
