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

	// "ezpkg.io/errorz"
	// iterjson "ezpkg.io/iter.json"
	"github.com/foadmom/common/config"
	h "github.com/foadmom/common/http"
	q "github.com/foadmom/common/sql"

	// h "github.com/foadmom/common/http"
	l "github.com/foadmom/common/logger"
)

type httpConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type envConfig struct {
	Database q.DBProperties `json:"Database"`
	HTTP     httpConfig     `json:"Http"`
}

var _logger l.LoggerInterface
var LoggerConfig l.Config = l.Config{ConsoleLoggingEnabled: true,
	EncodeLogsAsJson: false, FileLoggingEnabled: true,
	Directory: "./", Filename: "testScaffold.log",
	MaxSize: 100, MaxBackups: 7, MaxAge: 30}

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
	_config, _err := getConfigFromFile(ConfigFile)
	if _err != nil {
		_logger.Printf(l.Fatal, "Failed to get config: %v", _err)
		return
	}
	// TestcHttp()
	TestSQL(_config)
	testGenerateStoredProcWrapper()
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
	flag.StringVar(&Env, "env", "Dev", "which environment you are running in")
	flag.StringVar(&ConfigFile, "config", "./test.config.json", "need a config file")
	flag.Parse()
	fmt.Println("environment", Env)
}

// ============================================================================
// read the config file and unmarshal it into the envConfig struct
// ============================================================================
func getConfigFromFile(fileName string) (envConfig, error) {
	var _err error
	var _found bool
	var _envConfig envConfig = envConfig{}
	ConfigData, _err = config.ReadConfigFile(ConfigFile)

	// unmarshal the config data into a map to get the environment specific config
	var _data map[string]interface{}
	_err = json.Unmarshal([]byte(ConfigData), &_data)
	if _err == nil {
		// get the environment specific config
		var _myEnv map[string]interface{}

		_myEnv, _found = _data["Environment"].(map[string]interface{})
		if _found {
			// get the config for the specified environment, like Dev or Prod
			_myEnv, _found = _myEnv[Env].(map[string]interface{})
			if _found {
				// convert map to json and then unmarshal it into the envConfig struct
				var _envConfigJson []byte
				_envConfigJson, _err = json.Marshal(_myEnv)
				_err = json.Unmarshal(_envConfigJson, &_envConfig)
			}
		}
		if !_found {
			_err = fmt.Errorf("environment %s not found in config file", Env)
		}
	}
	if _err != nil {
		_logger.Printf(l.Fatal, "unable to process config file: %v", _err)
	}
	return _envConfig, _err
}

// ============================================================================
// this is a simple test function to test the http server
// ============================================================================
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

// ============================================================================
// take the postgres config and test a connection to the database
// and a call to a stored procedure
//
//	_prop := sql.DBProperties{
//		Name:     configurations.Environment.Dev.Database.Name,
//		Host:     configurations.Environment.Dev.Database.Host,
//		Port:     configurations.Environment.Dev.Database.Port,
//		User:     configurations.Environment.Dev.Database.User,
//		Password: configurations.Environment.Dev.Database.Password,
//		Database: configurations.Environment.Dev.Database.Database,
//		Schema:   configurations.Environment.Dev.Database.Schema,
//		Driver:   configurations.Environment.Dev.Database.Driver,
//	}
//
// ============================================================================
func TestSQL(config envConfig) {
	_logger = l.Instance()
	_logger.Print(l.Trace, "starting TestSQL")
	var _post q.PostgresProperties = q.PostgresProperties{}

	_post.Setup(config.Database)

	conn, err := _post.Connect(config.Database.Name)
	if err != nil {
		_logger.Printf(l.Error, "Unable to connect to database: %s\n", err.Error())
		return
	}
	defer conn.Close()

	// This should return a json with array of operators and no exception.
	jsonResult, err := _post.CallStoredProc(conn, "network.operator_get_all", "")
	_logger.Printf(l.Info, "Stored Procedure Result: %s\n", jsonResult)

	// this should fail with duplicate key error as the code EN already exists in the network table.
	jsonResult, err = _post.CallStoredProc(conn, "network.network_insert", `{"code": "EN", "name": "England"}`)
	_logger.Printf(l.Info, "Stored Procedure Result: %s\n", jsonResult)

	_logger.Print(l.Trace, "exiting TestSQL")
}

// CREATE TABLE network.service_link (
//     id              BIGSERIAL PRIMARY KEY,
//     service_id      BIGINT REFERENCES network.service(id),
//     from_stop_id    BIGINT REFERENCES network.stop(id),
//     to_stop_id      BIGINT REFERENCES network.stop(id),
//     distance_meters INTEGER,
//     sequence_order  INTEGER NOT NULL,

var jsonInput string = `{
  "service_code": "NX230_S",
  "from_stop_code": "DGBT",
  "to_stop_code": "COVCEN",
  "facilities": {
    "toilet": true,
    "bar": true,
    "cafe": true
  },
  "distance_meters": 20000,
  "sequence_order": 1
}`

func testGenerateStoredProcWrapper() {
	// this is a sample json input for the stored procedure wrapper generator
	procWrapper, err := q.GenerateStoredProcWrapper("network.service_link_insert", jsonInput)
	if err != nil {
		_logger.Printf(l.Error, "Error generating stored procedure wrapper: %v", err)
	} else {
		_logger.Printf(l.Info, "Generated stored procedure wrapper: %s", procWrapper)
	}
}
