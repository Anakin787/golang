// 슬라이스 형태의 매개변수를 전달할 때 "함수명(변수명...)"으로 보내고 받는 매개변수에서는 "변수명 ...반환형"으로적는다.

//익명함수는 클로저,defer,Go루틴에서 많이 쓰인다.

// 선언함수와 익명함수는 반환값을 변수에 초기화 할 수 있는데, 익명함수는 변수를 함수의 이름처럼 사용가능하다.

// 선언함수보다 익명함수가 나중에 읽혀서,  익명함수가 선언함수를 가리는것과 같은 개념.

// 익명함수를 매개변수로 보내 선언함수에서 받을때
package main

import "fmt"

type multis func(int, int) int

func calc(f multis, a int, b int) int { //받는 형식 유심히 보자 & type으로 함수형식 보기좋게 축약
	result := f(a, b)
	return result
}

func main() {
	multi := func(i int, j int) int {
		return i * j
	}

	r1 := calc(multi, 10, 20)
	fmt.Println(r1)

	r2 := calc(func(x int, y int) int { return x + y }, 10, 20)
	fmt.Println(r2)
}

// 선언함수의 구조체를 메인함수에서 사용하기 위해서 "변수=new(패키지.구조체)"로 선언한다.

//생성자 함수는 호출하면 구조체 객체 생성 및 초기화, 입력한 필드 생성 및 초기화함과 동시에 구조체를 반환합니다.
type mapStruct struct { //mapStruct라는 구조체를 생성
	data map[int]string //map형태의 data필드생성
}

func newStruct() *mapStruct { //포인터 구조체(형)를 반환함
	d := mapStruct{}          //mapStruct구조체를 변수d에 저장(=구조체 생성 및 초기화)
	d.data = map[int]string{} //data필드생성 및 map초기화
	return &d                 //변수d를 참조하기위해 &를 붙여 반환
}

/*
new 키워드로 선언한 변수는 포인터 형이다. ex) wait:=new(w.wait)
	=> 보내는 매개변수로 사용할 때 &를 사용하지 않아도 자동으로 주소를 참조한다.
	=> 받는 매개변수로 사용할 때 *를 붙여 사용한다. ex) *w.wait 	*/

// 메인함수에서 파라미터로 구조체를 전달할때 외부함수의 파라미터에 (변수명 *패키지.구조체)으로 작성한다.

// 클로저는 익명함수를 감싸고있는 함수의 변수를 가져다 쓸 수 있다.

// 클로저는 main루틴에서 생성한 채널들을 매개변수 없이 접근한다. / 함수를 함수의 반환형으로 사용하는것.

// 채널생성 예시) result := make(chan int)

// range문에서는 채널에 송신한 데이터의 개수만큼 접근.

/*
range는 송신채널이 닫히지 않았다면 데이터를 계속받아서 데드락이 발생한다.
	=>채널이 닫힌후 사용해야한다.
	=>range뒤에는채널수신자(<-)를 사용하는것이 아니라 채널이름만 써야한다.	*/

// 수신전용 채널에서 채널을 반환할 때 그 함수의 반환형에 '<-chan 반환형'의 형식으로 반환하고,

// 메인함수에서반환값을 변수에 담아  사용할땐 '<-변수명'을 사용한다.

// 채널은 병렬처리중에서 루틴의 흐름을 제어하기 위해 쓴다.
