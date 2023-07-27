package message

import (
	"time"

	"github.com/foadmom/common/utils"
	c "github.com/foadmom/common/utils"
)

const HEART_BEAT_MESSAGE_CODE = "HeartBeat"

const ACTIVE = "active"
const DRAINING = "draining"
const DRAINED = "drained"
const DISABLED = "disabled"
const ERROR = "error"
const UNKNOWN = "unknown"

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

	h.UUID, _err = utils.GenerateUUID()
	if _err == nil {
		h.Version = VERSION
		h.Host = c.HostName()
		h.TimeStamp = time.Now()
	}
	return _err
}

// ========================================================
//
// ========================================================
func (g *GenericMessage) Instance() error {
	var _header MessageHeader

	_err := _header.Instance()
	if _err == nil {
		g.Header = _header
	}

	return _err
}
