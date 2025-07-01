//https://leetcode.com/problems/concatenation-of-array/description/
//Array

func getConcatenation(nums []int) []int {
    nums = append(nums, nums...)
    return nums
}