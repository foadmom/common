package message

import (
	"testing"
)

func TestGenericMessage_Instance(t *testing.T) {
	var _message GenericMessage

	_err := _message.Instance()
	if _err == nil {

	} else {
		t.Errorf("Error generating Generic message %v\n", _message)
	}
}

func Test_MessageMarshal(t *testing.T) {
	var _message GenericMessage

	_err := _message.Instance()
	if _err == nil {
		var _json []byte = make([]byte, 0, 1024)
		_json, _err = _message.JSONMarshal()
		if _err == nil {
			var _message2 GenericMessage
			_err = _message2.JSONUnmarshal(_json)
			if _err == nil {
				// if reflect.DeepEqual(_message2, _message) {
				if _message.Header == _message2.Header {
				} else {
					t.Errorf("marshal and unmarshall did not produce the same struct. \nexpected  %v\ngot %v\n", _message, _message2)
				}

			}
		}
	}
	if _err != nil {
		t.Errorf("Error generating Generic message %v\n", _message)
	}

}
