package main

import (
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
func main() {
	//content, err := ioutil.RadFile("main.go")
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
func main() {
	//resp, _ := http.Get("http://example.com")
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

	//1. URLを生成する(POSTの場合はhttpsのケースが多い)
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	//2.1 リクエストの生成(GET)
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-match", `W/"wyzzy"`)
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println(q)
	fmt.Println(q.Encode())

	req.URL.RawQuery = q.Encode()

	//2.2 リクエストの生成(POST)
	//今回はPostのURLがないため404が返ってくる。
	//req, _ := https.NewRequest("POST", endpoint, bytes.NewBuffer([]byte("password")))
	//req.Header.Add("If-None-match", `W/"wyzzy"`)
	//q := req.URL.Query()
	//q.Add("c", "3&%")
	//fmt.Println(q)
	//fmt.Println(q.Encode())
	//req.URL.RawQuery = q.Encode()


	//3. Clientの作成
	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

type T struct {

}

type Person struct{
	Name string `json:"name,omitempty"`
	Age int `json:"age,omitempty"`
	Nicknames []string `json:"nicknames,omitempty"`
	T *T `json:"T,omitempty"`
}

//func(p *Person) UnmarshalJSON(b []byte) error {
//	type Person2 struct{
//		Name string
//	}
//	var p2 Person2
//	err := json.Unmarshal(b, &p2)
//	if err != nil{
//		fmt.Println(err)
//	}
//	p.Name = p2.Name + "!"
//	return err
//}

//func(p Person) MarshalJSON() ([]byte, error) {
//	v, err := json.Marshal(&struct {
//		Name string
//	}{
//		Name: "Mr." + p.Name,
//	})
//	return v, err
//}

func main(){
	b := []byte(`{"name":"Mike","age":20,"nicknames":[]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil{
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}

var DB = map[string]string{
	"User1Key":"User1Secret",
	"User2Key":"User2Secret",
}

func Server(apiKey, sign string, data []byte){
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC)
}

func main(){
	const apikey = "User2Key"
	const apiSecret = "User2Secret"

	data := []byte("data")
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign)
	Server(apikey, sign, data)
}

var s *semaphore.Weighted = semaphore.NewWeighted(1)


func longProcess(ctx context.Context){
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}
	//if err := s.Acquire(ctx, 1); err != nil{
	//	fmt.Println(err)
	//	return
	//}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(2 * time.Second)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}

type ConfigList struct{
	Port int
	DbName string
	SQLDriver string
}

var Config ConfigList

func init(){
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(),
		DbName: cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main(){
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)

}

func main() {
	spy, _ := quote.NewQuoteFromYahoo("spy", "2019-01-01", "2020-01-01", quote.Daily, true)
	fmt.Print(spy.CSV())
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
	mva := talib.Ema(spy.Close, 14)
	fmt.Println(mva)
}
type JsonRPC2 struct {
	Version string `json:"jsonrpc"`
	Method string `json:"method"`
	Params interface{} `json:"params"`
	Result interface{} `json:"result,omitempty"`
	Id *int `json:"id,omitempty"`
}
type SubscribeParams struct {
	Channel string `json:"channel"`
}

func main() {
	u := url.URL{Scheme: "wss", Host: "ws.lightstream.bitflyer.com", Path:"/json-rpc"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(),nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	if err := c.WriteJSON(&JsonRPC2{Version: "2.0", Method: "subscribe", Params: &SubscribeParams{"lightning_ticker_BTC_JPY"}});
	err != nil {
		log.Fatal("subscribe:", err)
		return
	}

	for {
		message := new(JsonRPC2)
		if err := c.ReadJSON(message); err != nil {
			log.Println("read:", err)
			return
		}
		if message.Method == "channelMessage"{
			log.Println(message.Params)
		}
	}
}

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()
	cmd := `CREATE TABLE IF NOT EXISTS person(
			name STRING,
			age INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	//cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	//_, err = DbConnection.Exec(cmd, "Nansy", 20)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//cmd = "UPDATE person SET age = ? WHERE name = ?"
	//_, err = DbConnection.Exec(cmd, 25, "Mike")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	//cmd = "SELECT * FROM person"
	//rows, _ := DbConnection.Query(cmd)
	//defer rows.Close()
	//var pp []Person
	//for rows.Next(){
	//	var p Person
	//	err := rows.Scan(&p.Name, &p.Age)
	//	if err != nil{
	//		log.Println(err)
	//	}
	//	pp = append(pp, p)
	//}
	//err = rows.Err()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for _, p := range pp {
	//	fmt.Println(p.Name, p.Age)
	//}

	//cmd = "SELECT * FROM person where age = ?"
	//row := DbConnection.QueryRow(cmd, 1000)
	//var p Person
	//err = row.Scan(&p.Name, &p.Age)
	//if err != nil {
	//	if err == sql.ErrNoRows{
	//		log.Println("No row")
	//	}else {
	//		log.Println(err)
	//	}
	//}
	//fmt.Println(p.Name, p.Age)

	//cmd = "DELETE FROM person WHERE name = ?"
	//_, err = DbConnection.Exec(cmd, "Nancy")
	//if err != nil{
	//	log.Fatalln(err)
	//}
	tableName := "person; INSERT INTO person (name, age) VALUES('Mr.X',100)"
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := DbConnection.Query(cmd)
	defer rows.Close()
	var pp []Person
	for rows.Next(){
		var p Person
		err := rows.Scan(&p.Name, &p.Age)
		if err != nil{
			log.Println(err)
		}
		pp = append(pp, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
*/

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error{
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error){
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil{
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main(){
	p1 := &Page{Title: "test", Body: []byte("This is a sample Page.")}
	p1.save()

	p2, _ := loadPage(p1.Title)
	fmt.Println(string(p2.Body))
}