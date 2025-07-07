
import "slices"
func twoSum(nums []int, target int) []int {
    for i:=0; i<len(nums); i++ {
        var number = target - nums[i]
        var exists = slices.Index(nums, number)
        if exists != -1 && i != exists{
            return []int{i, exists}
        }
        }       
    
    return []int{}
}