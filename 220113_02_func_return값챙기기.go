package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name) //strings <<sμ£Όμ π
}

func main() {
	// πμ μμΆλ ₯. value2κ°
	// totalLenght, upperName := lenAndUpper("nico")
	// fmt.Println(totalLenght, upperName)

	//πλκ°μ κ°μ returnνλλ° μ νλλ§ λ°μκ°μΈμ??βμλ¬
	// totalLenght := lenAndUpper("nico")
	// fmt.Println(totalLenght)

	//πμ μμΆλ ₯. value1κ°
	totalLenght, _ := lenAndUpper("nico") //λ¬΄μν΄λ¬λΌλ_νμ ignored value
	fmt.Println(totalLenght)
}
