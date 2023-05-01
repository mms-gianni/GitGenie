package genie

import (
	"fmt"
	"os/exec"
)

func Diff() {
	out, err := exec.Command("git", "diff", "-u").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

func Status() {
	out, err := exec.Command("git", "status", "-s", "-uno").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
