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
	UUID       string    `json:"UUID"`
	SequenceID int       `json:"sequenceID"`
	Version    string    `json:"version"`
	Host       string    `json:"host"`
	Process    string    `json:"process"`
	UserID     string    `json:"userId"`
	Password   string    `json:"password"`
	TimeStamp  time.Time `json:"timestamp"`
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
