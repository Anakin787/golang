package main

import "fmt"

//슬라이스는 배열과 다르게 값을 복사해오는 것이 아니라 슬라이스 자체가 참조하고있는 주소값을 같이 참조하는 것을 의미
func main() {
	// names := make([]string, 0)

	// var name string

	// for {
	// 	fmt.Scanf("%s\n", &name)
	// 	if name == "1" {
	// 		break
	// 	}
	// 	names = append(names, name)
	// }

	// var result string = names[0]

	// for i := 0; i < len(names); i++ {
	// 	if len(result) < len(names[i]) {
	// 		result = names[i]
	// 	} else {
	// 		continue
	// 	}
	// }

	// fmt.Println(result, len(result))

	//map을이용한 평균점수(난이도 上)
	a := make(map[string]int)
	var b string
	var c, sum int
	var avg float32 = 0

	for {
		fmt.Scanf("%s %d\n", &b, &c)
		if b == "0" {
			break
		}
		a[b] = c
		sum += c
	}

	avg = float32(sum) / float32(len(a))

	for sub, scr := range a {
		fmt.Printf("%s %d\n", sub, scr)
	}

	fmt.Printf("%.2f", avg)

}
