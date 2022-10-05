package main

import "fmt"

func bubbleSort(nums []int) []int { //nums가 어레이(슬라이스)형식이라 []int이렇게 붙인것
	var temp int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp = nums[j]
				nums[j] = nums[i]
				nums[i] = temp
			}
		}
	}
	return nums //bubbleSort가 수행되어 nums로 반환
}

func inputNums() (nums []int) { //nums가 슬라이스 형식이라 []int형식으로 반환
	var num, num2 int
	fmt.Scanln(&num)
	for i := 0; i < num; i++ {
		fmt.Scanln(&num2)
		nums = append(nums, num2)
	}
	return
}

func outputNums(nums []int) {
	for i := 0; i < len(nums); i++ {
		fmt.Printf("%d ", nums[i])
	}
}

func main() {
	nums := inputNums() //nums라는 inputNums함수의 반환값을 nums라는 변수에 저장되는 꼴
	bubbleSort(nums)
	outputNums(nums)
}
