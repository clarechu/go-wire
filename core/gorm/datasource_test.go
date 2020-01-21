package gorm

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDataSource(t *testing.T) {
	dataSource := NewDataSource()
	ds := DataSource{
		User:            "",
		Type:            "",
		Password:        "",
		Database:        "",
		Host:            "",
		Port:            0,
		ConnMaxLifetime: "",
		MaxIdleConns:    0,
		MaxOpenConns:    0,
	}
	assert.Equal(t, dataSource, &ds)
}

func TestNewConnection(t *testing.T) {
	_, err := InitializeConnection()
	assert.Equal(t, "sql: unknown driver \"\" (forgotten import?)", err.Error())
}
