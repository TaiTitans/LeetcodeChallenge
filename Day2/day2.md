# Day 2

- [LeetCode-RichestCustomerWealth](https://leetcode.com/problems/richest-customer-wealth/)
---

## You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.

## A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.

---
### Example 1:

```sh
Input: accounts = [[1,2,3],[3,2,1]]
Output: 6
Explanation:
1st customer has wealth = 1 + 2 + 3 = 6
2nd customer has wealth = 3 + 2 + 1 = 6
Both customers are considered the richest with a wealth of 6 each, so return 6.
```

---

### Let do it :

```sh
int maximumWealth(std::vector<std::vector<int>>& accounts) {
    int maxWealth = 0; // variable to store the maximum wealth
    for (int i = 0; i < accounts.size(); ++i) { // iterate through each customer
        int wealth = 0; // variable to store the wealth of the current customer
        for (int j = 0; j < accounts[i].size(); ++j) { // iterate through each bank account of the current customer
            wealth += accounts[i][j]; // accumulate the money in each bank account to calculate the total wealth
        }
        maxWealth = std::max(maxWealth, wealth); // update the maximum wealth if necessary
    }
    return maxWealth; // return the maximum wealth
}
```

---
