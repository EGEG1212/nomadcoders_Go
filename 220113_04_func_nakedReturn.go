package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return //👆위에서 결과값을 다 정해놨어서 return이 깨끗해도 에러없다~~
	//이거시.. naked return 🛀
}

func main() {
	totalLenght, up := lenAndUpper("anne")
	fmt.Println(totalLenght, up)
}
