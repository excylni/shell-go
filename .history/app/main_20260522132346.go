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
		command, err := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if err != nil{
			fmt.Print("Error during input:", err)
		}

		if command != "exit" {
		fmt.Println(command + ": command not found")	
		}
		
		if command == "exit" {
			break
		}

		if strings.HasPrefix(command, "echo")
    	}	
	}