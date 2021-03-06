package orm

import (
  "log"

  "github.com/tsoonjin/raspy/internal/orm/migration"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "github.com/tsoonjin/raspy/pkg/utils"
)

var autoMigrate, logMode, seedDB bool
var dsn, dialect string

// ORM struct to holds the gorm pointer to db
type ORM struct {
    DB *gorm.DB
}

func init() {
    dialect = utils.MustGet("GORM_DIALECT")
    dsn = utils.MustGet("GORM_CONNECTION_DSN")
    seedDB = utils.MustGetBool("GORM_SEED_DB")
    logMode = utils.MustGetBool("GORM_LOGMODE")
    autoMigrate = utils.MustGetBool("GORM_AUTOMIGRATE")
}

// Factory creates a db connection with the selected dialect and connection string
func Factory() (*ORM, error) {
    db, err := gorm.Open(dialect, dsn)
    if err != nil {
        log.Panic("[ORM] err: ", err)
    }
    orm := &ORM{
        DB: db,
    }
    // Log every SQL command on dev, @prod: this should be disabled?
    db.LogMode(logMode)
    // Automigrate tables
    if autoMigrate {
        err = migration.ServiceAutoMigration(orm.DB)
    }
    log.Println("[ORM] Database connection initialized.")
    return orm, err
}
