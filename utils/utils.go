package utils

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
	// qrcode "github.com/skip2/go-qrcode"
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
