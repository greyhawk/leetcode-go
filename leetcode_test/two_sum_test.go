package leetcode_test

import "testing"
import "github.com/greyhawk/leetcode-go/leetcode"

func TestTwoSum(t *testing.T) {
	nums := []int{2,3,4,5,6,7}
	target := 9
	indexs := leetcode.TwoSum(nums, target)
	if indexs[0] != 2 {
		t.Errorf("index 0 expected 2 but %d", indexs[0])
	}
	if indexs[1] != 3 {
		t.Errorf("index 1 expected 3 but %d", indexs[1])
	}
}
