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

	<<golang_lesson 数値型>>
----------------------------------

	//golngのお作法で”=”のところを揃える
		var(
			u8 uint8 	  = 255
			i8 int8 	  = 127
			f32 float32   = 0.2
			c64 complex64 = -5 + 12i
		)
		fmt.Println(u8, i8, f32, c64)

		fmt.Printf("%T %v", u8, u8)
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
	x := 0
	fmt.Println(x)
	// x = x + 1
	x++
	fmt.Println(x)
	// x = x - 1
	x--
	fmt.Println(x)
	fmt.Println(1 << 0) //0001 0001
	fmt.Println(1 << 1) //0001 0010
	fmt.Println(1 << 2) //0001 0100
	fmt.Println(1 << 3) //0001 1000
----------------------------------

	<<golang_lesson 文字型>>
----------------------------------

	fmt.Println("Hello world")
	fmt.Println("Hello "+"world")
	//このままだとASCIIコード=>72
	fmt.Println("Hello world"[0])
	//Cast=>"H"
	fmt.Println(string("Hello world"[0]))
	var s string = "Hello world"
	//replace
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)
	//contains
	fmt.Println(strings.Contains(s, "world"))
	fmt.Println(`Test
Test`)
	fmt.Println("\"")
	fmt.Println(`"`)

----------------------------------

	<<golang_lesson 論理値型>>
----------------------------------
	//normal
	//var t, f bool = true, false
	//short
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, 1, t)
	fmt.Printf("%T %v %t\n", f, 0, f)
	// &&はどちらも同じならtrue
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	// ||はどちらかがtrueならtrue
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	//　!は否定
	fmt.Println(!true)
	fmt.Println(!false)
----------------------------------

	<<golang_lesson 型変換>>
----------------------------------
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d\n", yy, yy, yy)

	var s string = "14"
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)
	h := "Hello World"
	fmt.Println(string(h[0]))

----------------------------------

	<<golang_lesson 配列>>
