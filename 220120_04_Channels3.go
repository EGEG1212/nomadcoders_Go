package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [...]string{"nico", "anne", "jane", "joo", "elsa", "doodle", "groot"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for", i)
		fmt.Println(<-c) //msg를받는다는건 😎blocking operation이다! <-
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	//fmt.Println(person)
	//return true
	c <- person + " is sexy"
}
