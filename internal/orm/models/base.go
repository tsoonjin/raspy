package models

import (
  "time"

  // "github.com/jinzhu/gorm"
)

// BaseModel defines the common columns that all db structs should hold, usually
// db structs based on this have no soft delete
type BaseModel struct {
  // ID should use uuid_generate_v4() for the pk's
  ID        int        `gorm:"primary_key;AUTO_INCREMENT"`
  CreatedAt time.Time  `gorm:"index;not null;default:CURRENT_TIMESTAMP"` // (My|Postgre)SQL
  UpdatedAt *time.Time `gorm:"index"`
}

// BaseModelSoftDelete defines the common columns that all db structs should
// hold, usually. This struct also defines the fields for GORM triggers to
// detect the entity should soft delete
type BaseModelSoftDelete struct {
  BaseModel
  DeletedAt *time.Time `sql:"index"`
}
