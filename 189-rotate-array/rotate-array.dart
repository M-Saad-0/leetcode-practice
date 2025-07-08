class Solution {
  void rotate(List<int> nums, int k) {
    final int n = nums.length;
    k = k % n; // Fix: to handle cases where k > n

    
      reverse(nums, 0, n-1);
      reverse(nums, 0, k-1);
      reverse(nums, k, n-1);
    

    print(nums);
  }

  void reverse(List<int> nums, int begin, int end) {
    var tmp = 0;
    while(begin<end){
        tmp = nums[begin];
        nums[begin] = nums[end];
        nums[end] = tmp;
        begin++;
        end--;
    }
  }
}