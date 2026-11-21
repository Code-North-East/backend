package arraysandhashing

import (
	"errors"
	"fmt"
	"testing"
	"strconv"
	"strings"
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

func TestProductExceptSelf(t *testing.T) {
	res := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Printf("left suffix product: %v\n", res)
}

func TestEncode(t *testing.T) {
	s := &Solution{}
	er := "5#Hello7#Pratyay"
	rs := s.Encode([]string{"Hello", "Pratyay"})
	if rs != er {
		t.Errorf("Failed: Expected result %s; found result %s", er, rs)
	}
}

func TestDecode(t *testing.T) {
	s := &Solution{}
	i := "6#Hello#7##Ankita"
	rs := s.Decode(i)
	t.Errorf("%v", rs)
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

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	products := make([]int, len(nums))
	// find the left product of all the elements
	leftProduct := 1
	i := 0
	for i < len(nums) {
		products[i] = leftProduct
		leftProduct *= nums[i]
		i++
	}

	// output till now for input 1, 2, 3, 4
	// 1, 1, 2, 6

	// find the right product of all the left product elements
	i = len(nums)-1
	rightProduct := 1
	for i >= 0 {
		products[i] *= rightProduct
		rightProduct *= nums[i]
		i--
	}

	return products
}

type Solution struct{}

// Encode joins a list of strings into a single encoded string.
func (s *Solution) Encode(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    
    var sb strings.Builder
    for _, str := range strs {
        // Use byte length and strings.Builder for O(n) performance
        sb.WriteString(strconv.Itoa(len(str)))
        sb.WriteString("#")
        sb.WriteString(str)
    }
    return sb.String()
}

// Decode extracts the original list of strings from the encoded string.
func (s *Solution) Decode(encoded string) []string {
    if encoded == "" {
        return []string{}
    }

    var res []string
    i := 0
    n := len(encoded)

    for i < n {
        // Find the position of the next '#' delimiter
        j := i
        for j < n && encoded[j] != '#' {
            j++
        }

        // If we reach the end without finding '#', something is wrong
        if j >= n {
            break
        }

        // Parse the length string located between i and j
        lengthStr := encoded[i:j]
        length, err := strconv.Atoi(lengthStr)
        if err != nil {
            break
        }

        // The actual string starts right after '#' (at j + 1)
        start := j + 1
        end := start + length

        // Extract the string and append to results
        res = append(res, encoded[start:end])

        // Move pointer `i` past the current decoded string
        i = end
    }

    return res
}