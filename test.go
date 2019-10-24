package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strings"
)

// func main() {
// http.HandleFunc("/", HelloServer)
// http.ListenAndServe(":8081", nil)
// }

// func HelloServer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello, world")
// }

func getMap() map[string]string {
	nameMap := make(map[string]string)
	nameMap["a"] = "apple"
	nameMap["b"] = "ball"
	nameMap["c"] = "cup"
	nameMap["d"] = "desk"
	nameMap["e"] = "egg"
	nameMap["f"] = "frag"

	return nameMap
}

func avergeAndTop(list []int) (averge, top int) {
	top = 0
	averge = 0
	for _, num := range list {
		if num > top {
			top = num
		}
		averge += num
	}

	averge /= 2

	return
}

func deferTest(num int) {
	if num == 3 {
		panic(1)
	} else {
		fmt.Println("It's working")
	}
}

func deferAction() {
	if err := recover(); err != nil {
		fmt.Println("error messge: ", err)
	}
}

func changeName(namePoint *string) {
	*namePoint = "jack"
}

type person struct {
	name string
	age  int
}

type graphic interface {
	area() float64
	perimeter() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.width * s.height
}

func (s square) perimeter() float64 {
	return (2 * s.width) + (2 * s.height)
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c circle) perimeter() float64 {
	return 2 * c.radius * math.Pi
}

func measure(g graphic) {
	fmt.Println("graphic = ", g)
	fmt.Println("area = ", g.area())
	fmt.Println("perimeter = ", g.perimeter())
}

