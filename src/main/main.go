package main

import (
	"command"
	"fmt"
)

func main() {
	cmd := command.ParseCmd()

	if cmd.VersionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.HelpFlag || cmd.Class == "" {
		command.PrintUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *command.Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.CpOption, cmd.Class, cmd.Args)
}
