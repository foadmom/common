package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	// qrcode "github.com/skip2/go-qrcode"
)

var HOSTNAME string = ""

// ========================================================
// initializes the HOSTNAME variable with the machine's hostname.
// This is not really needed, but it can be useful for logging or debugging purposes.
// ========================================================
func init() {
	HOSTNAME, _ = os.Hostname() // not really needed
}

// ========================================================
// returns the machine's hostname. If the HOSTNAME variable is empty, it initializes it first.
// ========================================================
func HostName() string {
	if HOSTNAME == "" {
		HOSTNAME, _ = os.Hostname()
	}
	return HOSTNAME
}

// ========================================================
// generates a UUID string. This can be used for generating unique IDs for nodes, messages, etc.
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
	var _rand int = rand.Intn(upper-lower) + lower
	return _rand
}

type simpleQRCode struct {
	Content string
	Size    int
}

// ========================================================
//
// ========================================================
func (code *simpleQRCode) Generate() ([]byte, error) {
	qrCode, err := qrcode.Encode(code.Content, qrcode.Medium, code.Size)
	if err != nil {
		return nil, fmt.Errorf("could not generate a QR code: %v", err)
	}
	return qrCode, nil
}

// ========================================================
//
// ========================================================
func QRCode(content string, size int) ([]byte, error) {
	var _err error
	var _codeData []byte

	_qrCode := simpleQRCode{Content: content, Size: size}
	_codeData, _err = _qrCode.Generate()

	return _codeData, _err
}

// ==================================================================
//
// ==================================================================
func IPAddresses(hostname string) ([]net.IP, error) {
	_addrs, _error := net.LookupIP(hostname)
	// lookupIP looks up host using the local resolver.
	if _error != nil {
		log.Println("Failed to detect machine host name. ", _error.Error())
	}
	return _addrs, _error
}
