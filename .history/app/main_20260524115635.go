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

		switch {
			
			case command == "":
				continue

			case command == "exit":
				os.Exit(0)
			
			case strings.HasPrefix(command, "echo "):
				fmt.Println(command[5:])

			case strings.HasPrefix(command, "type "):
				target := command[5:]

				if target == "exit" || target == "echo" || target == "type" {
					fmt.Println(target + " is a shell builtin")
				}
			
			default:
				fmt.Println(command + ": command not found")
		}
    	}	
	}