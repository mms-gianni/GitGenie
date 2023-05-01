package genie

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

func Status() {

	r, err := git.PlainOpen(".")

	if err != nil {
		fmt.Println(err)
	}

	w, err := r.Worktree()

	if err != nil {
		fmt.Println(err)
	}

	status, err := w.Status()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(status)
}
