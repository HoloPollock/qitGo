package view

import (
	"image"

	. "github.com/gizak/termui/v3"
)

type TextBox struct {
	Block
	Buffer    string
	length    int
	TextStyle Style
	cursor    int
}

func NewTextBox() *TextBox {
	return &TextBox{
		Block:     *NewBlock(),
		TextStyle: Theme.Paragraph.Text,
		Buffer:    "",
		length:    0,
		cursor:    0,
	}
}

func (self *TextBox) Draw(buf *Buffer) {
	self.Block.Draw(buf)
	cells := ParseStyles(self.Buffer, self.TextStyle)

	for i := 0; i < len(cells); i++ {
		buf.SetCell(cells[i], image.Pt(i, 0).Add(self.Inner.Min))
	}
}

func (self *TextBox) AddLetter(l string) {
	self.length++
	self.Buffer = self.Buffer[:self.cursor] + l + self.Buffer[self.cursor:]
	self.cursor++
}

func (self *TextBox) CursorRight() {
	self.cursor++
}

func (self *TextBox) CursorLeft() {
	self.cursor--
}
