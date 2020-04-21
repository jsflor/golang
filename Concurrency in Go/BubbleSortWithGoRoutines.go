package main

import (
	"fmt"
)

/*
Write a program to sort an array of integers. The program
should partition the array into 4 parts, each of which
is sorted by a different goroutine. Each partition should
be of approximately equal size. Then the main goroutine
should merge the 4 sorted subarrays into one large sorted
array.

The program should prompt the user to input a series of
integers. Each goroutine which sorts ¼ of the array should
print the subarray that it will sort. When sorting is
complete, the main goroutine should print the entire
sorted list.
*/

// Swap func
func Swap(sli []int, i int) {
	if sli[i-1] > sli[i] {
		intermediate := sli[i]
		sli[i] = sli[i-1]
		sli[i-1] = intermediate
	}
}

// BubbleSort func
func BubbleSort(sli []int, c chan []int) {
	for i := len(sli); i > 0; i-- {
		for j := 1; j < i; j++ {
			Swap(sli, j)
		}
	}
	fmt.Println(sli)
	c <- sli
}

func merge(final *[]int, chunks []int) {
	for _, value := range chunks {
		*final = append(*final, value)
	}
}

func assigment3() {
	c := make(chan []int)
	x := []int{}
	y := []int{}
	var aux int

	fmt.Println("Insert 12 integers")
	for i := 0; i <= 12; i++ {
		fmt.Scan(&aux)
		x = append(x, aux)
	}

	var sli1 = x[0:3]
	var sli2 = x[3:6]
	var sli3 = x[6:9]
	var sli4 = x[9:12]

	go BubbleSort(sli1, c)
	go BubbleSort(sli2, c)
	go BubbleSort(sli3, c)
	go BubbleSort(sli4, c)

	part1 := <-c
	part2 := <-c
	part3 := <-c
	part4 := <-c

	merge(&y, part1)
	merge(&y, part2)
	merge(&y, part3)
	merge(&y, part4)

	fmt.Println(y)
}
