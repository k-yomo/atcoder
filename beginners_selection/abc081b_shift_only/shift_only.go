package main

import (
	"../../utils"
	"fmt"
)

// https://atcoder.jp/contests/abs/tasks/abc081_b
func main() {
	var length int
	fmt.Scanf("%d", &length)
	nums := utils.ScanNums(length)
	dividedCount := 0
	for {
		oddNumberExist := false
		for i := 0; i < len(nums); i++ {
			if nums[i]%2 != 0 {
				oddNumberExist = true
				break
			} else {
				nums[i] /= 2
			}
		}
		if oddNumberExist {
			break
		}
		dividedCount++
	}
	fmt.Println(dividedCount)
}
