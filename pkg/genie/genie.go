package genie

import (
	"fmt"
	"os"
	"os/exec"
)

func Diff() string {
	out, err := exec.Command("git", "diff", "-u").Output()
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(out))

	if len(out) == 0 {
		fmt.Println("No changes to commit")
		os.Exit(1)
	}

	return string(out)
}

func Status() {
	out, err := exec.Command("git", "status", "-s", "-uno").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}
