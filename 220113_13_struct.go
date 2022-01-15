package main

import "fmt"

//type섞어넣는건 struct 구조체 활용 = class + object
type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi", "ramen"}
	// nico := person{"nico", 18, favFood} // 위 struct를 확인해야하기 때문에 강추하는 코드는 아니다.
	nico := person{name: "anne", age: 14, favFood: favFood}
	fmt.Println(nico)         //출력 {nico 18 [kimchi ramen]}
	fmt.Println(nico.favFood) //당연히 이렇게도 출력가능
}
