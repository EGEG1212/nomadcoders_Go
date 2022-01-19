package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request failed")

func main() {
	//var results = map[string]string{} //empty map을 만들고싶어서 =, {} 두가지 추가해주면됨
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	for _, url := range urls { //_생략한건 인덱스😎
		//fmt.Println(url)
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	//fmt.Println(results) //옆으로줄줄이나오는게 보기싫어서 아래와같이 줄바꿈!😋
	for url, result := range results {
		fmt.Println(url, result)
	}

}

//go lang std library 검색ㄱㄱ
func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode) //에러의이유가궁금하니까!내용을찍어보자
		return errRequestFailed
	}
	return nil
}
