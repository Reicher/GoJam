package main

import (
	"log"
	"fmt"
	"strconv"
)

func damage(program []rune) int{
	dmg := 0
	charge := 1
	for i := 0; i < len(program); i++ {
		if program[i] == 'S'{
			dmg += charge
		} else if program[i] == 'C' {
			charge *= 2
		} else {
			log.Fatal("Unknown program format")
		}
	}

	return dmg
}

func solver(shield int, program []rune) int{
	moves := 0
	for damage(program) > shield {
		improved := false

		// Hacking seasion started!
		for i := len(program)-1; i > 0; i-- {
			if program[i] == 'S'&& program[i-1] != 'S'{
				tmp := program[i]
				program[i] = program[i-1]
				program[i-1] = tmp
				moves++
				improved = true
 			}
		}

		if improved == false {
			return -1
		}
	}

	return moves
}

func main() {
	T := 0
	fmt.Scanf("%d", &T)
	for i := 0; i < T; i++ {
		shield := 0
		program := ""
		fmt.Scanf("%d %s", &shield, &program)
		moves := solver(shield, []rune(program))
		solution := "IMPOSSIBLE"
		if moves > -1 {
			solution = strconv.Itoa(moves)
		}

		line := "Case #" + strconv.Itoa(i+1) + ": " + solution
		fmt.Println(line)
	}
}
