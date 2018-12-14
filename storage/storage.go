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

func init() {

}

//func DatabaseFromConf(path string) *goqu.Database {
//	viper.SetConfigName("db") // name of config file (without extension)
//
//	viper.AddConfigPath(path + "/conf") // optionally look for config in the working directory
//	if err := viper.ReadInConfig(); err != nil {
//		panic(fmt.Errorf("Fatal errors config file: %s \n", err))
//	}
//	conn := viper.GetString("conn")
//
//	dbServer := viper.GetString("server")
//	user := viper.GetString(dbServer + ".user")
//	pass := viper.GetString(dbServer + ".password")
//	host := viper.GetString(dbServer + ".host")
//	port := viper.GetString(dbServer + ".port")
//	dbName := viper.GetString(dbServer + ".db")
//
//	pgDb, err := sql.Open("postgres", fmt.Sprintf(
//		conn, user, pass, dbName, host, port))
//	if err != nil {
//		panic(err.Error())
//	}
//	db = goqu.New("postgres", pgDb)
//	return db
//
//}

type DataFetchMode int

const (
	FromStorage DataFetchMode = 1
	FromCache   DataFetchMode = 2
)
