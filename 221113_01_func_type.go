package main

import "fmt"

// func multiply(a int, b int) int { //int *3번 타입추가
//func multiply(a, b int) int { // int *2번 타입추가해도 같은 결과.
func multiply(a, b int) int { //💎두번째 int는 결과값의 타입을 나타내는것. 꼭 필요함💎
	return a * b
}

func main() {
	fmt.Println(multiply(2, 2))

}
