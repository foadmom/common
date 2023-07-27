package time

import (
	"encoding/json"
	"testing"
)

func TestUTCTime(t *testing.T) {
	var _goodTimes []string = []string{"\"2023-07-25T09:26:27.213838\"",
		"\"1901-01-01T00:01:09.123456\""}
	var _badTimes []string = []string{"\"2023-07-25T09:26:27.213\"",
		"\"1901-01-01T00:01:09Z\""}

	for _, _testTime := range _goodTimes {
		UTCTimeShouldSucceed(t, _testTime)
	}

	for _, _testTime := range _badTimes {
		UTCTimeShoulFail(t, _testTime)
	}
}

func UTCTimeShouldSucceed(t *testing.T, testTime string) {
	// var testTime string = "\"2023-07-25T09:26:27.213838\""
	var _time UTCTime
	_err := json.Unmarshal([]byte(testTime), &_time)
	if _err != nil {
		t.Errorf("Unmarshal of %s failed with error %v\n", testTime, _err)
	} else {
		_res, _err := json.Marshal(_time)
		if _err != nil {
			t.Errorf("Marshal of time failed with error %v\n", _err)
		} else {
			if string(_res) != testTime {
				t.Errorf("Marshal did not get back to original string of ")
			}
		}
	}
}

func UTCTimeShoulFail(t *testing.T, testTime string) {
	// var testTime string = "\"2023-07-25T09:26:27.213838\""
	var _time UTCTime
	_err := json.Unmarshal([]byte(testTime), &_time)
	if _err != nil {
		t.Logf("Unmarshal of %s failed with error %v\n", testTime, _err)
	} else {
		_res, _err := json.Marshal(_time)
		if _err != nil {
			t.Logf("Marshal of UTCTime failed with error %v\n", _err)
		} else {
			if string(_res) != testTime {
				t.Logf("Marshal did not get back to original string of ")
			}
		}
	}
	if _err == nil {
		t.Errorf("Test succeeded where it was expected to fail")
	}
}
