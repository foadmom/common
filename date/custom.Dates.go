package date

import (
	"fmt"
	"strings"
	"time"
)

// ============================================================================
// standard format for our timestamp is "2006-01-02T15:04:05.000"
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

// =============================================================================
// UnmarshalJSON implements the json.Unmarshaler interface for dateOnly type
// It uses the CustomTimeUnMarshalJSON function to parse the date string in the specified layout
// and assigns the parsed time to the dateOnly type.
// =============================================================================
func (t *dateOnly) UnmarshalJSON(data []byte) error {
	_time, _err := CustomTimeUnMarshalJSON(data, dateOnlyLayout)
	*t = dateOnly(_time)
	return _err
}

// =============================================================================
// MarshalJSON implements the json.Marshaler interface for dateOnly type
// It uses the CustomTimeMarshalJSON function to format the dateOnly
// type into a JSON string in the specified layout.
// =============================================================================
func (t *dateOnly) MarshalJSON() ([]byte, error) {
	return CustomTimeMarshalJSON(time.Time(*t), dateOnlyLayout)
}

// =============================================================================
// Set is a method for dateOnly type that takes a date string in the
// specified layout and parses it into a time.Time value.
// It uses the time.Parse function to parse the date string and assigns
// the parsed time to the dateOnly type. It returns an error if the parsing fails.
// =============================================================================
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

// =============================================================================
// UnmarshalJSON implements the json.Unmarshaler interface for sqlServerTimestamp type
// It uses the CustomTimeUnMarshalJSON function to parse the date string in the specified layout
// and assigns the parsed time to the sqlServerTimestamp type.
// =============================================================================
func (t *sqlServerTimestamp) UnmarshalJSON(data []byte) error {
	_time, _err := CustomTimeUnMarshalJSON(data, SQLServerTimeStampLayout)
	*t = sqlServerTimestamp(_time)
	return _err
}

// =============================================================================
// MarshalJSON implements the json.Marshaler interface for sqlServerTimestamp type
// It uses the CustomTimeMarshalJSON function to format the sqlServerTimestamp
// type into a JSON string in the specified layout.
// =============================================================================
func (t *sqlServerTimestamp) MarshalJSON() ([]byte, error) {
	return CustomTimeMarshalJSON(time.Time(*t), SQLServerTimeStampLayout)
}

// =============================================================================
// Set is a method for sqlServerTimestamp type that takes a date string in the
// specified layout and parses it into a time.Time value.
// It uses the time.Parse function to parse the date string and assigns
// the parsed time to the sqlServerTimestamp type. It returns an error if the parsing fails.
// =============================================================================
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
