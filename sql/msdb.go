package sql

import (
	"database/sql"
	"fmt"

	l "github.com/foadmom/common/logger"
	_ "github.com/microsoft/go-mssqldb"
)

type MsSqlProperties DBProperties

var sqlserverLogger l.LoggerInterface

// ============================================================================
//
//	setup any
//
// ============================================================================
func init() {

	// "sqlserver://SA:myStrong(!)Password@localhost:1433?database=tempdb"

	var _sqlMS MsSqlProperties = MsSqlProperties{"localMsSQL", "pgx", "localhost", "5432",
		"SA", "Pa55w0rd", "", "postgres", ""}
	_sqlMS.ConnString = fmt.Sprintf("sqlserver://%s:%s@%s:%s/%s", _sqlMS.UserId, _sqlMS.Password, _sqlMS.Host, _sqlMS.Port, _sqlMS.Database)
	DBServers[_sqlMS.Name] = DBProperties(_sqlMS)
}

// ============================================================================
//
// ============================================================================
func (ms *MsSqlProperties) NewConnection() (*sql.DB, error) {
	_mssql := GetDBProperty(ms.Name)
	_conn, _err := Connect(&_mssql)
	if _err != nil {
		sqlserverLogger.Printf(l.Error, "Unable to connect to database: %s", _err.Error())
	}
	return _conn, _err
}

// ============================================================================
//
// ============================================================================
func (ms *MsSqlProperties) CallStoredProc(conn *sql.DB, funcName string, query string) (string, error) {
	_jsonResult, _err := CallStoredProc(conn, funcName, query)

	return _jsonResult, _err
}
