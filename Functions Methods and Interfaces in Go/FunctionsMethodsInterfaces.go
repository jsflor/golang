package main

import (
	"fmt"
	"os"
)

// Call by Value
func foo(x, y int) int {
	return x * y
}

// Double return
func foo2(x int) (int, int) {
	return x, x + 1
}

// Call by Reference
func foo3(x *int) {
	*x = *x + 5
}

//Passing Array Pointer
func foo4(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}

//Passing Slices
func foo5(sli []int) {
	sli[0] = sli[0] + 1
}

// Variables as Functions
var funcVar func(int) int

func increment(x int) int {
	return x + 1
}

func decrement(x int) int {
	return x - 1
}

// Functions as Arguments
func applyIt(myfunc func(int) int, x int) int {
	return myfunc(x)
}

// Variable Argument Number
func getMax(values ...int) int {
	maxValue := -1
	for _, v := range values {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}

/*	Return Function
	func double(x, y ,z int) func (a) int {
		fn := func (int a) int {
			return a + x * y * z
		}
	}
*/

//Associating Methods with Data && Structs with Methos
type person struct {
	name string
	age  int
}

//Passing by Value
func (p person) greetings() {
	fmt.Printf("Welcome %s!!!\n", p.name)
}

//Passing by Reference with Pointer Receiver
func (p *person) changeName(n string) {
	p.name = n
}

// AnimalType Interface for a function which takes
// multiples types of parameters
type AnimalType interface {
	Speak()
}

// Dog type for animal
type Dog struct {
	name string
}

// Speak sound of a dog
func (d Dog) Speak() {
	fmt.Printf("Wowf!\n")
}

// Cat type for animal
type Cat struct {
	name string
}

// Speak sound of a cat
func (c Cat) Speak() {
	fmt.Printf("Miauu!\n")
}

func getSound(a AnimalType) {
	a.Speak()
}

// Empty Interface with Type Assertions
func add(val interface{}) {
	intNumber, ok := val.(int)
	if ok {
		fmt.Printf("Integer %v\n", intNumber)
	}
	floatNumber, ok := val.(float64)
	if ok {
		fmt.Printf("Float %v\n", floatNumber)
	}
}

func main() {

	defer fmt.Println("Bye")
	fmt.Println("Hello")

	x := foo(2, 3)
	fmt.Println(x)

	y, z := foo2(5)
	fmt.Println(y, z)

	w := 5
	foo3(&w)
	fmt.Println(w)

	a := [3]int{1, 2, 3}
	foo4(&a)
	fmt.Println(a)

	b := []int{1, 2, 3}
	foo5(b)
	fmt.Println(a)

	funcVar := increment
	fmt.Println(funcVar(1))

	fmt.Println(applyIt(decrement, 5))

	anonymous := applyIt(func(x int) int { return x * 2 }, 5)
	fmt.Println(anonymous)

	fmt.Println(getMax(1, 2, 3))
	vslice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(getMax(vslice...))

	me := person{"sebas", 22}
	me.greetings()
	fmt.Println(me)
	me.changeName("Juan")
	fmt.Println(me)

	perro := Dog{"Brian the dog"}
	gato := Cat{"Garfield the cat"}

	getSound(perro)
	getSound(gato)

	add(5)
	add(6.2)

	// Handling Errors
	f, err := os.Open("/harris/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f)
}
