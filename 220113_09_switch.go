package main

import "fmt"

func canIDrink(age int) bool {
	switch { // π{μμ :=λ³μμ μΈκ°λ₯
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}
	// switch age {
	// case 10:
	// 	return false
	// case 18:
	// 	return true
	// }
	return false
}

func main() {
	fmt.Println(canIDrink(51))
}
