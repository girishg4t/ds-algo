package main

import "fmt"

func main() {
	fibo(10)
}

var count = 0
var firstPoint = 0
var secountPoint = 1

func fibo(i int) {
	if i > count {
		newPoint := firstPoint + secountPoint
		fmt.Println(newPoint)
		firstPoint = secountPoint
		secountPoint = newPoint
		i--
		fibo(i)
	}
}

// [abcd abdc acbd acdb adcb adbc bacd badc bcad bcda bdca bdac cbad cbda cabd cadb cdab cdba
// dbca dbac dcba dcab dacb dabc]
