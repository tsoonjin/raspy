package main

import (
    "log"

    "github.com/tsoonjin/raspy/pkg/server"
    "github.com/tsoonjin/raspy/internal/orm"

)

func main() {
    orm, err := orm.Factory()
    if err != nil {
        log.Panic(err)
    }

    server.Run(orm)
}

