package main


func main() {
	//dictionary["hello"] = "hello"
	//dictionary.Add() 
	//dictionary.Delete()
	//dictionary.Search()
	//fmt.Println(dictionary["first"]) //이렇게해도 출력: Firsr word 나오긴하는데 멋진코드는 아니다🤐
	//🥽Search TEST-----------
	// dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }
	//🥽Search TEST-----------end
	//🥽Add TEST-----------
	dictionary := mydict.Dictionary{}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}
	hello, _ := dictionary.Search(word) //생략한것은 err
	fmt.Println("found", word, "definition:", hello)
	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
	//🥽Add TEST-----------end
}
