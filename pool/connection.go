package pool

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

const (
	INITIAL_POOL_SIZE int    = 10
	DIALECT           string = "mysql"
)

type IConnectionPool interface {
	GetConnection() (db *gorm.DB)
	ReleaseConnection(db *gorm.DB) bool
	GetSize() (index int)
	/*	GetUrl() string
		GetUser() string
		GetPassword() string*/
}

type BasicConnectionPool struct {
	Url             string     `json:"url"`
	User            string     `json:"user"`
	Password        string     `json:"password"`
	DbName          string     `json:"dbName"`
	ConnectionPool  []*gorm.DB `json:"connectionPool"`
	UsedConnections []*gorm.DB `json:"usedConnections"`
}

func (pool *BasicConnectionPool) Create() (err error) {
	dbUrl := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		pool.User,
		pool.Password, pool.DbName)
	for i := 0; i < INITIAL_POOL_SIZE; i++ {
		db, err := gorm.Open(DIALECT, dbUrl)
		if err != nil {
			return err
		}
		pool.ConnectionPool = append(pool.ConnectionPool, db)
	}
	return nil
}

func (pool *BasicConnectionPool) ReleaseConnection(db *gorm.DB) bool {
	pool.ConnectionPool = append(pool.ConnectionPool, db)
	for i := 0; i < len(pool.ConnectionPool); i++ {
		if db == pool.ConnectionPool[i] {
			pool.ConnectionPool = append(pool.ConnectionPool[:i], pool.ConnectionPool[+1:]...)
			return true
		}
	}
	return false
}

func (pool *BasicConnectionPool) GetConnection() (db *gorm.DB) {
	index := len(pool.ConnectionPool)
	if index <= 0 {
		return nil
	}
	connect := pool.ConnectionPool[index-1]
	pool.ConnectionPool = append(pool.ConnectionPool[:(index-1)], pool.ConnectionPool[+1:]...)
	pool.UsedConnections = append(pool.UsedConnections, connect)
	return connect
}

func (pool *BasicConnectionPool) GetSize() (index int) {
	connectionSize := len(pool.ConnectionPool)
	usedConnectionSize := len(pool.UsedConnections)
	return connectionSize + usedConnectionSize
}

