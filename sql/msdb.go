package nexsql

import (
	"database/sql"
	"errors"
	"fmt"
)

type instanceData struct {
	server   string
	port     string
	userId   string
	password string
	database string
}

var servers map[string]instanceData = make(map[string]instanceData)

func Init(instanceKey, server, port, userId, password, database string) {
	_, _found := servers[instanceKey]
	if _found == false {
		var _newInstance instanceData = instanceData{server: server, port: port, userId: userId,
			password: password, database: database}
		servers[instanceKey] = _newInstance
	}
}

func Connection(instanceKey string) (*sql.DB, error) {
	var _connectionError error
	var _sqlObj *sql.DB
	_instance, _found := servers[instanceKey]

	if _found == true {
		_connectionString := fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s",
			_instance.server, _instance.port, _instance.userId, _instance.password, _instance.database)

		_sqlObj, _connectionError = newConnection("mssql", _connectionString)
	} else {
		_connectionError = errors.New("connection has not been initialised")
	}

	return _sqlObj, _connectionError
}

// ============================================================================
//
// ============================================================================
func callStoredProc(conn *sql.DB, procName string, inputJson string) (string, error) {
	var _jsonOutput string
	_query := "EXECUTE " + procName + " ?"

	_stmt, _err := conn.Prepare(_query)
	if _err == nil {
		defer _stmt.Close()

		row := _stmt.QueryRow(inputJson)
		_err = row.Scan(&_jsonOutput)
	}

	return _jsonOutput, _err
}
