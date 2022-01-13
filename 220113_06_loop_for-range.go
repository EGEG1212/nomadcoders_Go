package main

import (
	"fmt"
)

func superAdd(numbers ...int) int {
	for index, number := range numbers { //range는 index를 주기때문에 0~5까지 출력됨
		fmt.Println(index, number)

		// 👌 1~6까지 찍고싶다면
		// for i := 0; i < len(numbers); i++ {
		// 	fmt.Println(numbers[i])
	}
	return 1
}

func main() {
	superAdd(1, 2, 3, 4, 5, 6)
}
