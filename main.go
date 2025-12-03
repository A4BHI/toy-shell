package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

func searchCmd(cmd string) {

	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s: not found\n", cmd)
		return
	}
	fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, path)
}

func main() {
	// TODO: Uncomment the code below to pass the first stage
	// path := os.Getenv("PATH")
	// fmt.Println(path)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error occured")
		}
		cleanedcmd := strings.TrimSpace(command)
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
				searchCmd(strings.TrimSpace(cleanedcmd[len(cmd)+1:]))
			}

		} else {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}

	}

}

