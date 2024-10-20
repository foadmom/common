package nexsql

import (
	"database/sql"
	"fmt"

	_ "github.com/microsoft/go-mssqldb"
	_ "github.com/microsoft/go-mssqldb/integratedauth/krb5"
)

// ============================================================================
//
// ============================================================================
func newConnection(driverName, connectionString string) (*sql.DB, error) {

	sqlObj, connectionError := sql.Open("mssql", connectionString)
	if connectionError != nil {
		fmt.Println(fmt.Errorf("error opening database: %v", connectionError))
	}

	return sqlObj, connectionError
}

func executeProc(procParams storedProcData) (string, error) {
	var _rc string

	_conn, _err := connection()
	if _err == nil {
		defer _conn.Close()
		_rc, _err = callStoredProc(_conn, procParams.ProcName, procParams.InputData)
		if _err == nil {
			fmt.Printf("return data=%s\n", _rc)
		}
	}
	return _rc, _err
}
