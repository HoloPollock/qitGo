package main

import (
	"fmt"
	"log"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	repo, err := git.PlainOpen("./")
	check(err)
	tree, err := repo.Worktree()
	sta, err := tree.Status()
	check(err)
	fmt.Printf("%+v", sta)
	it, err := repo.CommitObjects()
	err = it.ForEach(func(r *object.Commit) error {
		fmt.Println(r)
		return nil
	})
	check(err)

}
