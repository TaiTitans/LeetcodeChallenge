class Solution{
	public: int maxProduct(vector<int>& nums){
		int n = nums.size();
		if(n == 0){
			return 0;
		}
		int result = nums[0];
		int maxE = nums[0];
		int minE = nums [0];
		
		for(int i = 0; i<n;i++){
			if(nums[i]<0){
				swap(maxE,minE);
			}
			maxE = max(nums[i],maxE * nums[i]);
			minE = min(nums[i], minE * nums[i]);
			
			result = max(result, maxE);
		}
		return result
	}
};