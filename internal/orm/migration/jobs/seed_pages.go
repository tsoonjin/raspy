package jobs

import (
    "github.com/tsoonjin/raspy/internal/orm/models"
    "github.com/jinzhu/gorm"
    "gopkg.in/gormigrate.v1"
)

var (
    src = "http://www.astroawani.com/berita-malaysia/bayi-sembilan-bulan-antara-kes-terbaharu-covid-19-247265"
    firstPage *models.Page = &models.Page{
        Src: src,
    }
)

// SeedUsers inserts the first pages
var SeedPages *gormigrate.Migration = &gormigrate.Migration{
    ID: "SEED_PAGES",
    Migrate: func(db *gorm.DB) error {
        return db.Create(&firstPage).Error
    },
    Rollback: func(db *gorm.DB) error {
        return db.Delete(&firstPage).Error
    },
}
