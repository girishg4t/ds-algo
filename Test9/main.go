package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type values struct {
	numberOfInputs int
	inputs         []string
}
type inputs struct {
	Items []values
}

var argsSeq int
var data inputs
var index int
var result int

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	if err != nil {
		fmt.Println("plese give valid input")
	}
	numberOfInputs, err := strconv.Atoi(string(input))

	if err != nil {
		fmt.Println("Not a valid inputs:", numberOfInputs)
	}
	takeInput(numberOfInputs)
	run(len(data.Items))
}

func run(length int) {
	calculate(data.Items[index].inputs, 0)
	fmt.Println(result)
	index = index + 1
	if index != length {
		result = 0
		run(length)
	}
}

func calculate(Items []string, seq int) {
	length := len(Items)
	validInt, err := strconv.Atoi(string(Items[seq]))
	if err != nil {
		panic("Not a valid Test case number")
	}
	if validInt > 0 {
		result = result + validInt*validInt
	}
	seq = seq + 1
	if seq != length {
		calculate(Items, seq)
	}
}

func takeInput(noOfInput int) {
	noOfInput = noOfInput - 1
	reader := bufio.NewReader(os.Stdin)
	noofTests, err := reader.ReadString('\n')
	noofTests = strings.Replace(noofTests, "\n", "", -1)
	if err != nil {
		panic("plese give valid input")
	}
	numberOfTestcase, err := strconv.Atoi(string(noofTests))
	if err != nil {
		panic("Not a valid Test case number")
	}
	tests, err := reader.ReadString('\n')
	if err != nil {
		panic("Not a valid Tests")
	}
	tests = strings.Replace(tests, "\n", "", -1)
	arguments := strings.Fields(tests)
	if len(arguments) != numberOfTestcase {
		panic("Number of test and tests does not match")
	}
	input := values{
		numberOfInputs: numberOfTestcase,
		inputs:         arguments,
	}
	data.Items = append(data.Items, input)
	if noOfInput == 0 {
		return
	}
	takeInput(noOfInput)
}
