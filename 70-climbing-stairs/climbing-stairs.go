func climbStairs(n int) int {
     if n == 0 || n == 1 {
        return 1
    }
    a:=1; b:=1
    for i:=2;i<=n;i++{
        b, a =a, a+b
    }
    return a

}

// func climbStairs(n int) int {
//     if n == 0 || n == 1 {
//         return 1
//     }
//     return climbStairs(n-1) + climbStairs(n-2)
// }

