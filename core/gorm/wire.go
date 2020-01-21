//+build wireinject

package gorm

// wire_gen.go
import "github.com/google/wire"

func InitializeConnection() (*Connection, error) {
	wire.Build(DataSourceSet, ConnectionSet)
	return &Connection{}, nil
}

func InitializeDataSource() *DataSource {
	wire.Build(DataSourceSet)
	return &DataSource{}
}