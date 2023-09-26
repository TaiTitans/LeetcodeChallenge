#include<iostream>
#include<stdio.h>
#include<vector>
#include<string.h>

#define infinity -10000000
using namespace std;
class Solution{
public:
		int maxSubArray(vector<int>& nums){
			int result = infinity;
			int current_max = 0;
			for(int i = 0; i < nums.size();i++){
				current_max = max(nums[i],nums[i] + current_max);
				result = max(result, current_max);
			}
			return result;
		}
	
};
int main() {
	vector<int>nums={5,4,3,2,1};
	Solution solution;
	int max_sum_array = solution.maxSubArray(nums);
	cout<< max_sum_array << endl;
	return 0;
}