----------------------------------
	/*
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b = [2]int{300, 400}
	//配列はサイズを変更できない
	fmt.Println(b)

	var b = []int{500, 600}
	b = append(b, 700)
	fmt.Println(b)
----------------------------------

	<<golang_lesson 配列>>
----------------------------------
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[:2])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	n[2] = 100
	fmt.Println(n)
	fmt.Printf("%T %v\n", n, n)

	var board = [][]int{
		{0,1,2},
		{3,4,5},
		{6,7,8},
	}
	fmt.Println(board)
	//最後に追加されていく
	n = append(n, 100, 200, 300, 400, 500)
	fmt.Println(n)
----------------------------------

	<<golang_lesson スライスのmap, cap>>
----------------------------------
	//makeで初期化([]型, 長さ, キャパシティ)
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)
	b := make([]int, 0)
	var c []int
	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

	//c = make([]int, 5)
	c = make([]int, 0, 5)
	for i := 0; i < 5; i++{
		c = append(c, i)
		fmt.Println(c)
	}

	//自分で書いてみる
	d := make([]int,0, 10)
	for x := 0; x < 10; x++{
		d = append(d, x)
		fmt.Println(d)
	}
----------------------------------

	<<golang_lesson map>>
----------------------------------
 m := map[string]int{"apple": 100, "banana": 200}
 fmt.Println(m)
 fmt.Println(m["apple"])
 m["banana"] = 300
 fmt.Println(m)
 m["new"] = 500
 fmt.Println(m)

 fmt.Println(m["nothing"])

 v, ok := m["banana"]
 fmt.Println(v, ok)

 v2, ok2 := m["nothing"]
 fmt.Println(v2, ok2)

 m2 := make(map[string]int)
 m2["pc"] = 5000
 fmt.Println(m2)

 //これはエラーになる
 var m3 map[string]int
 m3["pc"] = 5000
 fmt.Println(m3)

 var s []int
 if s == nil{
  fmt.Println("Nil")
 }
----------------------------------

	<<golang_lesson byte>>
----------------------------------
 b := []byte{72, 73}
 fmt.Println(b)
 fmt.Println(string(b))

 c := []byte("HI")
 fmt.Println(c)
 fmt.Println(string(c))

func otameshi(x string, y string){
	fmt.Println("とりあえずお決まりの")
	fmt.Println(x + y)
}

----------------------------------

	<<golang_lesson 関数>>
----------------------------------
func add(x, y int) (int, int) {
	return x + y, x - y
}

func cal(price, item int) (result int){
	result = price * item
	return
}

func convert(price int) float64{
	return float64(price)
}

func main() {
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)
	r3 := cal(300, 5)
	fmt.Println(r3)
	r4 := convert(500)
	fmt.Println(r4)

	f := func(x int){
		fmt.Println("inner func", x)
	}
	f(1)

	func(x int){
		fmt.Println("inner func",x)
	}(1)
}

----------------------------------

	<<golang_lesson クロージャ>>
----------------------------------

func incrementGenerator() (func() int) {
	x := 0
	return func() int {
		x++
		return x
	}
}

func cicleArea(pi float64) func(radius float64) float64{
	return func(radius float64) float64{
		return pi * radius * radius
	}
}

func main() {
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	c1 := cicleArea(3.14)
	fmt.Println(c1(2))

	c2 := cicleArea(3)
	fmt.Println(c2(2))

}
----------------------------------

	<<golang_lesson 可変長変数>>
----------------------------------
// 可変長変数
func foo(params ...int){
	fmt.Println(len(params), params)
	for _, param := range params{
		fmt.Println(param)
	}

}


func main() {
	foo()
	foo(10,20)
	foo(10, 20, 30)
	//スライスの可変長変数の渡し方
	s := []int{1, 2, 3}
	fmt.Println(s)
	foo(s...)
}

----------------------------------

	<<golang_lesson 演習>>
----------------------------------

func foo() {
	mm := map[string]int{
		"Mike": 20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v\n", mm, mm)
}


func main() {
	f := 1.11
	i := int(f)
	//fmt.Println(f)
	fmt.Println("------------------------------")
	fmt.Printf("演習Q1：%T %v\n", i, i )
	fmt.Println("演習Q2：", 5,6)
	foo()
	//fmt.Println("演習Q3：", int(f))
	fmt.Println("------------------------------")

}

----------------------------------

	<<golang_lesson for文>>
----------------------------------
func main() {
	//roop毎に+1
	//変数（i）の宣言をforに含む方法
	for i := 0; i < 10; i++{
		if i == 3{
			fmt.Println("continue")
			continue
		}

		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
	//roop毎に+sumの値
	//変数（sum）を先に宣言しておく方法
	sum := 1
	for ; sum < 10; {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
	//無限ループ
	for {
		fmt.Println("hello")
	}
}

----------------------------------

	<<golang_lesson range文>>
----------------------------------
func main() {
	l := []string{"python", "go", "java"}
	l = append(l, "ruby")
	fmt.Printf("%T %v\n", l, l)
	//forで書いてみる
	for i := 0; i < len(l); i++{
		fmt.Println(i, l[i])
	}
	//rangeで書いてみる(上と一緒)
	for i, v := range l{
		fmt.Println(i, v)
	}
	//rangeのインデックス番号不要バージョン（"_"でそうなるらしい）
	for _, v := range l{
		fmt.Println(v)
	}
	//mapでの場合
	m := map[string]int{"test1": 100, "test2": 200}
	//mapに要素追加
	m["test3"] = 300
	for k, v := range m{
		fmt.Println(k, v)
	}
	//keyのみは"_"が不要
	for k := range m{
		fmt.Println(k)
	}
	//valueのみはスライスと同様
	for _, v := range m{
		fmt.Println(v)
	}
}

----------------------------------

	<<golang_lesson switch文>>
----------------------------------

func getOsname()string{
	return "mac"
}

func main() {
	os := getOsname()
	switch os {
	case "mac":
		fmt.Println("Mac!")
	case "windows":
		fmt.Println("Windows!")
	default:
		fmt.Printf("default")
	}
	fmt.Println(os)

	/*　switch文以降で変数を使わないなら変数のところもワンライナーで書ける
	switch os := getOsname(); os {
	case "mac":
		fmt.Println("Mac!")
	case "windows":
		fmt.Println("Windows!")
	default:
		fmt.Printf("default")
	}
	fmt.Println(os)


----------------------------------

	<<golang_lesson defer>>
----------------------------------

func foo(){
	defer fmt.Println("foo world")

	fmt.Println("foo hello")
}

func main() {
	/*
	foo()
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")


file, _ := os.Open("./lesson.go")
defer file.Close()
data := make([]byte, 100)
file.Read(data)
fmt.Println(string(data))
}


----------------------------------

	<<golang_lesson log>>
----------------------------------
func loggingSettings (logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func main() {
	loggingSettings("test.log")
	_, err := os.Open("ffsfsf")
	if err != nil{
		log.Fatalln("Exit..", err)
	}
	log.Println("Logging")
	log.Printf("%T %v", "test", "test")

	log.Fatalf("%T %v", "test", "test" )
	log.Fatalln("error!!")

	fmt.Println("OK!")
}
----------------------------------

	<<golang_lesson エラーハンドリング>>
----------------------------------
func main() {
	file, err := os.Open("./lesson.go")
	//errorじゃなければnilが返ってくる。
	if err != nil{
		log.Fatalln("Error!")
	}
	defer file.Close()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil{
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	if err = os.Chdir("test"); err != nil{
		log.Fatalln("Error")
	}
}

----------------------------------
<<golang_lesson panicとrecover>>
----------------------------------
func thirdPartyConnectDB() {
	panic("Unable to database!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}
func main() {
	save()
	fmt.Println("OK?")
}

func thirdPartyConnectDB() {
	panic("Unable to database!")
}

func save() {
	defer func() {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}
func main() {
	save()
	fmt.Println("OK?")
}
----------------------------------
<<golang_lesson 演習２>>
----------------------------------

func ques1(l []int){
	sort.Sort(sort.IntSlice(l))
	fmt.Println("Ans >>", l[0])
}

func ques2(m map[string]int){
	sum := 0
	for _, v := range m{
		//fmt.Println(v)
		sum += v
		//fmt.Println(sum)
	}
	fmt.Println("Ans >>",sum)
}

func main() {
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	fmt.Println("Q1: 以下から最小値を出力せよ\n", l)
	ques1(l)
	m := map[string]int{
		"apple": 200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi": 90,
	}
	fmt.Println("Q2: 果物の価格を合計せよ\n", m)
	ques2(m)
}

----------------------------------
<<golang_lesson ポインタ>>
----------------------------------

func one(x *int){
	*x = 1
}

func main() {
	var n int = 100
	//&nはポインタ
	one(&n)
	fmt.Println(n)
	/*
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)

	fmt.Println(*p)

}
----------------------------------
<<golang_lesson >>newとmakeの違い
----------------------------------
func main() {
	//空のスライス
	s := make([]int, 0)
	fmt.Printf("%T\n", s)
	//空のmap
	m := make(map[string]int)
	fmt.Printf("%T\n", m)
	//チャネル
	ch := make(chan int)
	fmt.Printf("%T\n", ch)
	//ポインタ
	var p *int = new(int)
	fmt.Printf("%T\n", p)
	//ストラクト
	var st = new(struct{})
	fmt.Printf("%T\n", st)
	/*
	var p *int = new(int)
	fmt.Println(p)
	*p++
	fmt.Println(*p)

	var p2 *int
	fmt.Println(p2)
	*p2++
	fmt.Println(p2)
	}
*/

type Vertex struct{
	X int
	Y int
	S string
}

func changeVertex(v Vertex){
	v.X = 1000
}

func changeVertex2(v *Vertex){
	v.X = 1000
}

func main() {
	v := Vertex{1, 2, "test"}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "test"}
	changeVertex2(v2)
	fmt.Println(*v2)
	/*
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)
	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1,2,"test"}
	fmt.Println(v3)

	v4 := Vertex{}
	fmt.Println(v4)
	fmt.Printf("%T %v\n", v4, v4)

	var v5 Vertex
	fmt.Println(v5)
	fmt.Printf("%T %v\n", v5, v5)

	v6 := new(Vertex)
	fmt.Println(v6)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Println(v7)
	fmt.Printf("%T %v\n", v7, v7)
	*/

}