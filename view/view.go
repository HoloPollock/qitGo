package view

import (
	"github.com/gizak/termui/v3/widgets"
)

type Header struct {
	Repo   *widgets.Paragraph
	Branch *widgets.Paragraph
}

func NewHeader(repoName, branchName string) *Header {
	r := widgets.NewParagraph()
	r.Text = repoName
	b := widgets.NewParagraph()
	b.Text = branchName
	return &Header{
		Repo:   r,
		Branch: b,
	}

}

type View struct {
	Header *Header
	List   *widgets.List
}

// func newView() *View {

// }
