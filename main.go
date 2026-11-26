package main

import (
	"fmt"

	"github.com/grachmannico95/treasury-coding-test/question1"
	"github.com/grachmannico95/treasury-coding-test/question2"
	"github.com/grachmannico95/treasury-coding-test/question3"
	"github.com/grachmannico95/treasury-coding-test/question4"
	"github.com/grachmannico95/treasury-coding-test/question5"
)

func main() {
	fmt.Println("question 1:")
	fmt.Println(question1.GO([]int{-1, 1}))
	fmt.Println(question1.GO([]int{-1, -7, -5}))
	fmt.Println(question1.GO([]int{1, 2, 1, 6}))

	fmt.Println("")
	fmt.Println("question 2:")
	fmt.Println(question2.GO([]int{0, 1, 0}, 1000, 1010))
	fmt.Println(question2.GO([]int{0, 1, 0}, 1259, 2601))

	fmt.Println("")
	fmt.Println("question 3:")
	fmt.Println(question3.GO([]int{1, 2}, []int{1, 3}))
	fmt.Println(question3.GO([]int{1, 2, 2}, []int{1, 2, 4}))

	fmt.Println("")
	fmt.Println("question 4:")
	fmt.Println(question4.GO(1223334))

	fmt.Println("")
	fmt.Println("question 5:")
	fmt.Println(question5.GO(1, 2, 3, 4))
	fmt.Println(question5.GO(1, 2, 3, 4, 5))
}
