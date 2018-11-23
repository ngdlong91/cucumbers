package storage

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"gopkg.in/doug-martin/goqu.v5"
)

type LiteOrm struct {
	db     *gorm.DB
	logger *logrus.Entry
}

func (orm *LiteOrm) Close() {
	orm.db.Close()
}

func (orm *LiteOrm) Orm() *gorm.DB {
	return orm.db
}

type LiteStorage struct {
	db     *goqu.Database
	logger *logrus.Entry
}

func (storage *LiteStorage) Builder() *goqu.Database {
	return storage.db
}

func (storage *LiteStorage) Close() {

}

func NewLiteOrm(logger *logrus.Entry) *LiteOrm {
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	db := OrmFromConf(dir)
	orm := &LiteOrm{
		db:     db,
		logger: logger.WithField("database", "postgres"),
	}
	return orm
}

func NewLiteStorage(logger *logrus.Entry) *LiteStorage {
	logger.Debug("Setup new postgres storage connection")
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	db := DatabaseFromConf(dir)
	orm := &LiteStorage{
		db:     db,
		logger: logger.WithField("database", "postgres"),
	}
	return orm
}
