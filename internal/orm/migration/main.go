package migration

import (
    "fmt"
    "log"

    "github.com/tsoonjin/raspy/internal/orm/migration/jobs"
    "github.com/tsoonjin/raspy/internal/orm/models"
    "github.com/jinzhu/gorm"
    "gopkg.in/gormigrate.v1"
)

func updateMigration(db *gorm.DB) error {
    return db.AutoMigrate(
        &models.Page{},
    ).Error
}

// ServiceAutoMigration migrates all the tables and modifications to the connected source
func ServiceAutoMigration(db *gorm.DB) error {
    // Keep a list of migrations here
    options := &gormigrate.Options{TableName: "Pages"}
    m := gormigrate.New(db, options, nil)
    m.InitSchema(func(db *gorm.DB) error {
        log.Println("[Migration.InitSchema] Initializing database schema")
        if err := updateMigration(db); err != nil {
            return fmt.Errorf("[Migration.InitSchema]: %v", err)
        }
        // Add more jobs, etc here
        return nil
    })

    if err := updateMigration(db); err != nil {
        return err
    }
    m = gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
        jobs.SeedPages,
    })
    return m.Migrate()
}
