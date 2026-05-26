package main
import "fmt"

func subsetSum(nums []int, index, target int) bool {
    if target == 0 {
        return true
    }
    if index == len(nums) || target < 0 {
        return false
    }

    return subsetSum(nums, index+1, target-nums[index]) || 
              subsetSum(nums, index+1, target)
}

func main() {
    var n, k int
    fmt.Scan(&n, &k) 

    nums := make([]int, n)
    for i := range nums {
        fmt.Scan(&nums[i])
    }

    if subsetSum(nums, 0, k) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}