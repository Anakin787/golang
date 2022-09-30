package main

import (
	"fmt"
)

/*변수는 함수밖에서는 var 변수명 변수형=값으로 선언하고, 안에서는 변수명:=값으로도 선언가능하다.
상수는 초기화후 변경이 불가능하기때문에 선언만 할 순 없고, 선언과 동시에 초기화를 해주어야한다.
변수는 초기화 후 사용하지 않으면 에러가 발생한다.*/
const username = "kim"

func main() {
	var ji int = 1
	const a int = 1
	const b, d = 10, 23
	const c = "groom"
	fmt.Println(ji, a, b, c, d)
	fmt.Println(username)
	const (
		x = "io"
		y        //선언만하면 위의 값을 그대로 가져온다
		z = iota //iota는 변수의 index값을 가져온다
	)
	fmt.Println(x, y, z)
}
