package server

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_getInstalledMigrationVersion(t *testing.T) {
	db, mock := GetMockDB()

	// Mock for successful query
	rows := sqlmock.NewRows([]string{"version"}).AddRow(5)
	mock.ExpectQuery("SELECT version FROM migrations ORDER BY version DESC LIMIT 1").WillReturnRows(rows)

	// Channel to collect the result
	resultChan := getInstalledMigrationVersionCh(db)
	version := <-resultChan

	// Validate
	assert.Equal(t, 5, version)
}
