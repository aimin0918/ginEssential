package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const Day = time.Hour * 24
const TimeFormat = "2006-01-02 15:04:05"

type XTime struct {
	time.Time
}

func (t *XTime) UnmarshalJSON(data []byte) (err error) {
	t.Time, err = time.ParseInLocation(TimeFormat, string(data[1:len(data)-1]), time.Local)
	return
}

func (t *XTime) MarshalJSON() ([]byte, error) {
	return []byte(t.Time.Format("\"2006-01-02 15:04:05\"")), nil
}

func (t *XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t *XTime) Scan(value interface{}) error {
	v, ok := value.(time.Time)
	if ok {
		t.Time = v
		return nil
	}

	return fmt.Errorf("can not convert %v to time.Time", value)
}

func (t *XTime) String() string {
	return t.Time.Format(TimeFormat)
}

func (t *XTime) Unix() int64 {
	loc := time.Local
	kTime, _ := time.ParseInLocation(TimeFormat, t.String(), loc)
	return kTime.Unix()
}

func (t *XTime) UnixNano() int64 {
	loc := time.Local
	kTime, _ := time.ParseInLocation(TimeFormat, t.String(), loc)
	return kTime.UnixNano()
}
