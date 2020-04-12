package main

import (
	"fmt"
)

/*
Write a program which allows the user to create a set
of animals and to get information about those animals.
Each animal has a name and can be either a cow, bird,
or snake. With each command, the user can either create
a new animal of one of the three types, or the user
can request information about an animal that he/she
has already created. Each animal has a unique name,
defined by the user. Note that the user can define
animals of a chosen type, but the types of animals
are restricted to either cow, bird, or snake. The
following table contains the three types of animals
and their associated data.
Animal	Food eaten	Locomtion method	Spoken sound
cow		grass		walk				moo
bird	worms		fly					peep
snake	mice		slither				hsss
Your program should present the user with a
prompt, “>”, to indicate that the user can type a
request. Your program should accept one command at
a time from the user, print out a response, and
print out a new prompt on a new line. Your program
should continue in this loop forever. Every command
from the user must be either a “newanimal” command
or a “query” command.
Each “newanimal” command must be a single line
containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be
the name of the new animal. The third string is the type
of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by
creating the new animal and printing “Created it!” on
the screen.
Each “query” command must be a single line containing
3 strings. The first string is “query”. The second string
is the name of the animal. The third string is the name
of the information requested about the animal, either
“eat”, “move”, or “speak”. Your program should process
each query command by printing out the requested data.
Define an interface type called Animal which describes
the methods of an animal. Specifically, the Animal
interface should contain the methods Eat(), Move(),
and Speak(), which take no arguments and return no
values. The Eat() method should print the animal’s
food, the Move() method should print the animal’s
locomotion, and the Speak() method should print the
animal’s spoken sound. Define three types Cow, Bird,
and Snake. For each of these three types, define
methods Eat(), Move(), and Speak() so that the types
Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of
the appropriate type. Your program should call the
appropriate method when the user issues a query command.
*/

// Animal Interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow Type
type Cow struct {
	food       string
	locomotion string
	noise      string
}

// Bird Type
type Bird struct {
	food       string
	locomotion string
	noise      string
}

// Snake Type
type Snake struct {
	food       string
	locomotion string
	noise      string
}

// Eat print Cow's food
func (c Cow) Eat() {
	fmt.Printf("%s\n", c.food)
}

// Move print Cow's locomotion
func (c Cow) Move() {
	fmt.Printf("%s\n", c.locomotion)
}

// Speak print Cow's spoken sound
func (c Cow) Speak() {
	fmt.Printf("%s\n", c.noise)
}

// Eat print Bird's food
func (b Bird) Eat() {
	fmt.Printf("%s\n", b.food)
}

// Move print Bird's locomotion
func (b Bird) Move() {
	fmt.Printf("%s\n", b.locomotion)
}

// Speak print Bird's spoken sound
func (b Bird) Speak() {
	fmt.Printf("%s\n", b.noise)
}

// Eat print Snake's food
func (s Snake) Eat() {
	fmt.Printf("%s\n", s.food)
}

// Move print Snake's locomotion
func (s Snake) Move() {
	fmt.Printf("%s\n", s.locomotion)
}

// Speak print Snake's spoken sound
func (s Snake) Speak() {
	fmt.Printf("%s\n", s.noise)
}

func assigment4() {
	var command string
	var animalName string
	var animalTypeAction string

	animals := make(map[string]Animal)

	for {
		fmt.Printf(">")
		fmt.Scan(&command, &animalName, &animalTypeAction)

		if command == "newanimal" {
			switch animalTypeAction {
			case "cow":
				a := Cow{"grass", "walk", "moo"}
				animals[animalName] = a
			case "bird":
				a := Bird{"worms", "fly", "peep"}
				animals[animalName] = a
			case "snake":
				a := Snake{"mice", "slither", "hsss"}
				animals[animalName] = a
			}
		} else if command == "query" {
			a, ok := animals[animalName]
			if ok {
				switch animalTypeAction {
				case "eat":
					a.Eat()
				case "move":
					a.Move()
				case "speak":
					a.Speak()
				}
			}
		}
	}
}
