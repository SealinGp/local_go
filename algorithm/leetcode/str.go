package main

func main() {

}

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	tmp := 0
	for _, v := range nums {
		if v == 1 {
			tmp++
		}

		if v == 0 {
			if max < tmp {
				max = tmp
			}

			tmp = 0
		}
	}

	if max < tmp {
		max = tmp
	}

	return max
}

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := range nums {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

func twoSum(numbers []int, target int) []int {
	index1 := 0
	index2 := len(numbers) - 1
	for index1 < index2 {
		if numbers[index1]+numbers[index2] == target {
			break
		}

		if numbers[index1]+numbers[index2] < target {
			index1++
			continue
		}

		index2--
	}

	return []int{index1 + 1, index2 + 1}
}

func arrayPairSum(nums []int) int {
	numsLen := len(nums) - 1
	Sort(nums, 0, numsLen)

	i := 0
	sum := 0
	for i < numsLen {
		next := i + 1
		sum += min(nums[i], nums[next])
		i += 2
	}

	return sum
}

func min(a, b int) int {
	if b < a {
		a = b
	}
	return a
}

func Sort(arr []int, start, end int) {
	if start >= end {
		return
	}

	privotIndex := split(arr, start, end)
	Sort(arr, start, privotIndex-1)
	Sort(arr, privotIndex+1, end)
}

func split(arr []int, start, end int) int {
	mid := arr[start]
	border := start

	for next := start + 1; next <= end; next++ {
		if arr[next] < mid {
			border++
			arr[next], arr[border] = arr[border], arr[next]
		}
	}

	arr[start] = arr[border]
	arr[border] = mid
	return border
}
