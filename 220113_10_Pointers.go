package main

import "fmt"

func main() {
	// a := 2
	// b := a //한번만 복사를 하기때문에 2
	// a = 10 //10으로 재지정
	// fmt.Println(a, b) //출력: 10 2

	// a := 2
	// b := 5
	// fmt.Println(&a, &b) //💎& 메모리주소출력 0xc000012078 0xc000012090

	// a := 2
	// b := &a
	// fmt.Println(&a, b) //👌같은메모리주소가출력됩니다~ 0xc000012078 0xc000012078

	// a := 2
	// b := &a
	// fmt.Println(a, b)//2 0xc000012078 b는 계속 a를 지켜보고있다🧐

	// a := 2
	// b := &a
	// fmt.Println((*b)) //⭐*은 살펴보다?훓어보다! 모두연결되어있어 출력:2

	// a := 2
	// b := &a
	// a = 5
	// fmt.Println(*b) //⭐*은 계속 관통해서 보는 느낌이야 !  출력:5

	a := 2
	b := &a
	*b = 20
	fmt.Println(a) //b를 이용해서 a를 변경시킨 경우. 출력20
}
