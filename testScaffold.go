package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	s "database/sql"

	h "github.com/foadmom/common/cHttp"
	l "github.com/foadmom/common/logger"
	"github.com/foadmom/common/sql"
)

var _logger l.LoggerInterface

func main() {
	_logger = l.Instance()
	l.SetLogLevel(l.Trace)
	TestcHttp()
	// TestSQL()
}

func TestcHttp() {
	var _http *h.CommonHttp
	_http = _http.Init("localhost", "8001")
	h.AddHandler("/", localHttpHandler)
	_http.Listen()

	_logger.Print(l.Trace, "exiting testScaffold")
}

func localHttpHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s Hello\n", time.Now().String())
}

func TestSQL() {
	_logger = l.Instance()
	_logger.Print(l.Trace, "starting TestSQL")

	PostgresProperties := sql.PostgresProperties{"localPostgres", "pgx", "localhost", "5432",
		"postgres", "postgres", "", "postgres"}
	conn, err := PostgresProperties.NewConnection()
	if err != nil {
		_logger.Printf(l.Error, "Unable to connect to database: %s\n", err.Error())
		return
	}
	defer conn.Close()

	jsonResult, err := TestGetUserId(PostgresProperties, conn, 1)

	_logger.Printf(l.Info, "Stored Procedure Result: %s\n", jsonResult)
	_logger.Print(l.Trace, "exiting TestSQL")
}

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

func TestGetUserId(dbp sql.PostgresProperties, conn *s.DB, userId int) (string, error) {
	var funcName string = "test.get_user_by_id"
	var query string = fmt.Sprintf(`{"user_id": %d}`, userId)

	var _jsonResult string
	var _err error
	var _structResult User = User{}
	_jsonResult, _err = dbp.CallStoredProc(conn, funcName, query)
	if _err == nil {
		_logger.Printf(l.Info, "Stored Procedure %s executed successfully", funcName)
		json.Unmarshal([]byte(_jsonResult), &_structResult)
		_logger.Printf(l.Info, "UserID: %d, Username: %s, Email: %s",
			_structResult.Id, _structResult.LastName, _structResult.Email)
		fmt.Printf("%v\n", _structResult)
	} else {

		return "", fmt.Errorf("error calling stored procedure: %w\n", _err)
	}

	return _jsonResult, _err
}
