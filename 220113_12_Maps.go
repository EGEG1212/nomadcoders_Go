package main

import "fmt"

func main() {
	//👇[key type]value type { "key":"value"} 약간 object너낌😎
	nico := map[string]string{"name": "nico", "age": "12"} //string을 사용한다고했기 때문에 나이조차도 12int가 안되고 str으로 들어감 🏃‍♂️
	//type섞어넣는건 struct 구조체 활용 = class + object
	// fmt.Println((nico))  //출력:map[age:12 name:nico]

	for key, value := range nico { //물론 생력가능 _ ignore하면 됨~
		fmt.Println(key, value) //출력: name nico <br> age 12
	}
}
