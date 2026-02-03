package sql

import (
	"database/sql"
	"fmt"

	l "github.com/foadmom/common/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresProperties DBProperties

var psgLogger l.LoggerInterface

type ExceptionPostgres struct {
	Message string `json:"message"`
	Details string `json:"details"`
	Hint    string `json:"hint"`
	Context string `json:"context"`
}

// ============================================================================
//
//	setup any
//
// ============================================================================
func init() {
	psgLogger = l.Instance()
	var _pgx PostgresProperties = PostgresProperties{"localPostgres", "pgx", "localhost", "5432",
		"postgres", "postgres", "postgres", "mycoach", ""}
	_pgx.ConnString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", _pgx.UserId, _pgx.Password, _pgx.Host, _pgx.Port, _pgx.Database)
	DBServers[_pgx.Name] = DBProperties(_pgx)
}

// ============================================================================
//
// ============================================================================
func (p *PostgresProperties) NewConnection() (*sql.DB, error) {
	_pgx := GetDBProperty(p.Name)
	_conn, _err := _pgx.NewConnection()
	if _err != nil {
		psgLogger.Printf(l.Error, "Unable to connect to database: %s\n", _err.Error())
	}
	return _conn, _err
}

// ============================================================================
//
// ============================================================================
func (p *PostgresProperties) CallStoredProc(conn *sql.DB, funcName string, query string) (string, error) {
	_jsonResult, _err := (*DBProperties)(p).CallStoredProc(conn, funcName, query)
	return _jsonResult, _err
}
