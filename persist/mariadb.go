package persist

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/tminerio/tminer/tminer-profit/config"
)

// GMariadb GMariadb
var GMariadb Mariadb

// Mariadb Mariadb
type Mariadb struct {
	db *gorm.DB
}

// Init Init
func (maria *Mariadb) Init() error {
	db, err := gorm.Open(config.MariaDB.Dialect, config.MariaDB.URL)
	if err != nil {
		return err
	}

	db.LogMode(false)
	db.DB().SetMaxIdleConns(config.MariaDB.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MariaDB.MaxOpenConns)
	db.DB().SetConnMaxLifetime(10 * time.Minute)

	maria.db = db

	return nil
}
