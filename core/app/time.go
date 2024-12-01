package app

import (
	"database/sql/driver"
	"fmt"
	"time"
)

const (
	timeFormart = "2006-01-02 15:04:05"
)

type DateTime time.Time

// MarshalJSON implements the json.Marshaler interface.
func (d *DateTime) MarshalJSON() ([]byte, error) {
	dt := time.Time(*d)
	return []byte(fmt.Sprintf(`"%v"`, dt.Format("2006-01-02 15:04:05"))), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *DateTime) UnmarshalJSON(b []byte) error {
	t, err := time.ParseInLocation(fmt.Sprintf(`"%s"`, timeFormart), string(b), time.Local)
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

// String implements the Stringer interface.
func (d DateTime) String() string {
	return time.Time(d).Format(timeFormart)
}

// Scan implements the Scanner interface.
// 读取数据库字段时处理
func (d *DateTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*d = DateTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// Value implements the driver Valuer interface.
// 写数据库时处理,零值写为NULL
func (d DateTime) Value() (driver.Value, error) {
	var zero time.Time
	t := time.Time(d)
	if t.UnixNano() == zero.UnixNano() {
		return nil, nil
	}
	return t, nil
}
