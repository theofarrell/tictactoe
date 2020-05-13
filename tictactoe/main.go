package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	// This is the width/height size of my grid
	GridSize = 3
)

func writeLine(length int){
	for i:=0; i<length; i++{
		fmt.Print("-")
	}
	fmt.Println()
}

func printGrid(grid []string){
	linelen := 6 * GridSize
	writeLine(linelen)
	for idx:=0; idx<len(grid); idx++{
		if idx % GridSize == 0 {
			fmt.Printf("|")
		}
		fmt.Printf("%3s", grid[idx])
		fmt.Print(" | ")
		if (idx + 1) % GridSize == 0 {
			fmt.Println("")
			writeLine(linelen)
		}
	}
	fmt.Println()
}

func exitError(err error){
	if err !=nil {
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
}

func inviteComputerTurn(grid []string){
	//todo
}

func inviteUserTurn(grid []string){
	fmt.Println("Ok it's your turn. Please enter a number that is not taken.  You are Xs.")
	fmt.Println("")
	scanner := bufio.NewScanner(os.Stdin)
	var err error
	var sel int64
	for scanner.Scan() {
		sel, err = strconv.ParseInt(scanner.Text(), 10, 64)
		if err == nil {
			fmt.Printf("You have chosen the number: %d\n", sel)
			// todo check that the square has not already been taken
			break
		}
		fmt.Println("You have not entered a number!  Please try again.")
	}
	exitError(scanner.Err())
	grid[sel] = "X"
	printGrid(grid)
}


func main(){
	// making room in memory for our grid
	grid := make([]string, GridSize * GridSize)

	//
	for idx:=0; idx<len(grid); idx++{
		grid[idx] = fmt.Sprintf("%d", idx)
	}

	fmt.Println("Hello.  I'm Theo, the computer. Can you beat me at Tic-Tac-Toe?")
	fmt.Println()

	// this code prints the starting grid
	printGrid(grid)

	// todo loop...
	inviteUserTurn(grid)
}