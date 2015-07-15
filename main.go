package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	execute_command()
}

func execute_command() {
	command := os.Args[1]
	args := os.Args[2:]
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	start_time := time.Now()

	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus := exitError.Sys().(syscall.WaitStatus)
			os.Exit(waitStatus.ExitStatus())
		}
		return
	}

	total_time := int64(time.Since(start_time) / time.Millisecond)
	fmt.Printf("\nran `%v` in %vms\n", strings.Join(os.Args[1:], " "), total_time)
}
