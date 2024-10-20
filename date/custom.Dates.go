package date

import (
	"fmt"
	"strings"
	"time"
)

// ============================================================================

// =============================================================================
// =============================================================================
func CustomTimeMarshalJSON(ctime time.Time, DateLayout string) ([]byte, error) {
	var _timeString string

	_timeString = ctime.Format(DateLayout)
	_timeString = fmt.Sprintf("%q", _timeString)

	return []byte(_timeString), nil

}

// ============================================================================
func CustomTimeUnMarshalJSON(json []byte, DateLayout string) (time.Time, error) {
	var _err error
	var _time time.Time
	var _sTime string = strings.Trim(string(json), `"`)

	_time, _err = time.Parse(DateLayout, _sTime)
	return _time, _err
}

// =============================================================================
// =============================================================================
// =============================================================================
// =============================================================================
// =============================================================================
// this to json marshal and unmarshal date only in the form YYYY-MM-DD
// =============================================================================
type dateOnly time.Time

const dateOnlyLayout = "2006-01-02"

func (t *dateOnly) UnmarshalJSON(data []byte) error {
	_time, _err := CustomTimeUnMarshalJSON(data, dateOnlyLayout)
	*t = dateOnly(_time)
	return _err
}

func (t *dateOnly) MarshalJSON() ([]byte, error) {
	return CustomTimeMarshalJSON(time.Time(*t), dateOnlyLayout)
}

func (t *dateOnly) Set(date string) error {
	_time, _err := time.Parse(dateOnlyLayout, date)
	*t = dateOnly(_time)
	return _err
}

// =============================================================================
// =============================================================================

// ============================================================================
// this to json marshal and unmarshal SQLServer date
// which is in the form YYYY-MM-DDThh:mm:ss.sss. eg 2023-03-22T12:59:33.229
// ============================================================================
type sqlServerTimestamp time.Time

const SQLServerTimeStampLayout = "2006-01-02T15:04:05.000"

func (t *sqlServerTimestamp) UnmarshalJSON(data []byte) error {
	_time, _err := CustomTimeUnMarshalJSON(data, SQLServerTimeStampLayout)
	*t = sqlServerTimestamp(_time)
	return _err
}

func (t *sqlServerTimestamp) MarshalJSON() ([]byte, error) {
	return CustomTimeMarshalJSON(time.Time(*t), SQLServerTimeStampLayout)
}

func (t *sqlServerTimestamp) Set(date string) error {
	_time, _err := time.Parse(SQLServerTimeStampLayout, date)
	*t = sqlServerTimestamp(_time)

	return _err
}

// =============================================================================
// =============================================================================
// =============================================================================
// =============================================================================
// =============================================================================
// =============================================================================
