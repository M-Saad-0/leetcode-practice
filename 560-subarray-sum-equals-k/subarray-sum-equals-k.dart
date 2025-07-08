class Solution {
  int subarraySum(List<int> nums, int k) {
    int sum = 0;
    int startIndex = 0;
    int totalSubArrays = 0;
    Map<int, int> mapo = {0 : 1};
    for(int i=0; i<nums.length; i++){
        sum += nums[i];
        if(mapo.containsKey(sum-k)) {
            totalSubArrays += mapo[sum-k]!;
        }
        mapo[sum] = (mapo[sum] ?? 0) + 1;
    }
        return totalSubArrays;
  }
}