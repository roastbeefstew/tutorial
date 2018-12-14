package main

import (
	"fmt"
	"math"
	"runtime"
)

func add(x int, y int) int {
	return x + y
}

// alternate parameters
func add2(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// type can be omitted if initializer present
var c, python, java = true, false, "no!"

// initialize multiple variables at once
var i, j int = 1, 2

const pi = 3.14

func insideFunction() {
	fmt.Println("insideFunction()")

	// var can be omitted with :=
	foo, bar := true, 100

	fmt.Println(foo, bar)
}

func loop() {
	fmt.Println("loop()")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// while loop
	sum = 1
	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)
}

func conditional() float64 {
	fmt.Println("arrayTconditionalest()")

	x := 1
	if x < 10 {
		x++
	}
	fmt.Println(x)

	if v := math.Pow(10, 2); v < 10 {
		return v
	} else {
		// v is declared in the if statement but still accessible here
		return v * 10
	}
}

func switchStatement() {
	fmt.Println("switchStatement()")

	// breaks are automatic per case
	// case statement can even call functions

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}

	// switch w/out condition is always switching on true. Handy way of writing long if-then-else chains
	switch {
	case false:
		fmt.Println(false)
	case true:
		fmt.Println(true)
	}

}

type Vertex struct {
	X int
	Y int
}

type Graph struct {
	X, Y float64
}

// add member function to Graph struct
func (g Graph) Abs() float64 {
	return math.Sqrt(g.X*g.X + g.Y*g.Y)
}

func memberFuncTest() {
	fmt.Println("memberFuncTest()")

	g := Graph{3, 4}
	fmt.Println(g.Abs())

	m := MyFloat(2.0)
	fmt.Println(m.Abs())
}

// like java struct declaration statement, notice := not needed
// "var" can declare one or more variables, in this case, it's declaring more than 1 variable
var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func structTest() {
	fmt.Println("structTest()")

	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(p)
}

func arrayTest() {
	fmt.Println("arrayTest()")

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// array literals
	b := [2]string{"foo", "bar"}
	c := [...]string{"foo", "bar"}

	fmt.Println(b, c, len(b), len(c), cap(b), cap(c))

	// a slice is a view into an array
	fmt.Println("slice: ")
	var s []int = primes[1:4]
	fmt.Println(s)

	// slice literal
	letters := []string{"a", "b", "c", "d"}
	fmt.Println(letters)

	fmt.Println("Any changes into the slice will change the original array and other slices associated with it")

	// build slices independently of an array. array gets built behind the scenes
	var ss []byte
	ss = make([]byte, 5, 5)
	// or
	ss2 := make([]byte, 5)

	fmt.Println(len(ss), len(ss2))

	fmt.Println("slices cannot resize to greater than their capacity")

	// appending to a slice
	var sss []int
	fmt.Println(sss, len(sss), cap(sss))
	sss = append(sss, 0)
	fmt.Println(sss, len(sss), cap(sss))
	fmt.Println("if capacity of array is too small, the backing array gets reallocated and returned")
	sss = append(sss, 1, 2, 3)
	fmt.Println(sss, len(sss), cap(sss))
}

func rangeTest() {
	fmt.Println("rangeTest()")

	nums := []int{2, 3, 4}

	sum := 0
	for _, num := range nums { // replace _ with a variable name to get the index. Use _ when we don't need it
		sum += num
	}
	fmt.Println("sum:", sum)

	// get index and copy of the value
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// iterate over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// iterate over just the keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// iterates over unicode code points
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func mapTest() {
	fmt.Println("mapTest()")

	// use create maps use make()

	var m map[string]string
	m = make(map[string]string)
	m["test"] = "foo"
	fmt.Println(m)

	// map literals
	var mm = map[string]string{
		"hello": "world",
		"foo":   "bar",
	}

	fmt.Println(mm)

	// test key is present with two value assignment
	elem, ok := mm["hello"]
	fmt.Println(elem, ok)
	elem, ok = mm["doesnt exist"]
	fmt.Println(elem, ok) // prints <empty> and false
}

func functionTest() {
	fmt.Println("functionTest()")

	concat := func(s1, s2 string) string {
		return s1 + s2
	}

	fmt.Println(concat("hello", "world"))

	concat3 := func(s1, s2, s3 string) string {
		return s1 + s2 + s3
	}

	fmt.Println(concat("a", concat3("b", "c", "d")))
}

func closureTest() {
	fmt.Println("closureTest()")

	// closures capture variables as reference
	sum := 20
	summer := func(x int) int {
		sum += x
		return sum
	}

	fmt.Println(summer(200))
	fmt.Println(sum)
}

// this is like a typdef and necessary if you want to create member function for float64
// Without this typedef we cannot add member function to float64 because float64 lives in
// a different package
type MyFloat float64

// if m is a pointer, then it modifies the original MyFloat vs making a copy of it
// func (m *MyFloat) Abs() float64 {
func (m MyFloat) Abs() float64 {
	if m < 0 {
		return float64(-1)
	}

	return float64(m)
}

type Vertices struct {
	X, Y float64
}

func (v *Vertices) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Abser is an interface that has one mandatory member function
type Abser interface {
	Abs() float64
}

func interfaceTest() {
	fmt.Println("interfaceTest()")

	var a Abser

	f := MyFloat(3.14)
	a = f

	v := Vertices{3, 4}
	a = &v

	fmt.Println(a.Abs())

	fmt.Println("under the hood interfaces can be thought of as a tuple of value and concrete type")
	fmt.Println("(value, type)")
	fmt.Println("ie (Abs(), MyFloat)")
	fmt.Println("Calling a method on an interface value (Abs()) executes the method of the same name of its underlying type (MyFloat)")
}

func typeAssertions() {
	var i interface{} = "hello"

	// s is underlying string "hello" and ok is true indicating underlying type is indeed string
	s, ok := i.(string)
	fmt.Println(s, ok)

	// test if f is of type float64, ok should be false. Leaving out the ok variable will throw a panic
	f, ok := i.(float64)
	fmt.Println(f, ok)
}

func main() {

	typeAssertions()

	interfaceTest()

	memberFuncTest()

	closureTest()

	functionTest()

	mapTest()

	arrayTest()

	structTest()

	switchStatement()

	loop()
	fmt.Println(conditional())

	insideFunction()

	fmt.Println(i, c, python, java)

	fmt.Println(split(17))

	// := declares a and b, then does an assignment. = only does assignment
	a, b := swap("hello", "world")

	//var b, c string = swap("foo", "bar")

	// prints world hello
	fmt.Println(a, b)
}
