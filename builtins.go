package main

import "os"

var builtins = map[string]func(interface{}){
	"exit":  exit,
	"async": async,
}

func exit(line interface{}) {
	os.Exit(0)
}

func async(line interface{}) {
	prog := [][]string{line.([]string)[1:]}
	go execProgram(prog)
}
