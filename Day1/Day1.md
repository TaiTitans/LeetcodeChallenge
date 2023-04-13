# Day 1

- [LeetCode-RunningofArray](https://leetcode.com/problems/running-sum-of-1d-array/)
---

## Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

## Return the running sum of nums.

---
### Example 1:

```sh
Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
```

---

### Let do it :

```sh
std::vector<int> runningSum(std::vector<int>& nums) {
    int sum = 0;
    for (int i = 0; i < nums.size(); i++) {
        sum += nums[i]; // Add the current element of nums to the running sum
        nums[i] = sum; // Update the current element of nums with the running sum
    }
    return nums;
}
```

### FULL CODE :

```sh
#include <iostream>
#include <vector>

std::vector<int> runningSum(std::vector<int>& nums) {
    int sum = 0;
    for (int i = 0; i < nums.size(); i++) {
        sum += nums[i]; // Add the current element of nums to the running sum
        nums[i] = sum; // Update the current element of nums with the running sum
    }
    return nums;
}

int main(){
    int n;
    std::cout << "Enter the size of the array: ";
    std::cin >> n;

    std::vector<int> nums;
    for (int i = 0; i < n; i++) {
        nums.push_back(i + 1); // Populate nums with elements from 1 to n
    }

    std::vector<int> result = runningSum(nums);

    std::cout << "Running Sum of the Array: ";
    for (int j = 0; j < result.size(); j++) {
        std::cout << result[j] << " ";
    }

    return 0;
}
```
---