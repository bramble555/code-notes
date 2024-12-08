package main

import "fmt"

func getPerBit(n int) []int {
	res := make([]int, 0)

	for n > 0 {
		temp := n % 10
		res = append(res, temp)
		n /= 10
	}
	return res
}

func main() {
	arr := []int{1, 2, 3, 5, 5, 6}
	resCount := 0
	i := 0
	for i < len(arr) {
		flag := true
		for i < len(arr)-1 && arr[i] == arr[i+1] {
			i++
			flag = false
		}
		if flag == false {
			i++
		}
		arr[resCount] = arr[i]
		resCount++
		i++
	}
	fmt.Println(i)
	if i != len(arr) {
		arr[resCount] = arr[i]
		resCount++
		i++
	}
	fmt.Println(arr)
	fmt.Println(resCount)

}
