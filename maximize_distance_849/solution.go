package maximize_distance_849

// https://leetcode.com/problems/maximize-distance-to-closest-person/.

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDistToClosest(seats []int) int {
	maxSpace := 0
	startIndex := -1

	leftPersonIndex := -1
	rightPersonIndex := -1

	for index, seat := range seats {
		if seat == 1 {
			maxSpace = max(maxSpace, index-startIndex)
			startIndex = index
			rightPersonIndex = index
			if leftPersonIndex == -1 {
				leftPersonIndex = index
			}
		}
	}

	maxDist := max(maxSpace, len(seats)-startIndex) / 2
	maxDistWithLeft := max(maxDist, leftPersonIndex)
	maxDistWithLeftAndRight := max(maxDistWithLeft, len(seats)-rightPersonIndex-1)

	return maxDistWithLeftAndRight

}
