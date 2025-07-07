func findDuplicates(nums []int) []int {
    mapo := map[int]int{}
    arr := []int{}
    for _, v := range nums {
        _, exists := mapo[v]
        if exists {
            arr = append(arr, v)
            mapo[v] += 1
        }else {
            mapo[v] = 1
        }
    }
    return arr
}