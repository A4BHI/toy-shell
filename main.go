package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" and "os" imports in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var _ = os.Stdout

var commands = map[string]string{
	"echo": "echo",
	"type": "type",
	"exit": "exit",
}

func getCmd(cleanedcmd string) string {
	var cmd string
	for _, v := range cleanedcmd {

		if string(v) == " " {
			break
		}

		cmd += string(v)

	}
	return cmd
}

func searchCmd(path string, cmd string) string {
	var pathsearch string
	for _, v := range path {

		pathsearch += string(v)

		if string(v) == ":" || string(v) == "" {

			newpath := strings.Replace(pathsearch, ":", "", 1)

			entries, err := os.ReadDir(newpath)
			if err != nil {
				fmt.Println("Error in ReadinDIr", err)
				return ""
			}
			for _, entry := range entries {
				if !entry.IsDir() {
					if entry.Name() == cmd {
						return cmd + " is " + newpath + "/" + cmd
					}
				}
			}
			pathsearch = ""
		}

	}
	return cmd + ": not found"
}

func main() {
	// TODO: Uncomment the code below to pass the first stage
	path := os.Getenv("PATH")
	// fmt.Println(path)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error occured")
		}
		cleanedcmd := command[:len(command)-1]
		if command[:len(command)-1] == "exit" {
			return
		}
		cmd := getCmd(cleanedcmd)

		if cmd == commands["echo"] {
			fmt.Print(cleanedcmd[len(cmd)+1:] + "\n")
		} else if cmd == commands["type"] {
			switch cleanedcmd[len(cmd)+1:] {
			case "echo":
				fmt.Println(cleanedcmd[len(cmd)+1:] + " is a shell builtin")
			case "exit":
				fmt.Println(cleanedcmd[len(cmd)+1:] + " is a shell builtin")
			case "type":
				fmt.Println(cleanedcmd[len(cmd)+1:] + " is a shell builtin")
			default:
				// fmt.Println(cleanedcmd[len(cmd)+1:] + ": not found")
				searchResult := searchCmd(path, cleanedcmd[len(cmd)+1:])
				fmt.Println(searchResult)
			}

		} else {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}

	}

}
