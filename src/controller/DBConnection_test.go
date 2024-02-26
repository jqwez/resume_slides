package controller

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseConnection(t *testing.T) {
	db := GetDatabaseConnection()
	assert.NotNil(t, db)
}