package main

import (
	"fmt"
)

// ✨점3개로 모든arguments를 받아올수있다! [array]형태로 출력됨
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("naco", "anne", "joo", "dal", "marl")
}
