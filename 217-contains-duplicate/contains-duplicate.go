func containsDuplicate(nums []int) bool {
    uniques := map[int]int{}
    for _, v:=range nums {
        _, exist := uniques[v] 
        if exist{
            uniques[v]+=1
        }else {
            uniques[v]=1
        }
    }

    for _, v := range uniques {
        if v > 1 {
            return true
        }
    }
    return false
}