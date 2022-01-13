package main

import "fmt"

func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { //if 바로 뒤에 variable생성가능💎
		return false
	}
	return true //else일경우
}

func main() {
	fmt.Println(canIDrink(16))
}
