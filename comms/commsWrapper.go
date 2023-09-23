package comms

import "github.com/foadmom/common/comms/nats"

type CommsInterface interface {
	GetMessage(channel string) ([]byte, error)
	PutMessage(channel string, b []byte) error
}

func Instance() CommsInterface {
	return nats.Instance()
}
