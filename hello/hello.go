package main

import (
	"errors"
	"fmt"
	"golang.org/x/example/stringutil"
	"os"
)

func main() {
	fmt.Println(stringutil.Reverse("Hello world!"))

	var a int
	fmt.Println(a)
	b := 789
	fmt.Println(b)

	var (
		e int = 987
		f int = 654
	)
	fmt.Println(e)
	fmt.Println(f)

	f = 321
	fmt.Println(f)

	fmt.Println("hello, world")

	name, age := "a", 34
	fmt.Println(name)
	fmt.Println(age)

	const dummy = 100
	const (
		pom = 100
		qom = 200
	)
	fmt.Println(pom)
	fmt.Println(qom)
	//コメント
	/*
	   コメント
	*/

	var num int8 = 123
	str := "ABC"
	fmt.Println(num)
	fmt.Println(str)

	var h byte = 0xFF
	fmt.Println(h)
	var stri string = "aaaa"
	fmt.Println(stri)

	type UtcTime string
	type Jstime string

	var t1 UtcTime = "00:"
	var t2 Jstime = "09:"

	fmt.Println(t1, t2)

	type (
		UtcTime2 string
		Jstime2  string
	)

	var a1 uint16 = 1234
	var a2 uint32 = uint32(a1)
	fmt.Println(a2)

	type MyTepe uint64
	type MyTepe2 uint32

	var mytype MyTepe = 1234
	var mytype2 MyTepe2 = MyTepe2(mytype)
	fmt.Println("mytype2:", mytype2)

	arr1 := [3]string{}
	arr1[0] = "red"
	arr1[1] = "green"
	arr1[2] = "blue"
	fmt.Println(arr1[0], arr1[1], arr1[2])

	var ar [3]string
	ar[0] = "ar"
	ar[1] = "ar2"
	fmt.Println(ar)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	arr3 := [...]bool{true, false, true}
	fmt.Println(arr3)

	arr4 := []string{}
	arr4 = append(arr4, "sweet")
	arr4 = append(arr4, "bitter")
	arr4 = append(arr4, "spicy")
	fmt.Println(arr4)

	arr5 := []int{}
	for i := 0; i < 10; i++ {
		arr5 = append(arr5, i)
		fmt.Println("len:", len(arr5), "cap:", cap(arr5))
	}

	bufa := make([]byte, 0, 1024)
	fmt.Println(bufa)

	arr6 := map[string]int{"x": 100, "y": 200}
	fmt.Println(arr6)
	fmt.Println(arr6["x"])

	arr6["z"] = 300
	fmt.Println(arr6)

	delete(arr6, "z")
	fmt.Println(arr6)

	fmt.Println(len(arr6))

	_, ok := arr6["z"]
	fmt.Println(ok)
	if ok {
		fmt.Println("exist")
	} else {
		fmt.Println("not exist")
	}

	for key, value := range arr6 {
		fmt.Printf("%s=%d\n", key, value)
	}

	x, y := 1, 2
	if x > y {
		fmt.Println("x is greater then y")
	} else {
		fmt.Println("x is less then y")
	}

	x, y = 1, 1
	if x > y {
		fmt.Println("big")
	} else if x < y {
		fmt.Println("small")
	} else {
		fmt.Println("eq")
	}

	mode := "stop"
	switch mode {
	case "running":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("不明")
	}

	x, y = 1, 1
	switch {
	case x > y:
		fmt.Println("big")
	case x < y:
		fmt.Println("small")
	default:
		fmt.Println("eq")
	}

	dayOfWeek := "sat"
	switch dayOfWeek {
	case "sat":
		fallthrough
	case "sun":
		fmt.Println("horiday")
	default:
		fmt.Println("weekday")
	}

	x, y = 5, 10
	for x < y {
		x++
		fmt.Println("x:", x)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	n := 0

	for {
		n++
		if n > 10 {
			break
		} else if n%2 == 1 {
			continue
		} else {
			fmt.Println(n)
		}
	}

	colors := [...]string{"red", "green", "blue"}

	for i, color := range colors {
		fmt.Printf("%d:%s \n", i, color)
	}

	funcA()
	fmt.Println(add(1, 3))

	add, min := addMinus(8, 5)
	fmt.Println(addMinus(add, min))
	funcB(1, 2, 3, 4, 5)

	var p1 Person
	p1.SetPerson("yamasa", 26)
	set_name, set_age := p1.GetPerson()
	fmt.Printf("%s(%d)\n", set_name, set_age)

	// p2_1 := Person2{"yaamada", 26, 1}
	// p2_2 := Person2{Age: 26, Name: "yamada"}

	a7 := Person{name: "taro"}
	a8 := Book{title: "neko"}
	PrintOut(a7)
	PrintOut(a8)

	// p2 := map[string]interface{}{
	// 	"name": "yamada",
	// 	"age":  26,
	// }

	p3 := dict{
		"name": "a",

		"a": dict{
			"name": "b",
		},
	}
	name2 := p3["name"]
	tel := p3["a"].(dict)["name"]
	fmt.Println(name2)
	fmt.Println(tel)

	var a9 int
	var ptr1 *int
	ptr1 = &a9
	*ptr1 = 123
	fmt.Println(ptr1, *ptr1)

	aa1 := 123
	aa2 := 123
	fn(aa1, &aa2)
	fmt.Println(aa1, aa2)

	aa3 := Person{"tanaka", 26}
	ptr3 := &aa3
	fmt.Println(aa3.name)
	fmt.Println((*ptr3).name)
	fmt.Println(ptr3.name)
	fmt.Println(ptr3.name)

	bookList := []*Book{}
	for i := 0; i < 10; i++ {
		book := new(Book)
		book.title = fmt.Sprintf("Title#%d", i)
		bookList = append(bookList, book)
	}
	for _, book := range bookList {
		fmt.Println(book.title)
	}

	funcE()

}

func funcE() {
	fp, err := os.Create("sample.txt")
	if err != nil {
		return
	}
	defer fp.Close()

	n := 0
	for n < 10 {
		conut, err := fp.Write([]byte("helloo\n"))
		if err != nil {
			fmt.Println(err)
			goto M
		}
		n++
		fmt.Println(conut)
	}
M:
}

func fn(b1 int, b2 *int) {
	b1 = 456
	*b2 = 456
}

type any interface{}
type dict map[string]any

func funcA() (string, error) {
	var err error
	filename := ""
	data := ""

	filename, err = GetFileName()
	if err != nil {
		fmt.Println(err)
		goto Done
	}

	data, err = ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		goto Done

	}
	fmt.Println(data)
Done:
	return data, err
}

