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

	gcnf := execute("git config -l")

	config := strings.Split(gcnf, "\n")
	for i, v := range config {
		entry := strings.Split(v, "=")
		if len(entry) < 2 {
			continue
		}
		key := entry[0]
		val := entry[1]
		fmt.Printf("%d: %s = '%s' \n", i, key, val)
	}

	glog := execute("git log --oneline")
	logarray := strings.Split(glog, "\n")
	for _, v := range logarray {
		entryarray := strings.SplitN(v, " ", 2)
		if len(entryarray) < 2 {
			continue
		}
		commitid := entryarray[0]
		reason := entryarray[1]
		fmt.Println(commitid, reason)
	}
}
