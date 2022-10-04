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

	//출력함수실습
	fmt.Printf("%-8s%-14s%5s\n", "이름", "전공학과", "학년")
	fmt.Printf("%-8s%-14s%5s\n", "유현수", "전자공학", "3")
	fmt.Printf("%-8s%-14s%5s\n", "김윤욱", "컴퓨터공학", "4")
	fmt.Printf("%-8s%-14s%5s\n", "김나영", "미술교육학", "2")

	//입력함수실습
	var RRNf, RRNt int
	var name string
	var height float32

	fmt.Scanf("%d-%d", &RRNf, &RRNt)
	fmt.Scanf("%s", &name)
	fmt.Scanf("%f", &height)

	fmt.Printf("주민등록번호 앞자리는 %d, 뒷자리는 %d, 이름은 %s입니다.\n", RRNf, RRNt, name)
	fmt.Printf("그리고 키는 %.2f입니다.", height)

}

// go에서는 증감연산자에서 후위연산자만 된다.
// = 변수나 상수에 값을 대입한다.
// 변수는 변수끼리 대입이 가능합니다.
// := 변수를 선언 및 대입한다.
// Go언어에서는 오로지 'true'와 'false'만 사용하여 할당
/* string으로 선언한 문자열 타입은 immutable 타입으로서 값을 수정할 수 없습니다.
=> var str string = "hello"와 같이 선언하고 str[2] = 'a'로 수정이 불가능합니다.*/
// 문자열을 표현할 때 백틱(``)을 사용하면 전부 문자열로 인식하고, 큰따옴표("")를 쓰면 개행문자나 %d같은 걸 인식합니다.
/* Go언어에서는 형 변환을 할 때 변환을 명시적으로 지정해주어야합니다.
만약 명시적인 지정이 없다면 런타임 에러가 발생합니다.*/
// Go언어에서는 배열을 한번에 출력이 가능 ex) var arr [3]int = [5]int{1, 2, 3}일 때  fmt.Printf("%d", arr)이 가능
