package main

func repeatedNTimes(nums []int) int {
	mMap := make(map[int]int)

	for _, count := range nums {
		mMap[count] = mMap[count] + 1
	}

	maxV := -1
	maxI := -1

	for k, v := range mMap {
		if v > maxV {
			maxV = v
			maxI = k
		}
	}

	return maxI
}