func httpGetTest() {
	response, err := http.Get("https://blog.syhlion.tw/sitemap.xml")
	if err != nil {
		fmt.Println("err1: ", err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("err2: ", err)
		return
	}

	fmt.Println(string(body))
}

func httpPostTest() {
	response, err := http.Post("https://blog.syhlion.tw/sitemap.xml",
		"application/x-www-form-undencoded",
		strings.NewReader("name=test"))
	if err != nil {
		fmt.Println("err1: ", err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("err2: ", err)
		return
	}

	fmt.Println(string(body))
}

func httpAnyTest() {
	client := &http.Client{}
	request, err := http.NewRequest("POST",
		"https://blog.syhlion.tw/sitemap.xml",
		strings.NewReader("name=test"))
	if err != nil {
		fmt.Println("err1: ", err)
		return
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Cookie", "name=test")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("err1: ", err)
		return
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("err2: ", err)
		return
	}

	fmt.Println(string(body))
}

func main() {
	// numList := []int{11, 23, 4, 45, 88, 12}
	// go avergeAndTop(numList)

	// for i := 0; i < 10; i++ {
	// replay:
	// 	if i == 5 {
	// 		break
	// 	}

	// 	switch i {
	// 	case 1:
	// 		fmt.Println("one = ", i)
	// 	case 2:
	// 		fmt.Println("test = ", i)
	// 	case 3:
	// 		i++
	// 		goto replay
	// 	case 4:
	// 		fmt.Println("i = ", i, ", other = ", true)
	// 	}
	// }

	// var arr [5]int
	// var slice []int
	// slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	// for i, k := 0, 10; i < len(arr); i++ {
	// 	arr[i] = k
	// 	k += 10
	// }
	// for _, item := range arr {
	// 	fmt.Println("arr ", item)
	// }
	// for _, item := range slice {
	// 	fmt.Println("slice ", item)
	// }

	// namemap := getMap()
	// fmt.Println(namemap)

	// numList := []int{11, 23, 4, 45, 88, 12}
	// averge, top := avergeAndTop(numList)
	// fmt.Println("Top = ", top, ", Averge = ", averge)

	// messages := make(chan int)
	// go func() { messages <- 234 }()
	// msg := <-messages
	// fmt.Println(msg)

	// channl1 := make(chan string)
	// channl2 := make(chan string)

	// go func() {
	// 	time.Sleep(time.Second * 1)
	// 	channl1 <- "one"
	// }()

	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	channl2 <- "two"
	// }()

	// go func() {
	// 	time.Sleep(time.Second * 3)
	// 	channl2 <- "two again"
	// }()

	// for i := 0; i < 3; i++ {
	// 	select {
	// 	case msg1 := <-channl1:
	// 		fmt.Println("received ", msg1)
	// 	case msg2 := <-channl2:
	// 		fmt.Println("received", msg2)
	// 	}
	// }

	// defer deferAction()
	// for i := 0; i < 5; i++ {
	// 	deferTest(i)
	// }

	// name := "mark"
	// fmt.Println("name = ", name)
	// changeName(&name)
	// prinfmt.Printlntln("name = ", name)

	// one := person{"jack", 18}
	// fmt.Println(one.name)
	// one.name = "mark"
	// fmt.Println(one.name)
	// fmt.Println(one.age)

	// s := square{10, 20}
	// c := circle{10}
	// measure(s)
	// measure(c)

	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	return
	// }
	// defer file.Close()

	// start, err := file.Stat()
	// if err != nil {
	// 	return
	// }

	// buffer := make([]byte, start.Size())
	// _, err = file.Read(buffer)
	// if err != nil {
	// 	return
	// }

	// str := string(buffer)
	// fmt.Println(str)

	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	return
	// }
	// defer file.Close()

	// reader, err := ioutil.ReadAll(file)
	// fmt.Println(string(reader))

	// now := time.Now()
	// fmt.Println(now)
	// fmt.Println(now.Format("2006-01-02"))
	// fmt.Println(now.Year())
	// then := time.Date(2019, 11, 12, 12, 56, 0, 0, time.UTC)
	// fmt.Println(then)
	// fmt.Println(then.Format("2006-01-02"))

	// date, _ := time.Parse("2006-01-02 15:04:05", "2019-11-26 12:10:00")
	// fmt.Println(date)

	// str := "abba"
	// str = strings.ReplaceAll(str, "a", "n")
	// fmt.Println(str)

	// hash := crc32.NewIEEE()
	// hash.Write([]byte("test"))
	// v := hash.Sum32()
	// fmt.Println(v)

	// h := sha1.New()
	// h.Write([]byte("test22222"))
	// bs := h.Sum([]byte{})
	// fmt.Println(bs)

	// var listTest list.List
	// listTest.PushBack(123)
	// listTest.PushBack(456)
	// listTest.PushBack(789)

	// for each := listTest.Front(); each != nil; each = each.Next() {
	// 	fmt.Println(each.Value.(int))
	// }

	// match, _ := regexp.MatchString("p[a-z]+ch", "peach")
	// fmt.Println(match)

	// regexRule, _ := regexp.Compile("p[a-z]+ch")
	// fmt.Println("1. ", regexRule.MatchString("peach"))
	// fmt.Println("2. ", regexRule.FindString("peach punch"))
	// fmt.Println("3. ", regexRule.FindStringIndex("peach punch"))
	// fmt.Println("4. ", regexRule.FindStringSubmatch("peach punch"))
	// fmt.Println("5. ", regexRule.FindStringSubmatchIndex("peach punch"))
	// fmt.Println("6. ", regexRule.FindAllString("peach punch pinch", -1))
	// fmt.Println("7. ", regexRule.FindAllStringSubmatchIndex("peach punch pinch", -1))
	// fmt.Println("8. ", regexRule.FindAllString("peach punch pinch", 2))
	// fmt.Println("9. ", regexRule.Match([]byte("peach")))

	// regexRule = regexp.MustCompile("p([a-z]+)ch")
	// fmt.Println("10. ", regexRule)

	// httpGetTest()
	// httpPostTest()
	// httpAnyTest()

	// m := martini.Classic()
	// m.Get("/", func() string {
	// 	return "Hello world"
	// })
	// m.Run()

}

func getTicket() (ID, amount int, name string) {

	db, err := sql.Open("mysql", "root:root@/Ticket_DB")
	if err != nil {
		fmt.Println("DB err: ", err)
		panic(1)
	}
	defer db.Close()

	ticketRow, err := db.Query("SELET ID FROM Ticket_amount")
	if err != nil {
		fmt.Println("sql err: ", err)
	}
	defer ticketRow.Close()

	return
}
