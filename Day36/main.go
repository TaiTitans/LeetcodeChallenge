package main

import "fmt"

func minimumOperations(nums []int) int {
    count := 0
    for _, n := range nums { 
        if n % 3 != 0 {
            count++
        }
    }
    return count
}


func main() {
	nums := []int{1,2,3,4}
	fmt.Println("Result: ",minimumOperations(nums))
}
