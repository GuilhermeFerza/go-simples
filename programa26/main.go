package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	alvo := 4

	fmt.Println(twoSum(nums, alvo))
	fmt.Println(twoSumVanilla(nums, alvo))
}

func twoSum(nums []int, alvo int)[]int{
	seen := make(map[int]int)
	for i, num := range nums{
		complement := alvo - num
		if idx, ok := seen[complement]; ok{
			return []int{idx, i}
		}
		seen[num]=i
	}
	return nil
}

func twoSumVanilla(nums []int, target int) []int{
	for i := 0; i < len(nums); i++{
		for j := i; j< len(nums); j++{
			if nums[i]+nums[j] == target{
				return []int{i, j}
			}
		}
	}
	return nil
}