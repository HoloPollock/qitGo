package view

import (
	"image"

	. "github.com/gizak/termui/v3"
	rw "github.com/mattn/go-runewidth"
)

type Button struct {
	Block
	Label     string
	TextStyle Style
	OnClick   func()
	InFocus   bool
}

func NewButton() *Button {
	return &Button{
		Block:     *NewBlock(),
		TextStyle: Theme.Paragraph.Text,
		OnClick:   nil,
		InFocus:   false,
	}
}

func (self *Button) Draw(buf *Buffer) {
	self.Block.Draw(buf)
	if self.InFocus {
		self.TextStyle.Bg = ColorBlack
	} else {
		self.TextStyle.Bg = ColorClear
	}
	cells := ParseStyles(self.Label, self.TextStyle)
	point := self.Inner.Min
	buf.SetCell(NewCell('[', self.TextStyle), point)
	point = point.Add(image.Pt(rw.RuneWidth('['), 0))
	buf.SetCell(NewCell(' ', self.TextStyle), point)
	point = point.Add(image.Pt(1, 0))
	buf.SetCell(NewCell(' ', self.TextStyle), point)
	point = point.Add(image.Pt(1, 0))

	for _, cell := range cells {
		buf.SetCell(NewCell(cell.Rune, cell.Style), point)
		point = point.Add(image.Pt(rw.RuneWidth(cell.Rune), 0))

	}
	buf.SetCell(NewCell(' ', self.TextStyle), point)
	point = point.Add(image.Pt(1, 0))
	buf.SetCell(NewCell(' ', self.TextStyle), point)
	point = point.Add(image.Pt(1, 0))
	buf.SetCell(NewCell(']', self.TextStyle), point)
	point = point.Add(image.Pt(rw.RuneWidth(']'), 0))
}

func (self *Button) Toggle() {
	self.InFocus = !self.InFocus
}

func (self *Button) Call() {
	self.OnClick()
}