func GetFileName() (string, error) {
	return "sample.txt", nil
}
func ReadFile(filename string) (string, error) {
	return "Hello world!", errors.New("Cant red file")
}

func add(x int, y int) int {
	return x + y

}

func addMinus(x int, y int) (int, int) {
	return x + y, x - y
}

func funcB(a int, b ...int) {
	fmt.Printf("a=%d\n", a)
	for i, num := range b {
		fmt.Printf("b[%d]=%d\n", i, num)
	}
}

type Person struct {
	name string
	age  int
}

func (p *Person) SetPerson(name string, age int) {
	p.name = name
	p.age = age
}

func (p *Person) GetPerson() (string, int) {
	return p.name, p.age
}

type Person2 struct {
	Name   string
	Age    int
	status int
}

type Printable interface {
	ToString() string
}

func PrintOut(p Printable) {
	fmt.Println(p.ToString())

}

func (p Person) ToString() string {
	return p.name
}

type Book struct {
	title string
}

func (b Book) ToString() string {
	return b.title
}

func funcC(a interface{}) {
	fmt.Printf("%d\n", a.(int))
}

func PrintOut2(a interface{}) {
	q, ok := a.(Printable)
	if ok {
		fmt.Println(q)
		fmt.Println(q.ToString())
	} else {
		fmt.Println("Not printable")
	}
}

func funcD(a interface{}) {
	switch a.(type) {
	case bool:
		fmt.Printf("%t\n", a)

	case int:
		fmt.Printf("%t\n", a)

	case string:
		fmt.Printf("%t\n", a)

	}
}
