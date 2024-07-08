package utils

import (
	"database/sql"
	"fmt"

	"github.com/themadman159/starter-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IDatabase interface {
	ConnectMysql() (*gorm.DB, *sql.DB, error)
}

type Database struct{}

var (
	DB    *gorm.DB
	sqlDB *sql.DB
)

func NewDatabase() IDatabase {
	return &Database{}
}

func (r *Database) ConnectMysql() (*gorm.DB, *sql.DB, error) {

	dbConfig, err := config.LoadConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load config: %w", err)
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci&multiStatements=true",
		dbConfig.DB_USER,
		dbConfig.DB_PASSWORD,
		dbConfig.DB_HOST,
		dbConfig.DB_PORT,
		dbConfig.DB_NAME,
	)
	sqlDB, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	DB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open gorm connection: %v", err)
	}

	return DB, sqlDB, nil
}
