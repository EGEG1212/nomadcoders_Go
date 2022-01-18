package mydict

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound = errors.New("Not Found")
	errCantUpdate = errors.New("Cant update non-existing word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	of exists {
		return value, nill
	}
	return "",errNotFound //string도 return해줘야하니까 ""를 포함🎅
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error { //📢word, def 모두 string이고, error를 return해줄꺼야
	_, err := d.Search(word) // 생략부분은 Search의 definition. Search의 에러를 받아 단어가 없다면!? dict에 단어를 추가가능하게 해주지~
	// if err == errNotFound{
	// 	d[word] = def
	// } else if err ==nil {
	// 	return errWordExists
	// }
	// return nil
	//-------------- 위와 같은 내용, 아래는 switch를 써서~~
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
	//---------------
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word) //이미존재하는 단어여야 update가 가능하니까.
	switch err {
	case nil: //단어가 있다면
		d[word] = definition
	case errNotFound://만약 없는 단어라면,
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}