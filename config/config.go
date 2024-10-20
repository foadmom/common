package config

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
