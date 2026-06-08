package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"os/exec"
	"unicode"
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
		var quoteChar rune =  0

		for i := 0; i < len(input); i++ {
			char := rune(input[i])

			switch char {
			case '\'', '"':
				// Toggling state of being inside/outside of quotes
				if quoteChar == 0 {
					quoteChar = char
				} else if quoteChar == char {
					quoteChar = 0
				// If already already inside a quote, write normally 
				} else {
					currentArg.WriteRune(char)
				}
												
			case '\\': // for \ Backlash character
				if quoteChar == '"' { // If in double quotes, escapes next character
					if i+1 < len(input) {
						nextChar := rune(input[i+1])
						if nextChar == '$' || nextChar == '"' || nextChar == '\\' {
							currentArg.WriteRune(nextChar)
							i++ // skip next character
						} else { // if its normal character like \a
							currentArg.WriteRune(char)
						}
					} else {
						fmt.Println("syntax error: unterminated double quote")
					}
				} else if quoteChar == '\'' {
					currentArg.WriteRune(char)

				} else { // outside of quotes it escapes next character
					if i+1 < len(input) {
						currentArg.WriteRune(rune(input[i+1]))
						i++ // skip next character 
					}
				}
			
			case '$': 
				if quoteChar == 0 || quoteChar == '"' {
					// looking ahead to see when the variable ends
					start := i + 1 
					end := start
					 
					for end < len(input) {
						// convert character to rune for unicode check
						r := rune(input[end])
						if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
							end++
						} else {
							break
						}
					}
					if end > start {
						varName := input[start:end]
						value := os.Getenv(varName)
						currentArg.WriteString(value)
						i = end - 1
					} else { // treat it normally if its alone
						currentArg.WriteRune(char)

					}
				} else { // single quote logic
					currentArg.WriteRune(char)
				}

			case ' ':
				if quoteChar != 0 { // space is prote~cted
					currentArg.WriteRune(char)
				} else {
					if currentArg.Len() > 0 {
						args = append(args, currentArg.String())
						currentArg.Reset()
						}
					}
			default:
				currentArg.WriteRune(char)

			}
		}
		
		if currentArg.Len() > 0 {
			args = append(args, currentArg.String())
			}
		return args
	}
		