package model

import (
	"pbf_demo/internal/model/datatype"
)

type Pyramid struct {
	ID uint          `gorm:"column:id" json:"id,omitempty"`
	X  uint          `gorm:"column:x" json:"x"`
	Y  uint          `gorm:"column:y" json:"y"`
	Z  uint          `gorm:"column:z" json:"z"`
	P  datatype.JSON `gorm:"column:pyramid"`
}
