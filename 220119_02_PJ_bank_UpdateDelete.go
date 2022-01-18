package main


func main() {
	//🥽Update TEST----------- First를 Second로 바꾸는
	// dictionary := mydict.Dictionary{}
	// baseWord := "hello"
	// dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// word, _ := dictionary.Search(baseWord)
	// fmt.Println(word)
	//🥽Update TEST-----------end
	//🥽Delete TEST-----------더하고 찾아서 지우는 출력: Not Found 나오면 정상 ^^
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(word)
	}
	//🥽Delete TEST-----------end
}
