package main

import "fmt"

// Weekdays typedef for days of the week
type Weekdays int

// Weekdays enum
const (
	Monday Weekdays = iota
	Tuesday
	Wednesday
	Thursday
	Friday
)

func printDays(day Weekdays) {

	switch day {
	case 0:
		fmt.Printf("Monday")
	case 1:
		fmt.Printf("Tuesday")
	case 2:
		fmt.Printf("Wednesday")
	case 3:
		fmt.Printf("Thursday")
	case 4:
		fmt.Printf("Friday")
	}
	fmt.Printf("\n")

}

func typeDeclarations() {

	type ID int
	type Name string

	var dni ID = 54361019
	var myName Name = "Sebastian"

	fmt.Printf("typeDeclarations %s %v", myName, dni)
	fmt.Printf("\n")

}

func initializingVariables() {

	var x int = 100
	var y = 100

	var z int // z = 0
	z = 100

	w := 100

	fmt.Printf("initializingVariables %v %v %v %v", x, y, z, w)
	fmt.Printf("\n")

}

func pointers() {

	var x int = 1
	var y int
	var ip *int // ip is pointer to int

	ip = &x // ip now point to x
	y = *ip

	ptr := new(int)
	*ptr = 3

	fmt.Printf("pointers %v, %v", y, *ptr)
	fmt.Printf("\n")

}

func casting() {
	var x int32 = 1
	var y int16 = 2

	x = int32(y)

	fmt.Printf("casting %v", x)
	fmt.Printf("\n")
}

func flowControl() {
	x := 10

	fmt.Printf("flowControl")

	if x > 5 {
		fmt.Printf("greater")
	} else {
		fmt.Printf("lower")
	}

	fmt.Printf("\n")

	for i := 0; i < 10; i++ {
		fmt.Printf("%v", i)
	}

	fmt.Printf("\n")

	var i int = 10
	for i > 0 {
		fmt.Printf("%v", i)
		i--
	}
}

func scan() {
	var age int

	fmt.Printf("Age?")
	fmt.Scan(&age)
	fmt.Printf("Your age is %v", age)
}

func arrays() {

	var x [5]int
	x[0] = 2

	var y [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(x[0])
	fmt.Println(y[3])

	z := [...]int{1, 2, 3, 4, 5}

	for i, v := range z {
		fmt.Printf("%v %v", i, v)
		fmt.Printf("\n")
	}

}

func slices() {

	arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}

	s1 := arr[1:3]
	s2 := arr[2:5]

	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))

	sli := []int{1, 2, 3}
	fmt.Println(sli[1])

	sli2 := make([]int, 0, 3)
	sli2 = append(sli2, 100)
	fmt.Println(sli2[0])
	fmt.Println(len(sli2))

}

func maps() {

	//var idMap map[string]int
	//idMap = make(map[string]int)

	//idMap := make(map[string]int)

	idMap := map[string]int{"joe": 123}

	idMap["sebas"] = 10
	delete(idMap, "sebas")

	id, p := idMap["joe"]

	if p {
		fmt.Println(id)
	} else {
		fmt.Println(0)
	}

	fmt.Println(idMap["sebas"])

	for key, val := range idMap {
		fmt.Println(key, val)
	}

}

func structs() {

	type Person struct {
		name  string
		add   string
		phone int
	}

	var p1 Person
	p1.name = "Sebas"
	p1.add = "Moralzarzal"
	p1.phone = 666666

	//p2 := new(Person)

	//p3 := Person(name: "Ana", add: "Villalba", phone: 123)

	fmt.Printf("%s lives in %s and his phone is %v", p1.name, p1.add, p1.phone)

}

func main() {

	typeDeclarations()
	initializingVariables()
	pointers()
	casting()
	flowControl()
	printDays(Monday)
	arrays()
	slices()
	maps()
	structs()
	scan()

}
