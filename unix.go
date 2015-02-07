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
	return strings.Split(string(finds), ".git\n")
}

func GitPull(path string) bool {
	if *fix {
		exec.Command("git", "-C", path, "reset", "--hard").Run()
	}

	c := exec.Command("git", "-C", path, "pull")

	_, err := c.Output()
	if err != nil {
		fmt.Printf("\r\n" + CL_RED + "[" + path + "] Error:" + err.Error() + CL_DEFAULT + "\r\n")
		return false
	}
	return true
}
