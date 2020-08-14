package mysql

import (
	"fmt"
	"time"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/gonrails/gonrails/utils/common"
)

// Open open a mysql database to use
func Open(host, port, username, password, name string) *gorm.DB {
	timezone := "'Asia/Shanghai'"
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local&time_zone=%s",
		username,
		password,
		host,
		port,
		name,
		url.QueryEscape(timezone),
	)

	db, err := gorm.Open("mysql", dbConfig)
	common.PanicError(err)
	return db
}

// ConfigureMySQL configure db's idles connections, max connections
func ConfigureMySQL(db *gorm.DB, idles, connections int) {
	db.DB().SetMaxIdleConns(idles)
	db.DB().SetMaxOpenConns(connections)
	db.DB().SetConnMaxLifetime(time.Minute * 1)
}
