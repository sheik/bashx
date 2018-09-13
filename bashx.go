package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func eval(line pipedCommands) error {
	if len(line[0]) == 0 {
		return nil
	}
	if f, ok := builtins[line[0][0]]; ok {
		f(line)
	} else {
		execProgram(line)
	}
	return nil
}

func execProgram(cmds pipedCommands) error {
	var commands []*exec.Cmd
	var command *exec.Cmd
	var err error
	var wg sync.WaitGroup
	var nextStdin io.Reader

	firstCommand := true
	for _, cmd := range cmds {
		if len(cmd) == 0 {
			continue
		}
		command = exec.Command(cmd[0], cmd[1:]...)
		if firstCommand {
			command.Stdin = os.Stdin
			firstCommand = false
		} else {
			command.Stdin = nextStdin
		}
		nextStdin, err = command.StdoutPipe()
		if err != nil {
			panic(err)
		}
		commands = append(commands, command)
	}
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	for _, command := range commands {
		wg.Add(1)
		go func(command *exec.Cmd) {
			defer wg.Done()
			command.Run()
		}(command)
	}
	wg.Wait()

	return nil
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	line := ""
	for {
		r, _ := reader.ReadString('\n')
		line = strings.TrimSuffix(r, "\n")
		if line == "" {
			return "", nil
		}
		if !(line[len(line)-1] == '\\') {
			break
		}
	}
	return line, nil
}

func tokenize(line string) (pipedCommands, error) {
	var commands [][]string
	for _, command := range strings.Split(line, "|") {
		commands = append(commands, strings.Fields(command))
	}
	return commands, nil
}

type pipedCommands [][]string

func main() {
	for {
		p := prompt{ps1: "bashx$ "}
		fmt.Print(p.format())

		line, err := readInput()

		if err != nil {
			panic(err)
		}

		tokens, err := tokenize(line)

		if err != nil {
			panic(err)
		}

		eval(tokens)
	}
}
