package main
import "fmt"

func main(){
	const (
	  GridSize = 5
	)

	grid := make([]string, GridSize * GridSize)
	for idx:=0; idx<len(grid); idx++{
		grid[idx] = fmt.Sprintf("%d", idx)
	}

	fmt.Println("Hello.  I'm Theo, the computer. Can you beat me at Tic-Tac-Toe?")

	// this code prints the starting grid
	for idx:=0; idx<len(grid); idx++{
		fmt.Print(grid[idx])
		fmt.Print(" | ")
		if (idx + 1) % GridSize == 0 {
			fmt.Println("")
		}
	}
}