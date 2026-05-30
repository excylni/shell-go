package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"os/exec"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")
		command, err := reader.ReadString('\n')

		if err != nil{
			fmt.Print("Error during input:", err)
			break
		}

		command = strings.TrimSpace(command)
		args := strings.Fields(command)

		switch {
			
			case command == "":
				continue
			

			case command == "pwd":
				// returns path of the current directory
				dir, err := os.Getwd()
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error retrieving current directory", err)
				} else {
					fmt.Println(dir)
				}
			case command == "exit":
				os.Exit(0)
			
			case strings.HasPrefix(command, "echo "):
				fmt.Println(command[5:])

			case strings.HasPrefix(command, "type "):
				// checking the type 
				target := strings.TrimSpace(command[5:])

				if target == "exit" || target == "echo" || target == "type" || target == "pwd" {
					fmt.Println(target + " is a shell builtin")

				} else if path, err := exec.LookPath(target) ;err == nil {
					fmt.Println(target + " is " + path)

				} else {
					fmt.Println(target + ": not found")
				}
				
			
			default:
				// If not builtin or type, check if executable and run it
				if path, err := exec.LookPath(args[0]); err == nil {
					cmd := exec.Command(path, args[1:]...)
					
					cmd.Args = args
					cmd.Stdin = os.Stdin
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr

					err := cmd.Run()
					if err != nil {
						fmt.Print("Error running command", err)
					} 
					
				} else {
						fmt.Println(command + ": command not found")
										
			}
    	}	
		}
	}	