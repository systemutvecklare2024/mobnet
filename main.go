package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	RunCommand("gh", "auth", "login")
	RunCommand("gh", "api", "-H", "Accept: application/vnd.github+json", "-H", "X-GitHub-Api-Version: 2022-11-28", "/users/Wizti")
}

func RunCommand(path string, args ...string) {
	command := exec.Command(path, args...)

	stdout, _ := command.StdoutPipe()
	command.Stderr = command.Stdout

	errStart := command.Start()
	if errStart != nil {
		panic(errStart)
	}

	output := ""

	stdoutscanner := bufio.NewScanner(stdout)
	lineEnded := true
	stdoutscanner.Split(bufio.ScanBytes)

	for stdoutscanner.Scan() {
		character := stdoutscanner.Text()
		if character == "\n" {
			lineEnded = true
		} else {
			if lineEnded {
				lineEnded = false
				fmt.Print(" ")
			}
		}
		fmt.Print(character)
		output += character
	}

	errWait := command.Wait()
	if errWait != nil {
		//fmt.Print(output)
		return
	}

	//fmt.Print(output)
}
