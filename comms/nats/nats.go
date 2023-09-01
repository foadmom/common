package comms

import (
	"log"
	"time"

	ng "github.com/nats-io/nats.go"
)

// import (
//     ns "github.com/nats-io/nats-server/v2/server"
//     ng "github.com/nats-io/nats.go"
// )

type Nats struct {
	Connection   *ng.Conn
	Subscription *ng.Subscription
}

var nats *Nats = &Nats{}

func init() {
	_err := nats.Connect()
	if _err != nil {
		log.Fatal(_err)
	}
}

func (n *Nats) Connect() error {
	var _err error

	n.Connection, _err = ng.Connect(ng.DefaultURL, ng.UserInfo("foadm", "Pa55w0rd"))
	if _err == nil {
		// defer n.Connection.Close()
		n.Subscription, _err = n.Connection.SubscribeSync("appGateway.*")
		if _err == nil {

		}
	} else {
		log.Fatal(_err)
	}

	return _err
}

// ========================================================
// ========================================================
// ========================================================
// Mandatory functions to satisfy the interface
// ========================================================
func Instance() *Nats {
	return nats
}

func (n *Nats) GetMessage() ([]byte, error) {
	var _message *ng.Msg
	var _err error

	_message, _err = n.Subscription.NextMsg(10 * time.Second)
	if _err != nil {
		log.Fatal(_err)
	}
	return []byte(_message.Data), _err
}

func (n *Nats) PutMessage(b []byte) error {
	return nil
}

// ========================================================
// ========================================================
