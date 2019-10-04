package core

import "github.com/jinzhu/gorm"

var dbConn *gorm.DB


//func GetDbConn() *gorm.DB {
//	if nil != dbConn {
//		return dbConn
//	}
//
//
//}
//
//func DbInit(db *gorm.DB) error {
//	db, err := gorm.Open()
//	return nil
//}
