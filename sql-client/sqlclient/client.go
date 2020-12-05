package sqlclient

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

const (
	goEnvironment = "GO_ENVIRONMENT"
	production    = "production"
)

var (
	dbClient SqlClient
)

type client struct {
	db *sql.DB
}

type SqlClient interface {
	Query(query string, args ...interface{}) (rows, error)
}

func isProduction() bool {
	return os.Getenv(goEnvironment) == production
}

func Open(driverName, dataSourceName string) (SqlClient, error) {

	// モック有効かつ本番環境ではない場合、モックを返す
	if isMocked && !isProduction() {
		dbClient = &clientMock{}
		return dbClient, nil
	}

	if driverName == "" {
		return nil, errors.New("invalid driver name")
	}
	database, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	dbClient := &client{
		db: database,
	}

	return dbClient, nil
}

func (c client) Query(query string, args ...interface{}) (rows, error) {
	returnedRows, err := c.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	result := sqlRows{
		rows: returnedRows,
	}

	return &result, nil
}
