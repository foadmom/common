package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	l "github.com/foadmom/common/logger"
)

// ============================================================================
// This is a common interface to all functions for getting configuration
//   data for the apps
// The functions here are meant to be generic and hide the actual
//   implementation from the calling applications.
// eg. functions here could be calling ConfigStore in BitBucket or
//   etcd or any other config platform, but the caller should not be concerned
//   with that and should only be asking for what it needs.
// config structure:
//      Environment: "EName"   // eg dev, test1, perf, prod, .....
//          [
//              Category: "CName"        // eg LocationsDB
//              [
//                  {key: value}
//              ]
//          ]
// ============================================================================

// ==================================================================
// The cache and map below are for all the configs for one env
// not sure if we ever need to have the configs for all the envs in a cache
// not sure if we need both or just one of them
// ==================================================================
var configCache []byte
var configMap map[string]string
var _logger l.LoggerInterface

// ==================================================================
// initializes the config system
// ==================================================================
func init() {
	_logger = l.Instance()
	_logger.Print(l.Trace, "config package initialized")
}

// ==================================================================
// returns a JSON or any other format, eg YAML, as string.
// this returns everything under the Environment branch.
//
//	Can be used to cache all configs for a particular env
//
// ==================================================================
func GetConfigEnv(env string) (string, error) {
	return "", nil
}

// ==================================================================
// returns a JSON or any other format, eg YAML, as string
// ==================================================================
func GetConfigCategory(env, cat string) (string, error) {
	return "", nil
}

// ==================================================================
//
// ==================================================================
func GetConfigValue(env, cat, key string) (string, error) {
	return "", nil
}

// ==================================================================
//
// ==================================================================
func ReadConfigFile(fullPathAndFileName string) (string, error) {
	_config, _err := os.ReadFile(fullPathAndFileName)
	if _err != nil {
		_logger.Printf(l.Info, "unable to read file: %v", _err)
	}
	return string(_config), _err
}

// ==================================================================
// this is a common function to map a JSON string to any struct
// ==================================================================
func MapConfig(configStr string, configObj interface{}) error {
	_err := json.Unmarshal([]byte(configStr), configObj)
	return _err
}

// ==================================================================
// looks for a key in the nested map and returns a
// map[string]interface{} for the key.
// ==================================================================
func FindValueInJson(jsonStr, key string) (map[string]interface{}, bool) {
	var _mapData map[string]interface{}
	var _exists bool = false

	_keys := strings.Split(key, ".")
	err := json.Unmarshal([]byte(jsonStr), &_mapData)

	if err == nil {
		for _, k := range _keys {
			_mapData, _exists = SearchForKeyInMap(_mapData, k)
			if !_exists {
				return nil, false
			}
		}
	}
	return _mapData, _exists
}

// ==================================================================
// cheks if the key exists in the map and if it does.
// ==================================================================
func SearchForKeyInMap(thisMap map[string]interface{}, key string) (map[string]interface{}, bool) {
	var exists bool

	value, exists := thisMap[key]
	if exists {
		_result, ok := value.(map[string]interface{})
		if ok {
			return _result, true
		}
	}
	return nil, false
}

// ==================================================================
// find a nested JSON object for the compound key and return it as a JSON string
// eg key can be Environment.Database and it will return the JSON string for that category
// ==================================================================
func FindKeyedSubJson(fileName, key string) ([]byte, error) {
	var _jsonString string
	var _err error
	var _jResult []byte

	_jsonString, _err = ReadConfigFile(fileName)
	if _err == nil {
		_mapData, _exists := FindValueInJson(_jsonString, key)
		if _exists {
			_jResult, _err = json.Marshal(_mapData)
			return _jResult, _err
		}
	}
	return nil, _err
}

// ============================================================================
// this is just to test the config package and make sure it can read the
// config file and unmarshal it into the map[string]interface{}
// ============================================================================
func GetConfigFromFile(fileName string) (map[string]interface{}, error) {
	var _err error
	var ConfigData string

	// var _envConfig envConfig = envConfig{}
	ConfigData, _err = ReadConfigFile(fileName)

	// unmarshal the config data into a map to get the environment specific config
	var _data map[string]interface{}
	_err = json.Unmarshal([]byte(ConfigData), _data)
	if _err != nil {
		_logger.Printf(l.Fatal, "unable to process config file: %v", _err)
	}
	return _data, _err
}

// ============================================================================
// search the map for a keyed nested map and return it as a map[string]interface{}
// ============================================================================
func GetKeyMap(data map[string]interface{}, key string) (map[string]interface{}, error) {
	var _err error
	var _found bool

	var value map[string]interface{}

	value, _found = data[key].(map[string]interface{})
	if !_found {
		_err = fmt.Errorf("key %s not found in the map", key)
		_logger.Printf(l.Fatal, "environment %s not found in config file", key)
	}
	return value, _err
}

// ============================================================================
//
// ============================================================================
func GetKeyedStringValue(data map[string]string, key string) (string, bool) {
	_value, _found := data[key] // .(map[string]interface{})
	if !_found {
		_logger.Printf(l.Info, "environment %s not found in config file", key)
	}
	return _value, _found
}
