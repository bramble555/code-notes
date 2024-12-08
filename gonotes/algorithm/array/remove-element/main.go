package main
// le27
import "fmt"

// 快慢指针
// 快指针用来寻找，慢指针用来存储新数组
func removeElement(nums []int, val int) int {
	var slow int
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func main() {
	nums := []int{3, 2, 2, 3}
	fmt.Println(removeElement(nums, 3))
	fmt.Println(nums)

}
