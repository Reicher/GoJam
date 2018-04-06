package main

import (
	"os"
	"log"
	"github.com/reicher/GoJam/iohandler"
)

func solver(c iohandler.RawCase) iohandler.RawSolution{
	sol := iohandler.RawSolution{Id: c.Id, Row: "1234567890"}
	return sol
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("No inputfile given.")
	}

	cases := iohandler.GetRawCases(os.Args[1])

	var solutions []iohandler.RawSolution
	for _, c := range cases {
		s := solver(c)
		solutions = append(solutions, s)
	}

	out_filename := os.Args[1][:len(os.Args[1])-3] + ".out"
	iohandler.SaveSolutions(solutions, out_filename)
}
