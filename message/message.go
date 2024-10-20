package message

import (
	"encoding/json"
	"time"

	c "github.com/foadmom/common/utils"
)

// ============================================================================
// reveive a json message and unmarshal it to genericMessage
// read the requestCode from payload
//
//	for Health/status Unmarshal Data into HealthStatus
//
// ============================================================================
const VERSION = "V1.0"

type MessageHeader struct {
	UUID       string    `json:"UUID"`		// * added automatically, a function will generate this. 
							//   Unless one has already been passed to you from above
	SequenceID int       `json:"sequenceID"`	//   only if there are multiple message in one transaction
	Version    string    `json:"version"`		// * Should be set by calling app
	Host       string    `json:"host"`		// * Should be added automatically by the app
	Process    string    `json:"process"`		//   calling application name. Should be set automatically
	UserID     string    `json:"userId"`		//   only used when API authentication is needed
	Password   string    `json:"password"`
	TimeStamp  time.Time `json:"timestamp"`		// * set automatically to Now ()
}

type Payload struct {
	MessageCode string `json:"messageCode"`
	Data        []byte `json:"data"`
}

type GenericMessage struct {
	Header  MessageHeader `json:"header"`
	Payload Payload       `json:"payload"`
}

// ========================================================
//
// ========================================================
func (h *MessageHeader) Instance() error {
	var _err error

	h.UUID, _err = c.GenerateUUID()
	if _err == nil {
		h.Version = VERSION
		h.Host = c.HostName()
		h.TimeStamp = time.Now()
	}
	return _err
}

// ========================================================
// create and instance of GenericMessage with whatever
// defaults we know before specifics are populated by
// the app
// ========================================================
func (g *GenericMessage) Instance() error {
	var _header MessageHeader

	_err := _header.Instance()
	if _err == nil {
		g.Header = _header
	}

	return _err
}

// ========================================================
// Marshall a GenericMessage
// ========================================================
func (gm GenericMessage) JSONMarshal() ([]byte, error) {

	return json.Marshal(gm)
}

// ========================================================
// Unmarshall a GenericMessage
// ========================================================
func (gm *GenericMessage) JSONUnmarshal(buffer []byte) error {

	return json.Unmarshal(buffer, gm)
}
