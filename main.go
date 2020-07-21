package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

/*
import (
	"fmt"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2016-01-01", "2016-04-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
}
func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func main() {
	//正規表現 "a([a-z]+)e" => "aで始まってeで終わる、かつその間の文字列にa〜ｚまでの1文字が含まれているか。"
	//そのため、appleをapp9leとかでfalseになる
	//ただし、"a([a-z0-9]+)e"などにして数字も含めるとtrueになる
	match, _ := regexp.MatchString("a([a-z]+)e", "apple")
	fmt.Println(match)
	//何度も確認するなら下のほうが最適
	r := regexp.MustCompile("a([a-z]+)e")
	ms := r.MatchString("apple")
	fmt.Println(ms)
	// s := "/view/test"
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
	fs := r2.FindString("/view/test")
	fmt.Println(fs)
	fss := r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
	fss = r2.FindStringSubmatch("/save/test")
	fmt.Println(fss, fss[0], fss[1], fss[2])
}

func main() {
	i := []int{5, 3, 2, 8, 7}
	s := []string{"d","a","f"}
	p := []struct {
		Name string
		Age int
	}{
		{"Nancy", 20},
		{"Vera", 40},
		{"Mike", 30},
		{"Bob", 50},
	}
	fmt.Println(i, s, p)
	sort.Ints(i)
	fmt.Println(i)
	sort.Strings(s)
	fmt.Println(s)
	sort.Slice(p, func(i, j int ) bool {return p[i].Name < p[j].Name })
	sort.Slice(p, func(i, j int ) bool {return p[i].Age < p[j].Age })
	fmt.Println(i, s, p)
}

const (
	c1 = iota
	c2
	c3
)

const (
	_ = iota
	KB int = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println(c1,c2,c3)
	fmt.Println(KB, MB, GB)
}

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go longProcess(ctx, ch)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("##############")
}
*/

func main() {
	//content, err := ioutil.ReadFile("main.go")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(string(content))
	//
	//if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
	//	log.Fatal(err)
	//}
	r := bytes.NewBuffer([]byte("abc"))
	content, _ := ioutil.ReadAll(r)
	fmt.Println(string(content))
}
