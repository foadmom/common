package sql

import (
	"database/sql"

	l "github.com/foadmom/common/logger"
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
	Schema     string
	ConnString string
}

type DBInterface interface {
	NewConnection() (*sql.DB, error)
	CallStoredProc(c *sql.DB, f string, q string) (string, error)
}

var DBServers map[string]DBProperties = make(map[string]DBProperties)
var dbLogger l.LoggerInterface

// ============================================================================
//
// ============================================================================
func init() {
	dbLogger = l.Instance()
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
		dbLogger.Printf(l.Error, "Unable to connect to database: %s", _err.Error())
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
		dbLogger.Printf(l.Error, "QueryRow failed: %s", _err.Error())
	}

	return _jsonResult, _err
}
