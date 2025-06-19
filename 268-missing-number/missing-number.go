//Slow
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


//Slower


// import "sort"
// func missingNumber(nums []int) int {
//     n := len(nums)
//     c := 0
//     sort.Sort(sort.IntSlice(nums))
//     for c<n {
//         if c != nums[c] {
//             return c
//         }
//         c++
//     }
//     return n
// }


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

func missingNumber(nums []int) int {
    n := len(nums)
    expectedSum := (n * (n + 1))/2
    actualSum := 0
    for i:=0; i<n; i++ {
        actualSum += nums[i]
    }
    return expectedSum - actualSum
}