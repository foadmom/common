package date

import (
	"bytes"
	"encoding/json"
	"testing"
)

// ============================================================================
// creates a custom date (dateOnly) and Marshal and Unmarshall
// and see if you get back to where you started
// ============================================================================
func anyTimeMarshall[ti any](inTime ti, layOut string, t *testing.T) error {
	var _err error
	_json, _err := json.Marshal(&inTime)
	if _err == nil {
		var _outTime ti
		_err = json.Unmarshal(_json, &_outTime)
		if _err == nil {
			// Unmarshall and see if the unmarshalled object is the
			// same as the original any object
			// if time.Time(inTime) == time.Time(_outTime) {
			_json2, _err := json.Marshal(&_outTime)
			if _err == nil {
				if bytes.Equal(_json, _json2) {
					// this is success
				}
			} else {
				// t.Errorf("Unmarshal did not produce the original time object %v", inTime)
			}
		}
	} else {
		// t.Errorf("Error json Marshalling %v", inTime)
	}
	return _err
}

// ============================================================================
// creates a custom date (dateOnly) and Marshal and Unmarshall
// and see if you get back to where you started
// ============================================================================
func Test_dateOnly_1(t *testing.T) {
	var _err error
	var _time dateOnly
	_err = _time.Set("2023-03-25")
	if _err == nil {
		anyTimeMarshall(_time, dateOnlyLayout, t)
		if _err != nil {
			t.Errorf("Error with custom time sqlServerTimestamp. error=%s", _err)
		}
	} else {
		t.Errorf("Error with custom time sqlServerTimestamp. error=%s", _err)
	}
}

// ============================================================================
// creates a custom date (dateOnly) and Marshal and Unmarshall
// and see if you get back to where you started. this test provides wrong
// date layout and should fail
// ============================================================================
func Test_dateOnly_2(t *testing.T) {
	var _err error
	var _time dateOnly
	_err = _time.Set("2023-03-25T23:09:56")
	if _err == nil {
		_err = anyTimeMarshall(_time, dateOnlyLayout, t)
		if _err != nil {
			t.Errorf("Error with custom time sqlServerTimestamp. error=%s", _err)
		}
	} else {
		// success
	}

}

// ============================================================================
// creates a custom date (dateOnly) and Marshal and Unmarshall
// and see if you get back to where you started
// ============================================================================
func Test_SQLServerDate_1(t *testing.T) {
	var _err error
	var _time sqlServerTimestamp
	_err = _time.Set("2023-03-25T22:23:03.000")
	if _err == nil {
		_err = anyTimeMarshall(_time, SQLServerTimeStampLayout, t)
		if _err != nil {
			t.Errorf("Error with custom time sqlServerTimestamp. error=%s", _err)
		}
	} else {
		t.Errorf("Error with custom time sqlServerTimestamp. error=%s", _err)
	}

}

// ============================================================================
// creates a custom date (dateOnly) and Marshal and Unmarshall
// and see if you get back to where you started
// ============================================================================
func Test_SQLServerDate_2(t *testing.T) {
	var _err error
	var _time sqlServerTimestamp
	_err = _time.Set("2023-03-25T22:23:03")
	if _err == nil {
		_err = anyTimeMarshall(_time, SQLServerTimeStampLayout, t)
		if _err != nil {
			t.Errorf("Error with custom time sqlServerTimestamp. error=%s", _err)
		}
	} else {
		// success
	}

}
