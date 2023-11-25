package main

import "fmt"

func main() {
	N := 10

	numbers := make([]int, N)

	for i := 0; i < N; i++ {
		numbers[i] = i + 1
	}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
