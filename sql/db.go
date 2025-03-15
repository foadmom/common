package sql

import (
	"database/sql"
)

type StoredProcData struct {
	ProcName  string
	InputData string
}

type DBProperties struct {
	Name       string
	Driver     string
	Host       string
	Port       string
	UserId     string
	Password   string
	Database   string
	ConnString string
}

type DBInterface interface {
	NewConnection() (*sql.DB, error)
	CallStoredProc(c *sql.DB, f string, q string) (string, error)
}

var DBServers map[string]DBProperties = make(map[string]DBProperties)

// ============================================================================
//
// ============================================================================
func init() {
}

// ============================================================================
//
// ============================================================================
func AddDBProperty(key string, value DBProperties) {
	DBServers[key] = value
}

// ============================================================================
//
// ============================================================================
func GetDBProperty(key string) DBProperties {
	return DBServers[key]
}

// ============================================================================
//
// ============================================================================
func (db *DBProperties) NewConnection() (*sql.DB, error) {
	_conn, _err := sql.Open(db.Driver, db.ConnString)
	if _err != nil {
		logger.Printf("Unable to connect to database: %v\n", _err)
	}
	return _conn, _err
}

// ============================================================================
//
// ============================================================================
func (p *DBProperties) CallStoredProc(conn *sql.DB, funcName string, query string) (string, error) {
	var _jsonResult string
	var _queryString string = "SELECT * FROM " + funcName + "('" + query + "')"
	_row := conn.QueryRow(_queryString)
	_err := _row.Scan(&_jsonResult)
	if _err != nil {
		logger.Printf("QueryRow failed: %v\n", _err)
	}

	return _jsonResult, _err
}
