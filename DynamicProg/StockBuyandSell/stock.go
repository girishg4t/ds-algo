package stockbuyandsell

//  [7,1,5,3,6,4]
func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	sum := 0
	var max = prices[len(prices)-1]
	var maxPos = len(prices) - 1
	var min = prices[0]
	var minPos = 0
	if max > min && maxPos > minPos {
		return max - min
	}
	for i := 1; i < len(prices)-1; i++ {
		if prices[i] == 0 {
			continue
		}
		if prices[i] > max {
			max = prices[i]
			maxPos = i
		}
		if prices[i] < min && i != (len(prices)-1) {
			min = prices[i]
			minPos = i
		}

	}
	if max > min && maxPos > minPos {
		sum = sum + (max - min)
		min = len(prices)
		max = 0
		maxPos = 0
		minPos = 0
	}
	return sum
}
