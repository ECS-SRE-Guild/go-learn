package main

import (
	"fmt"
	"os/exec"
	"strings"
)


func execute(cmdstr string) (string, error) {
	cmdargs := strings.Split(cmdstr, " ")         // string arrayified
	cmd := cmdargs[0]                             // command
	cmdargs = append(cmdargs[:0], cmdargs[1:]...) // argument array sans cmd
	out, err := exec.Command(cmd, cmdargs...).CombinedOutput()
	return string(out[:]), err
}

func main() {
	gitver, _ := execute("git --version")
	verno := strings.Split(gitver, " ")[2]
	fmt.Println(verno)

	gcnf, _ := execute("git config -l")

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

	glog, _ := execute("git log --oneline")
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
