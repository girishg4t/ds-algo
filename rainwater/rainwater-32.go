package rainwater

func Trap(height []int) int {
	count := 0
	left := height[0]
	right := height[len(height)-1]
	i := 0
	j := len(height) - 1
	for i < j {
		if left <= right {
			i++
			if left < height[i] {
				left = height[i]
			} else {
				count = count + (left - height[i])
			}
		} else {
			j--
			if right < height[j] {
				right = height[j]
			} else {
				count = count + (right - height[j])
			}
		}
	}
	return count
}
