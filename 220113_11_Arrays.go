package main

import "fmt"

func main() {
	// names := [5]string{"nico", "anne", "joo"} //💎[array의 길의 제시]
	// names[3] = "haha"
	// names[4] = "hoho"
	// names[5] = "good" //길이가 넘어가면 에러남❌
	// fmt.Println((names))

	names := []string{"nico", "anne", "joo"}
	names = append(names, "flynn") //추가만 됩니다~👌
	fmt.Println(names)
}
