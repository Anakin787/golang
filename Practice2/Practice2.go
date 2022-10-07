package main

import "fmt"

func main() {
	//for문 실습
	var dan int
	fmt.Scanf("%d", &dan)

	for i := 1; i < 10; i++ {
		fmt.Printf("%d X %d = %d\n", dan, i, dan*i)
	}

	//for문 실습2 - 이등변 삼각형
	var n int

	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if i == j {
				fmt.Print("* ")
			} else {
				fmt.Print("o ")
			}
		}
		fmt.Print("\n")
	}

	//for문과 if/else문 실습 - 7과9의배수(단, 공배수는 한번만 출력)
	for i := 1; i <= 100; i++ {
		if i%7 == 0 {
			fmt.Printf("%d ", i)
		} else if i%9 == 0 {
			fmt.Printf("%d ", i)
		} else if i%7 == 0 && i%9 == 0 {
			fmt.Printf("%d ", i)
		}
	}

	//2단부터9단까지 구구단 - (단,홀수 단만 출력 & n단은 n*n까지 출력)
	var res int
	for i := 2; i <= 9; i++ {
		if i%2 == 0 {
			continue
		}
		for j := 1; j <= i; j++ {
			res = i * j
			fmt.Printf("%d x %d = %d\n", i, j, res)
		}
		fmt.Println("")
	}

	//두 수를 더하면99 - AB + BA = 99가 나오는 모든 수를 출력(if문)
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			if ((i*10)+j)+((10*j)+i) == 99 {
				fmt.Printf("%d%d + %d%d = 99\n", i, j, j, i)
			}
		}
	}

	//슬라이스는 배열과 다르게 값을 복사해오는 것이 아니라 슬라이스 자체가 참조하고있는 주소값을 같이 참조하는 것을 의미

	
}
