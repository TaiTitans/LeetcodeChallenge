//
// Created by TaiTitans on 4/15/2023.
//
#include <vector>

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
