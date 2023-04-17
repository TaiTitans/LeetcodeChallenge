//
// Created by TaiTitans on 4/17/2023.
//
#include <vector>
using namespace std;
class Solution {
public:
    vector<bool> kidsWithCandies(vector<int>& candies, int extraCandies) {
        int max_candies = 0;
        for(int candy : candies){
            max_candies = max(max_candies, candy);
        }
        vector<bool> result;
        for(int candy : candies){
            if(candy+extraCandies >= max_candies){
                result.push_back(true);
            }else {
                result.push_back(false);
            }
        }
        return result;
    }
};