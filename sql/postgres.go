package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"

	l "github.com/foadmom/common/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresProperties DBProperties

var psgLogger l.LoggerInterface

// ============================================================================
//
//	setup any
//
// ============================================================================
func init() {
	psgLogger = l.Instance()
	var _pgx PostgresProperties = PostgresProperties{"localPostgres", "pgx", "localhost", "5432",
		"postgres", "postgres", "", "postgres"}
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

// ============================================================================
// this is a function to produce a wrapper for a postgres function that takes a
// json input and creates local psql variables from json keys and values.
// ============================================================================
func generateStoredProcWrapper_2(procName string, jsonInput string) (string, error) {
	var _data map[string]interface{}
	var _output string

	_err := json.Unmarshal([]byte(jsonInput), &_data)
	if _err == nil {
		_output = fmt.Sprintf("CREATE OR REPLACE FUNCTION %s (input json) RETURNS text AS $$\n    DECLARE\n", procName)
		for _key, _value := range _data {
			switch v := _value.(type) {
			case string:
				_output += fmt.Sprintf("        _%s TEXT := input::json->>'%s';\n", _key, _key)
			case float64:
				_output += fmt.Sprintf("        _%s NUMERIC := (input::json->>'%s')::NUMERIC;\n", _key, _key)
			case bool:
				_output += fmt.Sprintf("        _%s BOOLEAN := (input::json->>'%s')::BOOLEAN;\n", _key, _key)
			default:
				return "", fmt.Errorf("unsupported type for key %s: %T", _key, v)
			}
		}
		_output += "    BEGIN\n"
		_output += "        -- function body here\n"
		_output += "        RETURN input;\n"
		_output += "    END;\n"
		_output += "$$ LANGUAGE plpgsql;\n"
	}
	return _output, _err
}
