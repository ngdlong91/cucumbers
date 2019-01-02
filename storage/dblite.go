package storage

import (
	"os"

	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	_ "gopkg.in/doug-martin/goqu.v4/adapters/sqlite3"
	"gopkg.in/doug-martin/goqu.v5"

	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type LiteOrm struct {
	db     *gorm.DB
	logger *logrus.Entry
}

func (orm *LiteOrm) OrmFromConf(path string) *gorm.DB {
	panic("implement me")
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

func (storage *LiteStorage) DatabaseFromConf(path string) *goqu.Database {
	viper.SetConfigName("db") // name of config file (without extension)

	viper.AddConfigPath(path + "/conf") // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal errors config file: %s \n", err))
	}

	dbServer := viper.GetString("server")
	file := fmt.Sprintf("%s%s", "./", viper.GetString(dbServer+".file"))
	storage.logger.Debug("Database file to open ", file)
	sqliteDB, err := sql.Open("sqlite3", file)
	if err != nil {
		panic(err.Error())
	}
	return goqu.New("sqlite3", sqliteDB)
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

	orm := &LiteOrm{
		logger: logger.WithField("database", "sqlite3"),
	}

	orm.db = orm.OrmFromConf(dir)
	return orm
}

func NewLiteStorage(logger *logrus.Entry) *LiteStorage {
	logger.Debug("Setup new lite storage connection")
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	orm := &LiteStorage{
		logger: logger.WithField("database", "sqlite3"),
	}
	orm.db = orm.DatabaseFromConf(dir)
	return orm
}
