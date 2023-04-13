//
// Created by TaiTitans on 4/14/2023.
// Running Sum of 1d Array

#include <iostream>
#include <vector>
using namespace std;
vector<int> runningSum(vector<int> &nums){
    int sum = 0;
    for(int i=0;i< nums.size();i++){
        sum += nums[i];
        nums[i] = sum;
    }
    return nums;
}

int main(){
    int n;
    cout << "Enter the size of the array: ";
    cin >> n;
    vector<int> nums;
    for(int i = 0;i<n;i++){
        nums.push_back(i+1);
    }
    vector<int> result = runningSum(nums);
    cout << "Running Sum of the Array: ";
    for(int j=0;j<result.size();j++){
        cout << result[j] << " ";
    }
    return 0;
}