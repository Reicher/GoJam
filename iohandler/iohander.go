package iohandler

import (
	"fmt"
	"os"
	"strconv"
	"bufio"
)

type RawCase struct{
	Id int
	Row []string
}

type RawSolution struct{
	Id int
	Row string
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// readLines reads a whole file into memory
// and returns a slice of its lines as strings.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func GetRawCases(path string) []RawCase{
	fmt.Println("Loading file:", path)

	file_data, err := readLines(path)
	check(err)

	cases_n, err := strconv.Atoi(file_data[0])
	check(err)

	rows_n := len(file_data) - 1
	rows_per_case := rows_n / cases_n

	var cases []RawCase
	for i := 0; i < cases_n; i++ {
		start := 1 + i * rows_per_case
		end   := start + rows_per_case

		rows := file_data[start : end]

		test_case := RawCase{Id: i, Row: rows}

		cases = append(cases, test_case)
	}
	return cases
}

func SaveSolutions(solutions []RawSolution, filename string){
	fmt.Println("Saving file:", filename)

	file, err := os.Create(filename)
 	check(err)

	defer file.Close()

	for index, solution := range solutions {

		if index != solution.Id {
			panic("SaveSolutions: Id missmatch!" )
		}

		line := "Case #" + strconv.Itoa(solution.Id) + ": " + solution.Row + "\r"
		file.WriteString(line)
	}
}
