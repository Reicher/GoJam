package main

import (
	"os"
	"fmt"
	"log"
	"github.com/reicher/GoJam/iohandler"
)

func solver(c iohandler.RawCase) iohandler.RawSolution{
	sol := iohandler.RawSolution{Id: c.Id, Text: "TEST"}
	return sol
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("No inputfile given.")
	}

	cases := iohandler.GetRawCases(os.Args[1])
	fmt.Println(len(cases))

	var solutions []iohandler.RawSolution
	for _, c := range cases {
		s := solver(c)
		solutions = append(solutions, s)
	}

	fmt.Println(len(solutions))
}
