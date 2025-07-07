func productExceptSelf(nums []int) []int {

    prev := 1

    tmpArr := make([]int, len(nums))
    for i:=0; i<len(nums);i++ {
        tmpArr[i] = prev 
        prev *= nums[i];
    }
    prev = 1
    for i:=len(nums)-1; i>=0;i-- {
        tmpArr[i] *= prev
        prev *= nums[i]
    }    
    return tmpArr
}



// 1 2 3 4
// 1 1*2 1*2*3 1*2*3*4
// 1 2 6 24
// 1*2*6*24 2*6*24 6*24 24