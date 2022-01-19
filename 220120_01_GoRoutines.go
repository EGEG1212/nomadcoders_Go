package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("nico") //🤩go는 1개만 ㅋㅋ(고루틴은 프로그램이 작동하는 동안만 유효하다. == 메인함수가 실행되는 동안만!)
	sexyCount("anne")
	time.Sleep(time.Second * 5) //고루틴이 5초만살아있고 이후엔 메인함수가 종료됨
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
