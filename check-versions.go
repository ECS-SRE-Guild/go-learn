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

func execute(cmdstr string) {

	validateos()

	var cmdargs = strings.Split(cmdstr, " ")      // string arrayified
	var cmd = cmdargs[0]                          // command
	cmdargs = append(cmdargs[:0], cmdargs[1:]...) // argument array sans cmd
	out, err := exec.Command(cmd, cmdargs...).CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	fmt.Println(output)
}

func main() {
	execute("/usr/sbin/lsof -iTCP -sTCP:LISTEN")
	execute("gcc --version")
	execute("java -version")
	execute("python3 --version")
	execute("php --version")

}
