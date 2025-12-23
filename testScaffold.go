package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	s "database/sql"

	"flag"

	h "github.com/foadmom/common/cHttp"
	"github.com/foadmom/common/config"
	l "github.com/foadmom/common/logger"
	"github.com/foadmom/common/sql"
)

var _logger l.LoggerInterface

func main() {
	_logger = l.Instance()
	l.SetLogLevel(l.Trace)
	commandLineArgs()
	getConfigFromFile(ConfigFile)
	// TestcHttp()
	// TestSQL()
}

type configLevel struct {
	Environment struct {
		Dev struct {
			HTTP struct {
				Host string `json:"host"`
				Port string `json:"port"`
			} `json:"http"`
			Database struct {
				Server     string `json:"server"`
				Port       string `json:"port"`
				User       string `json:"user"`
				Password   string `json:"password"`
				Database   string `json:"database"`
				Schema     string `json:"schema"`
				DriverName string `json:"DriverName"`
			} `json:"database"`
		} `json:"dev"`
	} `json:"environment"`
}

var configurations configLevel = configLevel{}

var Env string
var ConfigFile string
var ConfigData string

func commandLineArgs() {
	// flag.StringVar(&Env, "env", "dev", "which environment you are running in")
	// flag.StringVar(&ConfigFile, "config", "", "need a config file")
	flag.Parse()
	fmt.Println("environment", Env)
}

func getConfigFromFile(fileName string) {
	var _err error
	ConfigData, _err = config.ReadConfigFile(ConfigFile)
	if _err == nil {
		_logger.Printf(l.Info, "Config Data: %s", ConfigData)
		_err = config.MapConfig(ConfigData, &configurations)
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
			}
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
