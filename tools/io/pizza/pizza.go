/*
Package pizza

Problem description

You are organizing a Hash Code hub and want to order pizza for your hub’s
participants. Luckily, there is a nearby pizzeria with really good pizza.
The pizzeria has different types of pizza, and to keep the food ordering for your hub
interesting, you can only order at most one pizza of each type. Fortunately, there are
many types of pizza to choose from!
Each type of pizza has a specified size: the size is the number of slices in a pizza of this
type.
You estimated the maximum number of pizza slices that you want to order for your
hub based on the number of registered participants. In order to reduce food waste,
your goal is to order as many pizza slices as possible, but not more than the
maximum number.

To use the library you need to import it

	import "github.com/ilyakaznacheev/pretty-nice-tasks/tools/io/pizza"

	inData, err := pizza.Input("test_input.in")
	...
	err = pizza.Output("test_output.out", outData)
*/
package pizza

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// InputData is a pizza task input format:
//
// Each input data set is provided in a plain text le containing exclusively ASCII
// characters with lines terminated with a single '\n' character (UNIX-style line endings).
// When a single line contains multiple elements, they are separated by single spaces.
//
// The rst line of the data set contains the following data:
// - an integer M (1 ≤ M ≤ 10^9) – the maximum number of pizza slices to order
// - an integer N (1 ≤ N ≤ 10^5 ) – the number of different types of pizza
//
// The second line contains N integers – the number of slices in each type of pizza, in
// non-decreasing order:
// - 1 ≤ S0 ≤ S1 ≤ … ≤ SN-1 <= M
type InputData struct {
	Slices       int
	Types        int
	SliceNumbers []int
}

// OutputData s a pizza task output format:
//
// The output should contain two lines:
// - The rst line should contain a single integer K (0 ≤ K ≤ N) – the number of
// different types of pizza to order.
// - The second line should contain K numbers – the types of pizza to order (the
// types of pizza are numbered from 0 to N-1 in the order they are listed in the
// input).
//
// The total number of slices in the ordered pizzas must be less than or equal to M.
type OutputData struct {
	Types    int
	Ordering []int
}

// Input reads file and parses it's data into pizza task input format
func Input(fileName string) (*InputData, error) {
	rawData, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	data := strings.Split(string(rawData), "\n")
	if len(data) < 2 {
		return nil, fmt.Errorf("invalid data length: %d", len(data))
	}

	res := InputData{}

	ln1 := strings.Split(data[0], " ")
	if len(ln1) < 2 {
		return nil, fmt.Errorf("invalid first line length: %d", len(ln1))
	}
	res.Slices, err = strconv.Atoi(ln1[0])
	if err != nil {
		return nil, err
	}

	res.Types, err = strconv.Atoi(ln1[1])
	if err != nil {
		return nil, err
	}

	numbersStr := strings.Split(data[1], " ")
	res.SliceNumbers = make([]int, 0, len(numbersStr))
	for _, ns := range numbersStr {
		n, err := strconv.Atoi(ns)
		if err != nil {
			return nil, err
		}
		res.SliceNumbers = append(res.SliceNumbers, n)
	}

	return &res, nil
}

// Output converts the data to output format and writes to a file
func Output(fileName string, out OutputData) error {
	orderingStr := make([]string, 0, len(out.Ordering))
	for _, o := range out.Ordering {
		orderingStr = append(orderingStr, strconv.Itoa(o))
	}
	res := fmt.Sprintf("%d\n%s", out.Types, strings.Join(orderingStr, " "))
	return ioutil.WriteFile(fileName, []byte(res), 0644)
}
