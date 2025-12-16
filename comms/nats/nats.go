package nats

import (
	"log"
	"time"

	"github.com/foadmom/common/utils"

	l "github.com/foadmom/common/logger"
	ng "github.com/nats-io/nats.go"
)

// import (
//     ns "github.com/nats-io/nats-server/v2/server"
//     ng "github.com/nats-io/nats.go"
// )

type natsConfig struct {
	hosts    string
	user     string
	password string
}

var config natsConfig = natsConfig{ng.DefaultURL, "foadm", "Pa55w0rd"}

type Nats struct {
	Connection   *ng.Conn
	Subscription *ng.Subscription
}

var nats Nats = Nats{}
var hostName string
var hostNamePostfix string
var natsLogger l.LoggerInterface

func init() {
	natsLogger = l.Instance()
	_err := nats.Connect()
	if _err != nil {
		log.Fatal(_err)
	} else {
		hostName = utils.HostName()
		hostNamePostfix = hostName + "."
	}
}

func (n *Nats) Connect() error {
	var _err error

	n.Connection, _err = ng.Connect(config.hosts, ng.UserInfo(config.user, config.password))
	if _err == nil {
		// defer n.Connection.Close()
		n.Subscription, _err = n.Connection.SubscribeSync("appGateway.*")
		if _err != nil {
			natsLogger.Printf(l.Fatal, "nats connection failed. %s", _err.Error())
		}
	}

	return _err
}

// ========================================================
// ========================================================
// ========================================================
// Mandatory functions to satisfy the interface
// ========================================================
func Instance() Nats {
	return nats
}

func (n Nats) GetMessage(channel string) ([]byte, error) {
	var _message *ng.Msg
	var _err error

	_message, _err = n.Subscription.NextMsg(10 * time.Second)
	if _err != nil {
		natsLogger.Printf(l.Fatal, "nats subscription to channel %s failed. %s", channel, _err.Error())
	}
	return []byte(_message.Data), _err
}

func (n Nats) PutMessage(channel string, b []byte) error {
	var _err error
	var _queue string = channel + hostNamePostfix

	_err = n.Connection.Publish(_queue, b)
	return _err
}

// ========================================================
// ========================================================
