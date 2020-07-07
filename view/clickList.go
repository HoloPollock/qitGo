package view

import (
	"github.com/HoloPollock/qitGo/helpers"
	. "github.com/gizak/termui/v3"
)

type Row struct {
	Name     string
	Selected bool
	Status   helpers.Status
}

type ClickList struct {
	*Block
}
