package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		command = 
		command, err := reader.ReadString('\n')
		if err != nil{
			fmt.Print("Error during input:", err)
		}

		if command != "" {
		fmt.Println(command + ": command not found")	
		}
	
    	}	
	}