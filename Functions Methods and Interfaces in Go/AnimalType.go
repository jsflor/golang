package main

import (
	"fmt"
)

/*
Write a program which allows the user to get information about a
predefined set of animals. Three animals are predefined, cow, bird,
and snake. Each animal can eat, move, and speak. The user can issue
a request to find out one of three things about an animal:
1) the food that it eats, 2) its method of locomotion, and
3) the sound it makes when it speaks. The following table
contains the three animals and their associated data which
should be hard-coded into your program.
Animal	Food eaten	Locomotion method	Spoken sound
cow		grass		walk				moo
bird	worms		fly					peep
snake	mice		slither				hsss
Your program should present the user with a prompt, “>”,
to indicate that the user can type a request. Your program
accepts one request at a time from the user, prints out the
answer to the request, and prints out a new prompt. Your
program should continue in this loop forever. Every request
from the user must be a single line containing 2 strings. The
first string is the name of an animal, either “cow”, “bird”,
or “snake”. The second string is the name of the information
requested about the animal, either “eat”, “move”, or “speak”.
Your program should process each request by printing out the
requested data.
You will need a data structure to hold the information about
each animal. Make a type called Animal which is a struct
containing three fields:food, locomotion, and noise, all of
which are strings. Make three methods called Eat(), Move(),
and Speak(). The receiver type of all of your methods should
be your Animal type. The Eat() method should print the animal’s
food, the Move() method should print the animal’s locomotion,
and the Speak() method should print the animal’s spoken sound.
Your program should call the appropriate method when the
user makes a request.
*/

// Animal type
type animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat print animal's food
func (a animal) Eat() {
	fmt.Printf("%s\n", a.food)
}

// Move print animal's locomotion
func (a animal) Move() {
	fmt.Printf("%s\n", a.locomotion)
}

// Speak print animal's spoken sound
func (a animal) Speak() {
	fmt.Printf("%s\n", a.noise)
}

func assigment3() {
	cow := animal{"grass", "walk", "moo"}
	bird := animal{"worms", "fly", "peep"}
	snake := animal{"mice", "slither", "hsss"}

	var animalName string
	var infoRequested string

	for {
		fmt.Printf(">")
		fmt.Scan(&animalName, &infoRequested)

		switch animalName {
		case "cow":
			switch infoRequested {
			case "eat":
				cow.Eat()
			case "move":
				cow.Move()
			case "speak":
				cow.Speak()
			}
		case "bird":
			switch infoRequested {
			case "eat":
				bird.Eat()
			case "move":
				bird.Move()
			case "speak":
				bird.Speak()
			}
		case "snake":
			switch infoRequested {
			case "eat":
				snake.Eat()
			case "move":
				snake.Move()
			case "speak":
				snake.Speak()
			}
		}
	}
}
