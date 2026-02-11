package sql

import (
	"database/sql"
	"encoding/json"
	"fmt"

	l "github.com/foadmom/common/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresProperties struct{}

// type PostgresProperties DBProperties

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
}

// ============================================================================
//
// ============================================================================
func (p *PostgresProperties) Setup(prop DBProperties) {
	prop.ConnString = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", prop.UserId,
		prop.Password, prop.Host, prop.Port, prop.Database)
	AddDBProperty(prop)
}

// ============================================================================
//
// ============================================================================
func (p *PostgresProperties) Connect(name string) (*sql.DB, error) {
	_pgx := GetDBProperty(name)
	_conn, _err := Connect(&_pgx)

	if _err != nil {
		psgLogger.Printf(l.Error, "Unable to connect to database: %s\n", _err.Error())
	}
	return _conn, _err
}

// ============================================================================
//
// ============================================================================
func (p *PostgresProperties) CallStoredProc(conn *sql.DB, funcName string, query string) (string, error) {
	_jsonResult, _err := CallStoredProc(conn, funcName, query)
	return _jsonResult, _err
}

// ============================================================================
// This function generates a stored procedure wrapper given a
// procedure name and a sample json input
// ============================================================================
func GenerateStoredProcWrapper(procName string, jsonInput string) (string, error) {
	var _data map[string]interface{}
	var _output string

	_err := json.Unmarshal([]byte(jsonInput), &_data)
	if _err == nil {
		_output = fmt.Sprintf("CREATE OR REPLACE FUNCTION %s (input json) RETURNS text AS $$\n    DECLARE\n", procName)
		_output, _ = generateParamsFromMap("", _output, _data)
		_output += "        _result TEXT;\n"
		_output += "    BEGIN\n"
		_output += "        -- function body here\n"
		_output += "        RETURN _result;\n"
		_output += "    END;\n"
		_output += "$$ LANGUAGE plpgsql;\n"
	}
	return _output, _err
}

// ============================================================================
// this is a recursive function to handle nested json objects
// This is how you get individual values from nested json in psql:
// _contacts_email TEXT := input::json#>>'{contacts, email}';
// ============================================================================
func generateParamsFromMap(prefix, output string, _data map[string]any) (string, error) {
	if prefix == "" { // && prefix[0] != 'v' { // which means it is the first level
		prefix = "v" + prefix
	}
	for _key, _value := range _data {
		switch v := _value.(type) {
		case string:
			output += fmt.Sprintf("        %s_%s TEXT := input::json->>'%s';\n", prefix, _key, _key)
		case float64:
			output += fmt.Sprintf("        %s_%s NUMERIC := (input::json->>'%s')::NUMERIC;\n", prefix, _key, _key)
		case bool:
			output += fmt.Sprintf("        %s_%s BOOLEAN := (input::json->>'%s')::BOOLEAN;\n", prefix, _key, _key)
		case map[string]any:
			output, _ = generateParamsFromMap(prefix+"_"+_key, output, _value.(map[string]any))
		default:
			return "", fmt.Errorf("unsupported type for key %s: %T", _key, v)
		}
	}
	return output, nil
}
