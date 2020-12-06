package main

import "fmt"

func main() {
	nums := []int {3,4,8}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	fmt.Println("~~~~||||~~~~")

	for i, num := range nums {
        if num % 2 == 0 {
            fmt.Println("even index:", i, "|| num:", num)
        }
    }
}