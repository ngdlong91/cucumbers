package storage

import (
	"os"

	"fmt"

	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/doug-martin/goqu.v5"
)

type PostgresOrm struct {
	db     *gorm.DB
	logger *logrus.Entry
}

func (orm *PostgresOrm) OrmFromConf(path string) *gorm.DB {
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
//var err err
	db, err := gorm.Open("postgres", fmt.Sprintf(
		conn, user, pass, dbName, host, port))

	if err != nil {
		panic(err)
	}
	return db
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

func (storage *PostgresStorage) DatabaseFromConf(path string) *goqu.Database {
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

	db, err := sql.Open("postgres", fmt.Sprintf(
		conn, user, pass, dbName, host, port))
	if err != nil {
		panic(err.Error())
	}
	return goqu.New("postgres", db)
}

func (storage *PostgresStorage) Builder() *goqu.Database {
	return storage.db
}

func (storage *PostgresStorage) Close() {
	db.Close()
}

func NewPostgresOrm(logger *logrus.Entry) *PostgresOrm {
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	orm := &PostgresOrm{
		logger: logger.WithField("database", "postgres"),
	}
	orm.db = orm.OrmFromConf(dir)

	return orm
}

func NewPostgresStorage(logger *logrus.Entry) *PostgresStorage {
	logger.Debug("Setup new postgres storage connection")
	dir, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}

	orm := &PostgresStorage{
		logger: logger.WithField("database", "postgres"),
	}

	orm.db = orm.DatabaseFromConf(dir)
	return orm
}
