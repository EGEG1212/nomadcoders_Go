package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool) // 👩‍🚀channel만들기!  어떤타입을 보낼것인가 맨아래 return값이 bool이니까~
	people := [2]string{"nico", "anne"}
	for _, person := range people {
		go isSexy(person, c)
	}
	//time.Sleep(time.Second * 10)
	// channel의 msg를 받기위해 메인함수가 기다려줍니다..🤧
	//result := <-c
	//fmt.Println(result) // 여 2줄이 아래의 1줄과 같다 :-)
	fmt.Println(<-c)
	fmt.Println(<-c) //msg2개보냈으니까 2개까지 출력이 가능하다. 3개출력하려고하면 에러남ㅋ.ㅋ

}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	//return true
	c <- true
}
