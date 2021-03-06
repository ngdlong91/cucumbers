package storage

import (
	"os"

	"fmt"

	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/doug-martin/goqu.v5"
	"time"
)

var storageInfo PostgresStorageInfo

var storageMode int

func init() {
	storageMode = 0
}

type PostgresStorageInfo struct {
	IsUpdated bool
	Conn      string
	User      string
	Pass      string
	Host      string
	Port      string
	DbName    string
}

func SetupPostgresInfo(conn, user, pass, host, port, dbName string) {
	storageInfo.Conn = conn
	storageInfo.User = user
	storageInfo.Pass = pass
	storageInfo.Host = host
	storageInfo.Port = port
	storageInfo.DbName = dbName
	storageInfo.IsUpdated = true

}

type PostgresOrm struct {
	mode   int
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
	mode   int
	db     *goqu.Database
	logger *logrus.Entry
}

func (storage *PostgresStorage) DatabaseFromConf(path string) *goqu.Database {
	viper.SetConfigName("db")           // name of config file (without extension)
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
	storage.logger.Debug("Setup with postgres info ", fmt.Sprintf(
		conn, user, pass, dbName, host, port))
	if db == nil {
		db, err = sql.Open("postgres", fmt.Sprintf(
			conn, user, pass, dbName, host, port))
		if err != nil {
			panic(err.Error())
		}
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(time.Hour)

	return goqu.New("postgres", db)
}

func SetMode(mode int) {
	storageMode = mode
}

func (storage *PostgresStorage) DatabaseFromInfo() *goqu.Database {
	if !storageInfo.IsUpdated {
		panic("please update database info")
	}

	var err error
	storage.logger.Debug("Setup with postgres info ", fmt.Sprintf(
		storageInfo.Conn, storageInfo.User, storageInfo.Pass, storageInfo.DbName, storageInfo.Host, storageInfo.Port))
	if db == nil {
		db, err = sql.Open("postgres", fmt.Sprintf(
			storageInfo.Conn, storageInfo.User, storageInfo.Pass, storageInfo.DbName, storageInfo.Host, storageInfo.Port))
		if err != nil {
			panic(err.Error())
		}
		db.SetMaxOpenConns(5)
		db.SetMaxIdleConns(2)
		db.SetConnMaxLifetime(time.Hour)
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

	if path == "" {
		dir, err := os.Getwd()
		if err != nil {
			panic(err.Error())
		}
		path = dir
	}

	orm := &PostgresStorage{
		logger: logger.WithField("database", "postgres"),
	}

	if storageMode == 0 {
		orm.db = orm.DatabaseFromConf(path)
	} else {
		orm.db = orm.DatabaseFromInfo()
	}

	return orm
}
