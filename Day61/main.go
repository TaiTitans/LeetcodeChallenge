package main

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
    sort.Ints(arr)
    ans := 1
    for i := 1; i < len(arr); i++ {
        if arr[i] > ans+1 {
            ans++
        } else {
            ans = arr[i]
        }
    }
    return ans
}