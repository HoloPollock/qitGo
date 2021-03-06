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

var cur = Cell{
	Rune: ' ',
	Style: Style{
		Fg: ColorBlack,
		Bg: ColorWhite,
	},
}

func (self *TextBox) Draw(buf *Buffer) {
	self.Block.Draw(buf)
	cells := ParseStyles(self.Buffer, self.TextStyle)

	for i := 0; i < len(cells); i++ {
		if i == self.cursor {
			cells[i].Style.Bg = ColorWhite
			cells[i].Style.Fg = ColorBlack
		}
		buf.SetCell(cells[i], image.Pt(i, 0).Add(self.Inner.Min))
	}
	if self.cursor == self.length {
		buf.SetCell(cur, image.Pt(self.length, 0).Add(self.Inner.Min))

	}
}

func (self *TextBox) AddLetter(l string) {
	self.length++
	self.Buffer = self.Buffer[:self.cursor] + l + self.Buffer[self.cursor:]
	self.cursor++
}

func (self *TextBox) DeleteLetter() {
	if self.length == 0 {
		return
	}
	self.length--
	self.Buffer = self.Buffer[:self.cursor-1] + self.Buffer[self.cursor:]
	self.cursor--

}

func (self *TextBox) CursorRight() {
	self.cursor++
	if self.cursor > self.length {
		self.cursor = self.length
	}
}

func (self *TextBox) CursorLeft() {
	self.cursor--
	if self.cursor < 0 {
		self.cursor = 0
	}
}
