package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	defer fmt.Println("I'm done") //💎func종료후 출력문구
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return //naked return
}

func main() {
	totalLenght, up := lenAndUpper("anne")
	fmt.Println(totalLenght, up)
}
