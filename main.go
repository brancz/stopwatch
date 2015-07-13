package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	command := os.Args[1]
	args := os.Args[2:]
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	start_time := time.Now()

	err := cmd.Run()
	if err != nil {
		println(err.Error())
		return
	}

	total_time := int64(time.Since(start_time) / time.Millisecond)
	fmt.Printf("\nran `%v` in %vms\n", strings.Join(os.Args, " "), total_time)
}
