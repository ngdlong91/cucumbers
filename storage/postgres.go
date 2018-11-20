package storage

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"gopkg.in/doug-martin/goqu.v5"
)

type PostgresOrm struct {
	db     *gorm.DB
	logger *logrus.Entry
}

func (orm *PostgresOrm) Close() {
	orm.db.Close()
}

func (orm *PostgresOrm) Orm() *gorm.DB {
	return orm.db
}

type PostgresStorage struct {
	db     *goqu.Database
	logger *logrus.Entry
}

func (storage *PostgresStorage) Builder() *goqu.Database {
	return storage.db
}

func (storage *PostgresStorage) Close() {

}

func NewPostgresOrm(logger *logrus.Entry) *PostgresOrm {
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	db := OrmFromConf(dir)
	orm := &PostgresOrm{
		db:     db,
		logger: logger.WithField("database", "postgres"),
	}
	return orm
}

func NewPostgresStorage(logger *logrus.Entry) *PostgresStorage {
	logger.Debug("Setup new postgres storage connection")
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	db := DatabaseFromConf(dir)
	orm := &PostgresStorage{
		db:     db,
		logger: logger.WithField("database", "postgres"),
	}
	return orm
}
