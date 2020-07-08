package view

import (
	"image"

	"github.com/HoloPollock/qitGo/helpers"
	. "github.com/gizak/termui/v3"
	rw "github.com/mattn/go-runewidth"
)

type Row struct {
	Name     string
	Selected bool
	Status   helpers.Status
}

func (r *Row) toggle() {
	r.Selected = !r.Selected
}

type ClickList struct {
	Block
	Rows             []Row
	TextStyle        Style
	SelectedRow      int
	topRow           int
	SelectedRowStyle Style
}

func NewClickList() *ClickList {
	return &ClickList{
		Block:            *NewBlock(),
		TextStyle:        Theme.List.Text,
		SelectedRowStyle: NewStyle(ColorWhite, ColorBlack),
	}
}

func (self *ClickList) Draw(buf *Buffer) {
	self.Block.Draw(buf)

	point := self.Inner.Min

	if self.SelectedRow >= self.Inner.Dy()+self.topRow {
		self.topRow = self.SelectedRow - self.Inner.Dy() + 1
	} else if self.SelectedRow < self.topRow {
		self.topRow = self.SelectedRow
	}

	for row := self.topRow; row < len(self.Rows) && point.Y < self.Inner.Max.Y; row++ {
		cells := ParseStyles(self.Rows[row].Name, self.TextStyle)
		defstyle := Theme.List.Text
		buf.SetCell(NewCell('[', defstyle), point)
		point = point.Add(image.Pt(rw.RuneWidth('['), 0))
		if self.Rows[row].Selected {
			buf.SetCell(NewCell('✓', defstyle), point)
			point = point.Add(image.Pt(rw.RuneWidth('✓'), 0))
		} else {
			buf.SetCell(NewCell(' ', defstyle), point)
			point = point.Add(image.Pt(rw.RuneWidth(' '), 0))
		}
		buf.SetCell(NewCell(']', defstyle), point)
		point = point.Add(image.Pt(rw.RuneWidth(']'), 0))
		point = point.Add(image.Pt(3, 0)) // add spacing
		for j := 0; j < len(cells) && point.Y < self.Inner.Max.Y; j++ {
			style := cells[j].Style
			if row == self.SelectedRow {
				style = self.SelectedRowStyle
			}
			if cells[j].Rune == '\n' {
				point = image.Pt(self.Inner.Min.X, point.Y+1)
			} else {
				if point.X+1 == self.Inner.Max.X+1 && len(cells) > self.Inner.Dx() {
					buf.SetCell(NewCell(ELLIPSES, style), point.Add(image.Pt(-1, 0)))
					break
				} else {
					buf.SetCell(NewCell(cells[j].Rune, style), point)
					point = point.Add(image.Pt(rw.RuneWidth(cells[j].Rune), 0))
				}
			}
		}
		point = point.Add(image.Pt(2, 0))

		switch self.Rows[row].Status {
		case helpers.Added:
			buf.SetCell(NewCell('+', NewStyle(ColorGreen)), point)
			point = point.Add(image.Pt(rw.RuneWidth('+'), 0))
		case helpers.Deleted:
			buf.SetCell(NewCell('-', NewStyle(ColorRed)), point)
			point = point.Add(image.Pt(rw.RuneWidth('-'), 0))
		case helpers.Changed:
			buf.SetCell(NewCell('·', NewStyle(ColorYellow)), point)
			point = point.Add(image.Pt(rw.RuneWidth('·'), 0))
		}
		point = image.Pt(self.Inner.Min.X, point.Y+1)
	}
}

func (self *ClickList) ScrollAmount(amount int) {
	if len(self.Rows)-int(self.SelectedRow) <= amount {
		self.SelectedRow = len(self.Rows) - 1
	} else if int(self.SelectedRow)+amount < 0 {
		self.SelectedRow = 0
	} else {
		self.SelectedRow += amount
	}
}

func (self *ClickList) ScrollUp() {
	self.ScrollAmount(-1)
}

func (self *ClickList) ScrollDown() {
	self.ScrollAmount(1)
}

func (self *ClickList) ScrollPageUp() {
	// If an item is selected below top row, then go to the top row.
	if self.SelectedRow > self.topRow {
		self.SelectedRow = self.topRow
	} else {
		self.ScrollAmount(-self.Inner.Dy())
	}
}

func (self *ClickList) ScrollPageDown() {
	self.ScrollAmount(self.Inner.Dy())
}

func (self *ClickList) ScrollHalfPageUp() {
	self.ScrollAmount(-int(FloorFloat64(float64(self.Inner.Dy()) / 2)))
}

func (self *ClickList) ScrollHalfPageDown() {
	self.ScrollAmount(int(FloorFloat64(float64(self.Inner.Dy()) / 2)))
}

func (self *ClickList) ScrollTop() {
	self.SelectedRow = 0
}

func (self *ClickList) ScrollBottom() {
	self.SelectedRow = len(self.Rows) - 1
}

func (self *ClickList) Toggle() {
	self.Rows[self.SelectedRow].toggle()
}
