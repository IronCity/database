package mysql

import (
    "log"
    "github.com/xormplus/xorm"
    "strings"
)

type MySqlDB struct {
    engine *xorm.Engine
    url string
}

var db MySqlDB

func GetEngine(url string) (*xorm.Engine) {
    if db.engine == nil {
        db.url = url
        CreateEngine(url)
    } else {
        if err := db.engine.Ping(); err != nil {
            log.Println("connect is close")
            db.engine.Close()
        }
    }
    return db.engine
}

func GetEngineDefault() (*xorm.Engine) {
    if db.engine == nil {
        CreateEngineDefault()
    } else {
        if err := db.engine.Ping(); err != nil {
            log.Println("connect is close")
            db.engine.Close()
        }
    }
    return db.engine
}

func CreateEngine(url string)  {
    var err error
    if db.engine, err = xorm.NewEngine("mysql", url); err != nil{
        log.Println(err)
    }

    db.engine.ShowSQL(true)

    err = db.engine.SetSqlTemplateRootDir("../template").InitSqlTemplate(xorm.SqlTemplateOptions{Extension: ".stpl"})
    if err != nil {
        log.Println(err)
    }
}

func CreateEngineDefault() {
    var err error
    if db.engine, err = xorm.NewEngine("mysql", db.url); err != nil{
        log.Println(err)
    }
}

func ChangeUrl(url string)  {
    db.url = url
}

