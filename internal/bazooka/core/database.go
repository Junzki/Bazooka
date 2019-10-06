package core

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var dbConn *gorm.DB


func GetDbConn() (*gorm.DB, error) {
	if nil != dbConn {
		return dbConn, nil
	}

	dialect, connStr, err := GetConfig().Database.GetConnString()
	if nil != err {
		return nil, err
	}

	conn, err := gorm.Open(dialect, connStr)
	if nil != err {
		return nil, err
	}

	// Check if database is ok to connect.
	err = conn.DB().Ping()
	if nil != err {
		return nil, err
	}

	dbConn = conn
	return dbConn, err
}
