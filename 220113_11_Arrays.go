package main

import "fmt"

func main() {
	// names := [5]string{"nico", "anne", "joo"} //๐[array์ ๊ธธ์ ์ ์]
	// names[3] = "haha"
	// names[4] = "hoho"
	// names[5] = "good" //๊ธธ์ด๊ฐ ๋์ด๊ฐ๋ฉด ์๋ฌ๋จโ
	// fmt.Println((names))

	names := []string{"nico", "anne", "joo"}
	names = append(names, "flynn") //์ถ๊ฐ๋ง ๋ฉ๋๋ค~๐
	fmt.Println(names)
}
