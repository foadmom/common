package utils

import (
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
)

var HOSTNAME string = ""

// ========================================================
//
// ========================================================
func init() {
	HOSTNAME, _ = os.Hostname() // not really needed
}

// ========================================================
//
// ========================================================
func HostName() string {
	if HOSTNAME == "" {
		HOSTNAME, _ = os.Hostname()
	}
	return HOSTNAME
}

// ========================================================
//
// ========================================================
func GenerateUUID() (string, error) {
	var _uuid uuid.UUID
	var _err error

	_uuid, _err = uuid.NewUUID()

	return _uuid.String(), _err
}

// ========================================================
// the lower limit is inclusive and upper limit is excluded.
// so the random int is anything from lower to upper-1
// ========================================================
func GenerateRandomInt(lower int, upper int) int {
	rand.Seed(time.Now().UnixNano())
	var _rand int = rand.Intn(upper-lower) + lower
	return _rand
}
