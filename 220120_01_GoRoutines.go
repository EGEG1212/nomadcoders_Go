package main

import (
	"fmt"
	"time"
)

func main() {
	go sexyCount("nico") //π€©goλ 1κ°λ§ γγ(κ³ λ£¨ν΄μ νλ‘κ·Έλ¨μ΄ μλνλ λμλ§ μ ν¨νλ€. == λ©μΈν¨μκ° μ€νλλ λμλ§!)
	sexyCount("anne")
	time.Sleep(time.Second * 5) //κ³ λ£¨ν΄μ΄ 5μ΄λ§μ΄μμκ³  μ΄νμ λ©μΈν¨μκ° μ’λ£λ¨
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
