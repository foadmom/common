// {
//     // Use IntelliSense to learn about possible attributes.
//     // Hover to view descriptions of existing attributes.
//     // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
//     "version": "0.2.0",
//     "configurations": [
//         {
//             "name": "Launch Package",
//             "type": "go",
//             "request": "launch",
//             "mode": "auto",
//             "program": "${fileDirname}",
//             "args": ["-env", "dev", "-config", "./test.config.json"]
//         }
//     ]
// }

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	s "database/sql"

	// "ezpkg.io/errorz"
	// iterjson "ezpkg.io/iter.json"
	"github.com/foadmom/common/config"
	h "github.com/foadmom/common/http"
	q "github.com/foadmom/common/sql"

	// h "github.com/foadmom/common/http"
	l "github.com/foadmom/common/logger"
	"github.com/foadmom/common/sql"
)

type databaseConfig struct {
	Server     string `json:"server"`
	Port       string `json:"port"`
	User       string `json:"user"`
	Password   string `json:"password"`
	Database   string `json:"database"`
	Schema     string `json:"schema"`
	DriverName string `json:"DriverName"`
}

type httpConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type envConfig struct {
	Database databaseConfig `json:"database"`
	HTTP     httpConfig     `json:"http"`
}

var _logger l.LoggerInterface
var LoggerConfig l.Config = l.Config{true, false, true, "./", "testScaffold.log", 100, 7, 30}

// // Enable console logging
// ConsoleLoggingEnabled bool

// // EncodeLogsAsJson makes the log framework log JSON
// EncodeLogsAsJson bool
// // FileLoggingEnabled makes the framework log to a file
// // the fields below can be skipped if this value is false!
// FileLoggingEnabled bool
// // Directory to log to to when filelogging is enabled
// Directory string
// // Filename is the name of the logfile which will be placed inside the directory
// Filename string
// // MaxSize the max size in MB of the logfile before it's rolled
// MaxSize int
// // MaxBackups the max number of rolled files to keep
// MaxBackups int
// // MaxAge the max age in days to keep a logfile
// MaxAge int

func main() {
	_logger = l.Instance()
	_logger.Configure(LoggerConfig)
	l.SetLogLevel(l.Trace)
	commandLineArgs()
	getConfigFromFile(ConfigFile)
	// json_iterator()
	// TestcHttp()
	TestSQL()
	// testGenerateStoredProcWrapper()
}

type configLevel struct {
	Environment struct {
		Dev envConfig `json:"dev"`
	} `json:"environment"`
}

var configurations configLevel = configLevel{}

var Env string
var ConfigFile string
var ConfigData string

func commandLineArgs() {
	flag.StringVar(&Env, "env", "dev", "which environment you are running in")
	flag.StringVar(&ConfigFile, "config", "./test.config.json", "need a config file")
	flag.Parse()
	fmt.Println("environment", Env)
}

func getConfigFromFile(fileName string) {
	var _err error
	ConfigData, _err = config.ReadConfigFile(ConfigFile)
	if _err == nil {
		_logger.Printf(l.Info, "Config Data: %s", ConfigData)
		_err = config.MapConfig(ConfigData, &configurations)
		if _err == nil {
			_logger.Printf(l.Info, "configuration:  %v", configurations)
		}
	}
	if _err != nil {
		_logger.Printf(l.Fatal, "unable to read config file: %v", _err)
	}
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
	var _post q.PostgresProperties = q.PostgresProperties{}

	_prop := sql.DBProperties{
		Name:     "localPostgres",
		Driver:   configurations.Environment.Dev.Database.DriverName,
		Host:     configurations.Environment.Dev.Database.Server,
		Port:     configurations.Environment.Dev.Database.Port,
		UserId:   configurations.Environment.Dev.Database.User,
		Password: configurations.Environment.Dev.Database.Password,
		Database: configurations.Environment.Dev.Database.Database,
		Schema:   configurations.Environment.Dev.Database.Schema,
	}

	_post.Setup(_prop)

	conn, err := _post.Connect(_prop.Name)
	if err != nil {
		_logger.Printf(l.Error, "Unable to connect to database: %s\n", err.Error())
		return
	}
	defer conn.Close()

	jsonResult, err := _post.CallStoredProc(conn, "network.operator_get_all", "''")
	// jsonResult, err := TestGetUserId(PostgresProperties, conn, 1)

	_logger.Printf(l.Info, "Stored Procedure Result: %s\n", jsonResult)
	_logger.Print(l.Trace, "exiting TestSQL")
}

