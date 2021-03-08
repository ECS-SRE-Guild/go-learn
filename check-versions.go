package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func validateos() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
		os.Exit(1)
	}
}

func execute(cmdstr string) string {

	validateos()

	var cmdargs = strings.Split(cmdstr, " ")      // string arrayified
	var cmd = cmdargs[0]                          // command
	cmdargs = append(cmdargs[:0], cmdargs[1:]...) // argument array sans cmd
	out, err := exec.Command(cmd, cmdargs...).CombinedOutput()
	var output string

	if err != nil {
		output = err.Error()
	} else {
		output = string(out[:])
	}
	return output
}

func main() {
	gitver := execute("git --version")
	fmt.Println(gitver)
}
