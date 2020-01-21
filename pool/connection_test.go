package pool

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicConnectionPool_Create(t *testing.T) {
	pool := &BasicConnectionPool{}
	err := pool.Create()
	assert.Equal(t, "sql: unknown driver \"mysql\" (forgotten import?)", err.Error())
}