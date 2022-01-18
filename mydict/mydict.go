package mydict

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	of exists {
		return value, nill
	}
	return "",errNotFound //string도 return해줘야하니까 ""를 포함🎅
}
