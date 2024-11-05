package database

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBOption func(db *gorm.DB) *gorm.DB

type DBConnectionProps struct {
	Hostname     string
	ReadHostname string
	Username     string
	Password     string
	DBName       string
}

func CreateDBConn(
	isDebugMode bool,
	driver string,
	coreConnection *DBConnectionProps,
) (*gorm.DB, error) {
	coreDsn := dsn(
		driver, coreConnection.Username, coreConnection.Password, coreConnection.Hostname, coreConnection.DBName,
	)

	gormOption := &gorm.Config{}
	if isDebugMode {
		gormOption = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	var db *gorm.DB
	var err error
	switch driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(coreDsn), gormOption)
		// err = db.Use(dbresolver.Register(dbresolver.Config{
		// 	Replicas:          []gorm.Dialector{mysql.Open(replicaDsn)},
		// 	TraceResolverMode: isDebugMode,
		// }))
	default:
		return nil, errors.New("invalid driver")
	}
	return db, err
}

func dsn(driver, username, password, hostname, dbName string) string {
	switch driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", username, password, hostname, dbName)
	case "postgres":
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s", hostname, username, password, dbName)
	}

	return ""
}
