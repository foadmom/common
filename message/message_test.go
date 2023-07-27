package message

import "testing"

func TestGenericMessage_Instance(t *testing.T) {
	var _message GenericMessage

	_err := _message.Instance()
	if _err != nil {
		t.Errorf("Error generating Generic message %v\n", _message)
	}
}
