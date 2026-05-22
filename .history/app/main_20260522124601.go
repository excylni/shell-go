package main

import (
	"fmt"
	"bufio"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)
	command := reader.ReadString('\n')

	if command != nil {
		fmt.Println(command[:len(command[-1])])
	}
	fmt.Print("$ ")
}
