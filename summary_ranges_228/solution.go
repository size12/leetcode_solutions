package summary_ranges_228

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/summary-ranges/.

func summaryRanges(nums []int) []string {
	result := make([]string, 0)

	if len(nums) == 0 {
		return result
	}

	first := nums[0]

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] > 1 {
			last := nums[i]
			if first == last {
				result = append(result, strconv.Itoa(first))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", first, last))
			}
			first = nums[i+1]
		}
	}

	last := nums[len(nums)-1]
	if first == last {
		result = append(result, strconv.Itoa(first))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", first, last))
	}

	return result
}
