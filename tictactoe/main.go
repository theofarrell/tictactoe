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
	fmt.Printf("%+v\n", grid)
}