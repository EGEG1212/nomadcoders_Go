package main

//😭 mydict 왜import안돼!?

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	//dictionary["hello"] = "hello"
	//dictionary.Add() 
	//dictionary.Delete()
	//dictionary.Search()
	//fmt.Println(dictionary["first"]) //이렇게해도 출력: Firsr word 나오긴하는데 멋진코드는 아니다🤐
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
