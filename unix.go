//+build !windows

package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func GetGOPATHS() []string {
	gopath := os.Getenv("GOPATH")
	return strings.Split(gopath, ":")
}

func GetGitPath(gopath string) []string {
	cmd := exec.Command("find", gopath+"/src/github.com", "-name", ".git")

	finds, err := cmd.Output()
	if err != nil {
		fmt.Printf(CL_RED+"Error:%s"+CL_DEFAULT+".\r\n", err.Error())
		return make([]string, 0)
	}
	return strings.Split(string(finds), ".git")
}

func GitPull(path string) bool {
	c1 := exec.Command("cd", path)
	c2 := exec.Command("git", "pull")
	c2.Stdin, _ = c1.StdoutPipe()
	c1.Run()
	_, err := c2.Output()
	if err != nil {
		fmt.Printf("\r\n" + CL_RED + "[" + path + "] Error:" + "," + err.Error() + CL_DEFAULT + "\r\n")
		return false
	}
	return true
}
