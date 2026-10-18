package main

func main() {

}

// 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func onlyOnce(nums []int) int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}
	for num, count := range countMap {
		if count == 1 {
			return num
		}
	}
	return 0
}

// 判断回文数
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}
	return original == reversed
}

// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
func isValid(s string) bool {
	stack := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 最长公共字串
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, str := range strs[1:] {
		for len(prefix) > 0 && !startsWith(str, prefix) {
			prefix = prefix[:len(prefix)-1]
		}
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func startsWith(s, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if s[i] != prefix[i] {
			return false
		}
	}
	return true
}

func longestCommonPrefix2(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	sumStr := ""

loop:
	for length := 0; ; length++ {
		c := ""

		for i := 0; i < len(strs); i++ {
			if len(strs[i]) < length+1 {
				break loop
			}

			if i == 0 {
				c = string(strs[i][length])
			}

			if i > 0 && c != string(strs[i][length]) {
				break loop
			}
		}

		sumStr += c
	}

	return sumStr
}

// 加一
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	newDigits := make([]int, n+1)
	newDigits[0] = 1
	return newDigits
}

func plusOne2(digits []int) []int {
	sum := 0

	for i := 0; i < len(digits); i++ {
		sum = sum*10 + digits[i]
	}

	sum += 1

	arr := []int{}

	for {
		dig := sum % 10

		array := []int{dig}

		arr = append(array, arr...)
		sum = int(sum / 10)

		if sum == 0 {
			return arr
		}
	}
}

// 给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
func removeDuplicates(nums []int) int {
	i := 0

	for j := 1; j < len(nums); j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}

	return len(nums[:i+1])
}

// 合并区间
func merge(intervals [][]int) [][]int {
	mergeSlice := [][]int{}

	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}

	for l := 0; l < len(intervals); l++ {
		isMerge := false
		for m := len(intervals) - 1; m > l; m-- {
			if intervals[l][1] >= intervals[m][0] {
				mergeSlice = append(mergeSlice, []int{intervals[l][0], intervals[m][1]})
				l = m
				isMerge = true
				break
			}
		}

		if !isMerge {
			mergeSlice = append(mergeSlice, intervals[l])
		}
	}

	return mergeSlice
}

// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
