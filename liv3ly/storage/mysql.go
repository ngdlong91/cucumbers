package storage

import (
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"os"
	"time"

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
	dbServer := "test"
	if !IsTestMode {
		dbServer = viper.GetString("server")
	}

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

type MysqlStorageHelper struct {
}

func (h *MysqlStorageHelper) EncryptField(field, key string) goqu.SqlFunctionExpression {
	return goqu.Func("AES_ENCRYPT", field, key)
}

func (h *MysqlStorageHelper) DecryptField(field, key string) goqu.SqlFunctionExpression {
	return goqu.Func("AES_DECRYPT", goqu.I(field), key)
}

func NewMySQLHelper() *MysqlStorageHelper {
	return &MysqlStorageHelper{}
}

// CUSTOM NULL Handling structures

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

//// Scan implements the Scanner interface for NullInt64
//func (ni *NullInt64) Scan(value interface{}) error {
//	var i sql.NullInt64
//	if err := i.Scan(value); err != nil {
//		return err
//	}
//
//	// if nil then make Valid false
//	if reflect.TypeOf(value) == nil {
//		*ni = NullInt64{i.Int64, false}
//	} else {
//		*ni = NullInt64{i.Int64, true}
//	}
//	return nil
//}

// NullBool is an alias for sql.NullBool data type
type NullBool struct{ sql.NullBool }

//// Scan implements the Scanner interface for NullBool
//func (nb *NullBool) Scan(value interface{}) error {
//	var b sql.NullBool
//	if err := b.Scan(value); err != nil {
//		return err
//	}
//
//	// if nil then make Valid false
//	if reflect.TypeOf(value) == nil {
//		*nb = NullBool{b.Bool, false}
//	} else {
//		*nb = NullBool{b.Bool, true}
//	}
//
//	return nil
//}

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 struct{ sql.NullFloat64 }

//// Scan implements the Scanner interface for NullFloat64
//func (nf *NullFloat64) Scan(value interface{}) error {
//	var f sql.NullFloat64
//	if err := f.Scan(value); err != nil {
//		return err
//	}
//
//	// if nil then make Valid false
//	if reflect.TypeOf(value) == nil {
//		*nf = NullFloat64{f.Float64, false}
//	} else {
//		*nf = NullFloat64{f.Float64, true}
//	}
//
//	return nil
//}

// NullString is an alias for sql.NullString data type
type NullString struct{ sql.NullString }

//// Scan implements the Scanner interface for NullString
//func (ns *NullString) Scan(value interface{}) error {
//	var s sql.NullString
//	if err := s.Scan(value); err != nil {
//		return err
//	}
//
//	// if nil then make Valid false
//	if reflect.TypeOf(value) == nil {
//		*ns = NullString{s.String, false}
//	} else {
//		*ns = NullString{s.String, true}
//	}
//
//	return nil
//}

// NullTime is an alias for mysql.NullTime data type
type NullTime struct{ mysql.NullTime }

//// Scan implements the Scanner interface for NullTime
//func (nt *NullTime) Scan(value interface{}) error {
//	var t mysql.NullTime
//	if err := t.Scan(value); err != nil {
//		return err
//	}
//
//	// if nil then make Valid false
//	if reflect.TypeOf(value) == nil {
//		*nt = NullTime{t.Time, false}
//	} else {
//		*nt = NullTime{t.Time, true}
//	}
//
//	return nil
//}

// MarshalJSON for NullInt64
func (ni *NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ni.Int64)
}

// UnmarshalJSON for NullInt64
func (ni *NullInt64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ni.Int64)
	ni.Valid = (err == nil)
	return err
}

// MarshalJSON for NullBool
func (nb *NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nb.Bool)
}

// UnmarshalJSON for NullBool
func (nb *NullBool) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nb.Bool)
	nb.Valid = (err == nil)
	return err
}

// MarshalJSON for NullFloat64
func (nf *NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nf.Float64)
}

// UnmarshalJSON for NullFloat64
func (nf *NullFloat64) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &nf.Float64)
	nf.Valid = (err == nil)
	return err
}

// MarshalJSON for NullString
func (ns *NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

// UnmarshalJSON for NullString
func (ns *NullString) UnmarshalJSON(b []byte) error {
	err := json.Unmarshal(b, &ns.String)
	ns.Valid = (err == nil)
	return err
}

// MarshalJSON for NullTime
func (nt *NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	val := fmt.Sprintf("\"%s\"", nt.Time.Format(time.RFC3339))
	return []byte(val), nil
}

// UnmarshalJSON for NullTime
func (nt *NullTime) UnmarshalJSON(b []byte) error {
	s := string(b)
	// s = Stripchars(s, "\"")

	x, err := time.Parse(time.RFC3339, s)
	if err != nil {
		nt.Valid = false
		return err
	}

	nt.Time = x
	nt.Valid = true
	return nil
}
