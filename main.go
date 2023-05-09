package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var file = ".env"
var separator = "="

// Called before main, sets up command-line flags for the environment file and separator.
func init() {
	flag.StringVar(&file, "f", file, "environment file to load")
	flag.StringVar(&separator, "s", separator, "key value separator")
}

// Parse command-line flags, load the env and run the command.
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("usage: dotenv [[-f <name>][-s <separator>]] <command> [<args>...]")
		os.Exit(1)
	}
	env(file)
	run(args)
}

// Read the env file and set environment variables.
func env(name string) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(fmt.Sprintf("error reading file: %s", name))
		os.Exit(1)
	}
	variables := strings.Split(string(data), "\n")
	env := make(map[string]string)
	for _, v := range variables {
		v := strings.SplitN(v, separator, 2)
		if len(v) != 2 {
			continue
		}
		env[v[0]] = strings.TrimSpace(v[1])
	}
	for k, v := range env {
		err := os.Setenv(k, v)
		if err != nil {
			panic(err)
		}
	}
}

// Run the command using the current process input, output and error streams.
func run(args []string) {
	if len(args) == 0 {
		fmt.Println("no command provided")
		os.Exit(1)
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(fmt.Sprintf("command failed: %v\n", err))
		os.Exit(1)
	}
}
