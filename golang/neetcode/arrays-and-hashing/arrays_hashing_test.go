package arraysandhashing

import (
	"errors"
	"fmt"
	"testing"
)

func TestHasDuplicate(t *testing.T) {
	if exists := hasDuplicate([]int{1, 2, 3, 4, 5}); exists {
		t.Errorf("Expected no duplicates, but found duplicates")
	}
} 

func TestTwoSum(t *testing.T) {
	// descriptive test case, covering all the edge scenarios
	i := []int{3, 4, 5, 6}
	target := 7
	if r := twoSum(i, target); r == nil {
		t.Errorf("Should return the indexes, as the target can be met")
	} else {
		fmt.Printf("index: %v\n", r)
	}
}

func TestIsAnagram(t *testing.T) {
	res := isAnagram("abcab", "bbaac")
	if res != true {
		t.Errorf("Expected true, a valid anagram; but returned false")
	}
}

func TestUniqueSum(t *testing.T) {
	if sum, err := uniqueSum([]int{3, 3, 3}); err == nil {
		t.Errorf("Expected error, but got sum: %d", sum)
	}
}

func TestMaximumSubarraySum(t *testing.T) {
	maxSum := maximumSubarraySum([]int{1, 2, 3, 3, 3}, 3)
	if maxSum != 6 {
		t.Errorf("Expected maximum subarray sum to be 6, but got %d", maxSum)
	}
}

func TestLongestConsecutiveSequence(t *testing.T) {
	expectedRes := 4
	res := longestConsecutiveSequence([]int{2,20,4,10,3,4,5})
	if res != expectedRes {
		t.Errorf("the longest sequence should have been %d, but it is %d", expectedRes, res)
	}
}

func hasDuplicate(nums []int) bool {
    data := make(map[int]int)
	for i, num := range nums {
		if _, ok := data[num]; ok {
			return true
		}
		data[num] = i
	}
	return false
}

func maximumSubarraySum(nums []int, k int) int64 {
	maxSum := 0
	for i := range nums {
		if i+k-1 < len(nums) {
			sum, err :=uniqueSum(nums[i:i+k])
			if sum > maxSum && err == nil {
				maxSum = sum
			}
		}
	}
	return int64(maxSum)
}

func uniqueSum(nums []int) (int, error) {
	sum := 0
	data := make(map[int]int)
	for i, num := range nums {
		if _, exists := data[num]; !exists {
			data[num] = i
			sum += num
		} else {
			return -1, errors.New("Duplicate element found in the array")
		}
	}
	return sum, nil
}

// todo: a clean solution is there in place, need to fix accordingly.
func isAnagram(s, t string) bool {
	// Approach: use a map that will store the character occurrence frequency
	// Iterate through the first string and increase the frequency and then iterate through the second string and 
	// decrease the frequency.
	// if the map has all the frequency count equal to zero, this is awesome and you can return true

	// first condition check should be if the length of both the strings are equal or not
	if len(s) != len(t) {
		return false
	}
	// declare the frequency map
	f := make(map[rune]int)
	// iterate through the first string
 	for _, c := range s {
		if _, exists := f[c]; exists {
			// increment by 1
			f[c] = f[c]+1
		} else {
			// assign the initial value
			f[c] = 1
		}
	}
	// iterating through the second string
	for _, c := range t {
		if _, exists := f[c]; exists {
			f[c] = f[c]-1 
		} else {
			return false
		}
	}
	// iterating over the map
	for _, v := range f {
		// if the decrement has not boiled down to 0
		if v != 0 {
			return false
		}
	}
	// if every character in that map has gone down to zero, then it should be returning true
	return true
}

func twoSum(nums []int, target int) []int {
	// apporach is straight forward, iterate over the array
	// add the index and the number associated with it, compare if the difference exists, if it does simply return the indexes 
	s := make(map[int]int)
	for i, n := range nums {
		// check if the difference exists
		if _, exists := s[target-n]; exists {
			return []int{s[target-n], i}
		} else {
			// storing each number index
			s[n] = i
		}
	}
	return nil
}

func longestConsecutiveSequence(nums []int) int {
	if len(nums) == 0 {
		// no consecutive sequence
		return -1
	}
	// create a hash set, empty struct cause this will have no memory.
	data := make(map[int]struct{})
	for _, num := range nums {
		data[num] = struct{}{}
	}
	longestStreak := 0
	for key := range data {
		// check if the start is valid; if a value less than a particular key exists, 
		// then it is a false start
		if _, exists := data[key-1]; !exists {
			// assign the key to a variable
			currentNum := key
			currentStreak := 1
			// Upward check; use a infinite loop
			for {
				if _, exists := data[currentNum+1]; exists {
					currentNum++
					currentStreak++
				} else {
					break
				}
			}
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			} 
		}
	}
	return longestStreak
}