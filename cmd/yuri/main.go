package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	args := os.Args[1:]
	fmt.Printf("running %s\n", args[0])

	cmd := exec.Command(strings.Split(args[0], " ")[0], strings.Split(args[0], " ")[1:]...)

	stdout, _ := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	println()
	cmd.Wait()
}
