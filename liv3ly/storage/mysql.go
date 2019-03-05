package storage

import (
	"os"

	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/doug-martin/goqu.v5"
	_ "gopkg.in/doug-martin/goqu.v5/adapters/mysql"
)

type MysqlStorage struct {
	db     *goqu.Database
	logger *logrus.Entry
}

func (storage *MysqlStorage) DatabaseFromConf(path string) *goqu.Database {
	viper.SetConfigName("db") // name of config file (without extension)

	viper.AddConfigPath(path + "/conf") // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal errors config file: %s \n", err))
	}
	conn := viper.GetString("conn")

	dbServer := viper.GetString("server")
	user := viper.GetString(dbServer + ".user")
	pass := viper.GetString(dbServer + ".password")
	host := viper.GetString(dbServer + ".host")
	port := viper.GetString(dbServer + ".port")
	dbName := viper.GetString(dbServer + ".db")
	var err error
	if db == nil {
		db, err = sql.Open("mysql", fmt.Sprintf(
			conn, user, pass, host, port, dbName))
		if err != nil {
			panic(err.Error())
		}
	}

	return goqu.New("mysql", db)
}

func (storage *MysqlStorage) Builder() *goqu.Database {
	return storage.db
}

func (storage *MysqlStorage) Close() {
	db.Close()
}

func NewMysqlStorage(logger *logrus.Entry) *MysqlStorage {
	logger.Debug("Setup new mysql storage connection")

	if path == "" {
		dir, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}

		path = dir
	}

	orm := &MysqlStorage{
		logger: logger.WithField("database", "mysql"),
	}

	orm.db = orm.DatabaseFromConf(path)
	return orm
}
