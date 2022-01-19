package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("nico") //우어어어어어!!병렬처리된다!!!!!🤩go
	sexyCount("anne")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
