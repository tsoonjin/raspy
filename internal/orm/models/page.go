package models

type Page struct {
  BaseModelSoftDelete
  Src                 string  `gorm:"not null;unique_index:idx_src"`
  Title               *string
  Text                *string
  Summary             *string
  Videos              *string
  Images              *string
  Links               *string
}
