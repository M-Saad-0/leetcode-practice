func search(nums []int, target int) int {
  var n int = len(nums)
  if n==0 {
      return -1
  }
  var midIndex int = n/2
  var mid int = nums[midIndex]
  if mid == target {
      return midIndex
  }
  if mid > target {
      return search(nums[0:midIndex], target)
  } else {
      res := search(nums[midIndex+1:n], target)
      if res == -1 {
          return res
      }
      return midIndex + 1 + res
  }
}