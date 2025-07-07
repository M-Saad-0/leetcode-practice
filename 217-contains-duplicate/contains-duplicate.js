/**
 * @param {number[]} nums
 * @return {boolean}
 */
var containsDuplicate = function(nums) {
    const newArr = new Set();
    for(ele of nums){
        if(newArr.has(ele))
        return true
        newArr.add(ele)
    }
    return false ;
};