// ===============================
// ===============================
type RecordInfo struct {
	RecordStatus  string `json:"record_status"`
	CreatedBy     string `json:"created_by"`
	CreatedOn     string `json:"created_on"`
	LastUpdatedBy string `json:"last_updated_by"`
	LastUpdatedOn string `json:"last_updated_on"`
}

// ===============================
// ===============================
type zone struct {
	Id        int    `json:"id"`
	ShortName string `json:"short_name"`
	LongName  string `json:"long_name"`
	RecordInfo
}

type result struct {
	Data      zone                  `json:"data"`
	Exception sql.ExceptionPostgres `json:"exception"`
}

type resultTemplate struct {
	Rc     int    `json:"rc"`
	Result result `json:"result"`
}

func TestGetUserId(dbp sql.PostgresProperties, conn *s.DB, id int) (string, error) {
	var funcName string = "myCoach.zone_get_js"
	var query string = fmt.Sprintf(`{"id": %d}`, id)

	var _jsonResult string
	var _err error
	var _structResult resultTemplate = resultTemplate{}
	_jsonResult, _err = dbp.CallStoredProc(conn, funcName, query)
	if _err == nil {
		_logger.Printf(l.Info, "result = %s", _jsonResult)
		json.Unmarshal([]byte(_jsonResult), &_structResult)
		_logger.Printf(l.Info, "zone = %v", _structResult)
		_logger.Printf(l.Info, "created_on %s, last_updated_on %s", _structResult.Result.Data.CreatedOn, _structResult.Result.Data.LastUpdatedOn)
		// fmt.Printf("%v\n", _structResult)
	} else {

		return "", fmt.Errorf("error calling stored procedure: %w\n", _err)
	}

	return _jsonResult, _err
}

var sampleConfig string = `
{
	"environment":
	{
		"dev": 
		{
			"http":
			{
				"host": "localhost",
				"port": "8001"
			},
			"database": 
			{
				"server": "localhost", 
				"port": "5432", 
				"user": "postgres", 
				"password": "postgres", 
				"database": "postgres", 
				"schema": "test",
				"DriverName": "pgx"
			}
		}
	}
}`

// func generateStoredProcWrapper(procName string, jsonInput string) string {

// 	// _output :=  + procName + "param(input json) RETURNS text AS $$"

// 	_output := fmt.Sprintf("CREATE OR REPLACE FUNCTION %s (input json) RETURNS text AS $$\n    DECLARE\n", procName)
// 	for _item, _err := range iterjson.Parse([]byte(jsonInput)) {
// 		if _err == nil {
// 			// if _item.Key.Type() != 0 {
// 			if _item.IsObjectValue() {
// 				_tokenValue, _ := _item.GetValue()
// 				_output += fmt.Sprintf("        %s %s;\n", _item.Key, _tokenValue)
// 				_output += fmt.Sprintf("        %s TEXT := input::json->>'%s';\n", _item.Key, _item.Token)
// 			}
// 		}

// 	}

// 	return _output
// 	// DECLARE
// 	// --    _input text :=  '{"username": "johndoe", "age": 30, "email": "johndoe@example.com"}';
// 	//     _username text;
// 	// BEGIN
// 	//     _username = input::json->'username';
// 	//     RETURN input;
// 	// END;
// 	// $$ LANGUAGE plpgsql;

// 	// use iterjson to parse a json config and generate a stored procedure wrapper
// }

// var sampleJsonInput string = `{"username": "johndoe", "age": 30, "contacts": {"email": "johndoe@example.com","mobile": "07462 666 666"}, "accountHolder": true}`
var sampleJsonInput string = `
{
  "username": "johndoe",
  "age": 30,
  "contacts": {
    "email": "johndoe@example.com",
    "phones": {
      "mobile": "07462 666 666",
      "work": "07462 666 667"
    }
  },
  "accountHolder": true
}`

var operator_network_1 string = `{"operator_code": "NEX", "network_code": "EN", "trans_mode":"COACH", "country_iso2":"GB"}`

func testGenerateStoredProcWrapper() {
	procName := "test.my_stored_procedure"

	_output, _err := q.GenerateStoredProcWrapper(procName, operator_network_1)
	fmt.Printf("%s %v\n", _output, _err)
}
