package sql

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
)

type MsSqlProperties DBProperties

// ============================================================================
//
//	setup any
//
// ============================================================================
func init() {

	"sqlserver://SA:myStrong(!)Password@localhost:1433?database=tempdb"

	var _sqlMS MsSqlProperties = MsSqlProperties{"localMsSQL", "pgx", "localhost", "5432",
		"SA", "Pa55w0rd", "", "postgres"}
	_sqlMS.ConnString = fmt.Sprintf("sqlserver://%s:%s@%s:%s/%s", _sqlMS.UserId, _sqlMS.Password, _sqlMS.Host, _sqlMS.Port, _sqlMS.Database)
	DBServers[_sqlMS.Name] = DBProperties(_sqlMS)
}

// ============================================================================
//
// ============================================================================
func (ms *MsSqlProperties) NewConnection() (*sql.DB, error) {
	_pgx := GetDBProperty(ms.Name)
	_conn, _err := _pgx.NewConnection()
	if _err != nil {
		logger.Printf("Unable to connect to database: %v\n", _err)
	}
	return _conn, _err
}

// ============================================================================
//
// ============================================================================
func (ms *MsSqlProperties) CallStoredProc(conn *sql.DB, funcName string, query string) (string, error) {
	_jsonResult, _err := (*DBProperties)(ms).CallStoredProc(conn, funcName, query)

	return _jsonResult, _err
}
