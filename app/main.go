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
		args := parseInput(command)

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
			
			case args[0] == "cd":
				// change directory 
				targetDir := args[1]
				if targetDir == "~" || strings.HasPrefix(targetDir, "~/") {
					home, err := os.UserHomeDir()

					if err!= nil {
						fmt.Fprintln(os.Stderr, "cd: cannot find home directory:", err)
						continue
					} else {
						os.Chdir(home)
					}

				} else {
					err := os.Chdir(targetDir)

					if err != nil {
						fmt.Fprintln(os.Stderr, "cd: " + targetDir + ": No such file or directory" )
					}
				}

			case command == "exit":
				os.Exit(0)
			
			case args[0] == "echo":
				fmt.Println(strings.Join(args[1:]," "))

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

	func parseInput(input string) []string{
		var args []string 
		var currentArg strings.Builder
		insideQuotes := false

		for _, char := range input {
			switch char {
			case '\'':
				// Toggling state of being inside/outside of quote
				insideQuotes =  !insideQuotes
			
			case ' ':
				if insideQuotes { // spaces are treated as normal characters
					currentArg.WriteRune(char)
				} else { 
					if currentArg.Len() > 0 {
						args = append(args, currentArg.String())
						currentArg.Reset() // clean buffer for next argument
					}
				}
			
			default:
				currentArg.WriteRune(char)
			}
		}
		// adding the rest if command didnt end on spacebar
		if currentArg.Len() > 0 {
			args = append(args, currentArg.String())
		}

		return args
	}