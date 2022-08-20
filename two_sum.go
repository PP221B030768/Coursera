// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

// golang
package main

import (
	"fmt"
	"strconv"
)

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums) - 1; i++{
        num := target - nums[i]
        for j := i + 1; j < len(nums); j++{
            if num == nums[j]{
                if i > j{
                    return []int{j, i}
                }else{
                    return []int{i, j}
                }
            }
        }
    }
    return []int{}
}

func main(){
	nums := []int{1, 2, 3}
	target := 5
	fmt.Println(twoSum(nums, target))
}

// JS
// var twoSum = function(nums, target) {
//     for (var i = 0; i < nums.length; i++){
//         var num = target - nums[i];
//         for (var j = i + 1; j < nums.length; ++j){
//             if (nums[j] == num){
//                 return [i, j];
//             }
//         }
//     }
// };

// Python
// class Solution(object):
//     def twoSum(self, nums, target):
//         """
//         :type nums: List[int]
//         :type target: int
//         :rtype: List[int]
//         """
//         for i in range(0, len(nums) - 1):
//             for j in range(i + 1, len(nums)):
//                 if nums[i] + nums[j] == target:
//                     return [i, j]
        
// C++
// class Solution {
// 	public:
// 		vector<int> twoSum(vector<int>& nums, int target) {
// 			vector<int> sum;
// 			for (int i = 0; i < nums.size() - 1; i++){
// 				for (int j = i + 1; j < nums.size(); j++){
// 					if (nums[i] + nums[j] == target){
// 						sum.push_back(i);
// 						sum.push_back(j);
// 					}
// 				}
// 			}
// 			return sum;
// 		}
// 	};