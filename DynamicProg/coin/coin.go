package coin

var coins = []int{1}
var numCoins = 0

func GetTotal(amount int, coinPos int) int {
	if amount >= coins[coinPos] {
		numCoins++
		amount = amount - coins[coinPos]
		return GetTotal(amount, coinPos)
	}
	coinPos++
	if amount == 0 {
		return numCoins
	}

	if coinPos >= len(coins) {
		return -1
	}
	return GetTotal(amount, coinPos)
}
