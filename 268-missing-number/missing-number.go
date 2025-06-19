func missingNumber(nums []int) int {
    myMap := map[int]int{}
    for _, k := range nums {
        myMap[k] = 1
    }
    for i:=0; i<=len(nums); i++{
        _, e := myMap[i]
        if(!e){
            return i
        }
    }

    return -1
    
}