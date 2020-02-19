package main

import (
	"flag"
	"log"
	"sort"

	"github.com/ilyakaznacheev/pretty-nice-tasks/tools/io/pizza"
)

func main() {
	var (
		inputName, outputName string
		printLog              bool
	)
	flag.StringVar(&inputName, "i", "input.txt", "input file path")
	flag.StringVar(&outputName, "o", "output.txt", "output file path")
	flag.BoolVar(&printLog, "l", false, "print log")
	flag.Parse()

	input, err := pizza.Input(inputName)
	if err != nil {
		log.Fatal(err)
	}

	output := findTheBest(input)

	if printLog {
		logResults(*input, output)
	}

	err = pizza.Output(outputName, output)
	if err != nil {
		log.Fatal(err)
	}
}

func findTheBest(input *pizza.InputData) pizza.OutputData {
	type pizzaType struct {
		idx, slices int
	}

	pizzas := make([]pizzaType, 0, len(input.SliceNumbers))
	for idx, s := range input.SliceNumbers {
		pizzas = append(pizzas, pizzaType{idx, s})
	}

	sort.Slice(pizzas, func(i, j int) bool {
		return pizzas[i].slices > pizzas[j].slices
	})

	res := pizza.OutputData{
		Types:    0,
		Ordering: []int{},
	}

	total := input.Slices

	for _, p := range pizzas {
		if p.slices < total {
			total -= p.slices
			res.Types++
			res.Ordering = append(res.Ordering, p.idx)
		}
	}

	sort.Slice(res.Ordering, func(i, j int) bool {
		return res.Ordering[i] < res.Ordering[j]
	})

	return res
}

func logResults(i pizza.InputData, o pizza.OutputData) {
	pizzas := make(map[int]struct{}, o.Types)
	for _, p := range o.Ordering {
		pizzas[p] = struct{}{}
	}
	sum := 0
	for idx, s := range i.SliceNumbers {
		if _, ok := pizzas[idx]; ok {
			sum += s
		}
	}

	log.Printf("slices want: %d, solution: %d", i.Slices, sum)
}
