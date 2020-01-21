package gorm

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"time"
)

type DataSource struct {
	User            string `json:"username" wire:"user"`
	Type            string `json:"type" wire:"type"`
	Password        string `json:"password" wire:"password"`
	Database        string `json:"database" wire:"database"`
	Host            string `json:"host" wire:"host"`
	Port            int    `json:"port" wire:"port"`
	ConnMaxLifetime string `json:"connMaxLifetime" wire:"ConnMaxLifetime"`
	MaxIdleConns    int    `json:"maxIdleConns" wire:"maxIdleConns"`
	MaxOpenConns    int    `json:"maxOpenConns" wire:"maxOpenConns"`
}

var DataSourceSet = wire.NewSet(NewDataSource)

func NewDataSource() (dataSource *DataSource) {
	return &DataSource{
		User:     "",
		Password: "",
		Database: "",
		Host:     "",
		Port:     0,
	}
}

type Connection struct {
	DB *gorm.DB
}

var ConnectionSet = wire.NewSet(NewConnection)

func NewConnection(dataSource *DataSource) (connection *Connection, err error) {

	db, err := gorm.Open(dataSource.Type, "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	duration, err := time.ParseDuration(dataSource.ConnMaxLifetime)
	if err != nil {
		return nil, err
	}
	db.DB().SetConnMaxLifetime(duration)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool
	db.DB().SetMaxIdleConns(dataSource.MaxIdleConns)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(dataSource.MaxOpenConns)
	connection = &Connection{DB: db}
	return
}
