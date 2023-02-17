package main

import "fmt"

const fizz = "Fizz"
const buzz = "Buzz"

func main() {
	for i := 0; i <= 100; i++ {
		FizzBuzz(i)
	}
}

func FizzBuzz(i int) {
	if i%3 == 0 && i%5 == 0 {
		println(fizz + buzz)
	} else if i%3 == 0 {
		fmt.Println(fizz)
	} else if i%5 == 0 {
		fmt.Println(buzz)
	} else {
		fmt.Println(i)
	}
}
