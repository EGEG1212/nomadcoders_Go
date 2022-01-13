package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name) //strings <<s주의 👈
}

func main() {
	// 👍정상출력. value2개
	// totalLenght, upperName := lenAndUpper("nico")
	// fmt.Println(totalLenght, upperName)

	//🙌두개의 값을 return하는데 왜 하나만 받아가세요??❌에러
	// totalLenght := lenAndUpper("nico")
	// fmt.Println(totalLenght)

	//👍정상출력. value1개
	totalLenght, _ := lenAndUpper("nico") //무시해달라는_표시 ignored value
	fmt.Println(totalLenght)
}
