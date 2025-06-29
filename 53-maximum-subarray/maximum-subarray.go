func maxSubArray(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    sum := 0
    maxSum := nums[0]
    for i:=0; i<len(nums); i++ {
        sum += nums[i]
        if maxSum < sum {
            maxSum = sum
        }
        if sum < 0 {
            sum = 0
        }
          
    }
    fmt.Print(sum)

    return maxSum
}