package main

//π­ mydict μimportμλΌ!?

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	//dictionary["hello"] = "hello"
	//dictionary.Add() 
	//dictionary.Delete()
	//dictionary.Search()
	//fmt.Println(dictionary["first"]) //μ΄λ κ²ν΄λ μΆλ ₯: Firsr word λμ€κΈ΄νλλ° λ©μ§μ½λλ μλλ€π€
	definition, err := dictionary.Search("second")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
