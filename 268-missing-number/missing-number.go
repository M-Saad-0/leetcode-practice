// func missingNumber(nums []int) int {
//     myMap := map[int]int{}
//     for _, k := range nums {
//         myMap[k] = 1
//     }
//     for i:=0; i<=len(nums); i++{
//         _, e := myMap[i]
//         if(!e){
//             return i
//         }
//     }

//     return -1
    
// }

import "sort"
func missingNumber(nums []int) int {
    n := len(nums)
    sort.Sort(sort.IntSlice(nums))
    for i:=0; i<=n; i++ {
        if i == n || i != nums[i] {
            return i
        }
    }
    return -1
}