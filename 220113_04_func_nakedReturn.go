package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (lenght int, uppercase string) {
	lenght = len(name)
	uppercase = strings.ToUpper(name)
	return //πμμμ κ²°κ³Όκ°μ λ€ μ ν΄λ¨μ΄μ returnμ΄ κΉ¨λν΄λ μλ¬μλ€~~
	//μ΄κ±°μ.. naked return π
}

func main() {
	totalLenght, up := lenAndUpper("anne")
	fmt.Println(totalLenght, up)
}
