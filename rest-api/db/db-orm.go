package db

import (
	"os"
	"path/filepath"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SQLiteDB struct {
	DB *gorm.DB
}

var (
	instance *SQLiteDB
	once sync.Once
)

func GetInstance() *SQLiteDB {
	once.Do(func() {
		db, err := connectSQLite()
		if err != nil {
			panic("Could not connect to database.")
		}
		instance = &SQLiteDB{DB: db}
	})
	return instance
}

func connectSQLite() (*gorm.DB, error) {
	dbDir := "./data"
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		err := os.MkdirAll(dbDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	dbPath := filepath.Join(dbDir, "api.db")

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	
	return gorm.Open((sqlite.Open(dbPath)), config)
}