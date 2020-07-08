package main

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"

	// "github.com/HoloPollock/qitGo/helpers"
	"github.com/HoloPollock/qitGo/view"

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

	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()
	i := view.NewTextBox()
	i.SetRect(0, 0, 50, 10)
	ui.Render(i)
	uiEvents := ui.PollEvents()
	readKey := false
	for {
		e := <-uiEvents
		if e.Type == ui.KeyboardEvent {
			switch e.ID {
			case "<Left>":
				i.CursorLeft()
				readKey = true
			case "<Right>":
				i.CursorRight()
				readKey = true
			case "<Enter>":
				return
			case "<Backspace>":
				i.DeleteLetter()
				readKey = true
			case "<Mousedown>":
				fmt.Println(e)
			}

			if !readKey {
				if len(e.ID) == 1 {
					i.AddLetter(e.ID)
				}

			} else {
				readKey = false
			}
		} else if e.Type == ui.MouseEvent {

		}

		ui.Render(i)
	}
	// cl := view.NewClickList()
	// cl.Title = "test"
	// cl.Rows = []view.Row{
	// 	{
	// 		Name:     "view/view.go",
	// 		Selected: true,
	// 		Status:   helpers.Changed,
	// 	},
	// 	{
	// 		Name:     "main.go",
	// 		Selected: false,
	// 		Status:   helpers.Added,
	// 	},
	// }
	// cl.SetRect(0, 0, 50, 10)
	// ui.Render(cl)

	// previousKey := ""
	// uiEvents := ui.PollEvents()
	// for {
	// 	e := <-uiEvents
	// 	switch e.ID {
	// 	case "q", "<C-c>":
	// 		return
	// 	case "j", "<Down>":
	// 		cl.ScrollDown()
	// 	case "k", "<Up>":
	// 		cl.ScrollUp()
	// 	case "<C-d>":
	// 		cl.ScrollHalfPageDown()
	// 	case "<C-u>":
	// 		cl.ScrollHalfPageUp()
	// 	case "<C-f>":
	// 		cl.ScrollPageDown()
	// 	case "<C-b>":
	// 		cl.ScrollPageUp()
	// 	case "g":
	// 		if previousKey == "g" {
	// 			cl.ScrollTop()
	// 		}
	// 	case "<Home>":
	// 		cl.ScrollTop()
	// 	case "G", "<End>":
	// 		cl.ScrollBottom()
	// 	case "<Enter>":
	// 		cl.Toggle()
	// 	}

	// 	if previousKey == "g" {
	// 		previousKey = ""
	// 	} else {
	// 		previousKey = e.ID
	// 	}

	// 	ui.Render(cl)
	// }

}
