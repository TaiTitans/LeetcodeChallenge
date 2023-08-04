var containsDuplicate = function(nums){
    const seen = Set();
    for(const num of nums){
        if(seen.has(num)){
            return true;
        }
        seen.add(num);
    }
    return false;
}