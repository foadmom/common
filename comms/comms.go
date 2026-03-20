package comms

import (
	"fmt"
	"net"
	"time"
)

type commsInterface interface {
	Init(config string) error
	Close()
	Send(message []byte, channel string) error
	Subscribe(channel string, goChannel chan []byte) error
	UnSubscribe(channel string) error
	Cleanup() error
}

var Comms commsInterface
var gConfig string

func init() {
	Comms = GetInstance()
}

// ======================================================
// Init initializes the communication channel using the provided configuration string.
// The configuration string is expected to be a JSON string that defines all the
// necessary parameters for initializing the communication channel (e.g., NATS).
// If the initialization fails, it returns an error. In a real implementation,
// you would want to handle this error appropriately, such as retrying the connection
// or logging the error for further investigation.
// ======================================================
func Init(config string) error {
	gConfig = config
	return Comms.Init(gConfig)
}

// ======================================================
// Close closes the communication channel. This is important to free up resources and prevent memory leaks.
// In a real implementation, you would want to ensure that all
// pending messages are sent and all subscriptions are cleaned up before closing the connection.
// ======================================================
func Close() {
	Comms.Close()
}

// ======================================================
// Send sends a message to a specified channel. If the send fails,
// it tries to reinitialize the connection and resend the message.
// This is a simple retry mechanism to handle transient connection issues.
// In a real implementation, you might want
// to add more sophisticated error handling and retry logic.
// ======================================================
func Send(message []byte, channel string) error {
	var _err error = Comms.Send(message, channel)
	if _err != nil {
		_err = Init(gConfig)
		if _err == nil {
			_err = Comms.Send(message, channel)
		}
	}
	return _err
}

// ======================================================
// Subscribe subscribes to a channel and sends received messages to the provided Go channel.
// This allows the node to receive messages from other nodes in the cluster.
// ======================================================
func Subscribe(channel string, goChan chan []byte) error {
	return Comms.Subscribe(channel, goChan)
}

// ======================================================
// UnSubscribe unsubscribes from a channel.
// This is important to prevent memory leaks and unnecessary message
// processing when a node changes its role or leaves the cluster.
// ======================================================
func UnSubscribe(channel string) error {
	return Comms.UnSubscribe(channel)
}

// ======================================================
// code provided by: Rustavil Nurkaev
// This is just to test connectivity to the other nodes in the cluster.
// ======================================================
func RawConnect(host string, ports []string) {
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			fmt.Println("Connecting error:", err)
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Opened", net.JoinHostPort(host, port))
		} else {
			fmt.Println("Failed to connect to", net.JoinHostPort(host, port))
		}
	}
}
