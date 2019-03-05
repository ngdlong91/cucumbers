package storage

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"gopkg.in/doug-martin/goqu.v5"
)

type Storage interface {
	Builder() *goqu.Database
	DatabaseFromConf(path string) *goqu.Database
	Close()
}

type Orm interface {
	Orm() *gorm.DB
	OrmFromConf(path string) *gorm.DB
	Close()
}

var db *sql.DB
var path string

func SetPath(confPath string) {
	path = confPath
}

func init() {

}

type DataFetchMode int

const (
	FromStorage DataFetchMode = 1
	FromCache   DataFetchMode = 2
)
