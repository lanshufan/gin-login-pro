package modules

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// XTime time.Time
type XTime time.Time

// MarshalJSON 重写time.Time的方法，因为它默认时间格式化RFC 3339格式
func (t *XTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)

	if t == nil || time.Time(*t).IsZero() {
		return []byte("null"), nil
	}

	s := fmt.Sprintf("\"%s\"", tTime.Format("2006-01-02 15:04:05"))
	return []byte(s), nil
}

func (t *XTime) UnmarshalJSON(data []byte) error {
	if string(data) == "" {
		return nil
	}
	str := string(data)
	timeStr := strings.Trim(str, "\"")
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, time.Local)
	*t = XTime(t1)

	return err
}

// Scan 从数据库里读取的数据value，转换成自定义的格式
func (t *XTime) Scan(value interface{}) error {
	timeVal, ok := value.(time.Time)
	if !ok {
		return errors.New("无法解析时间")
	}

	*t = XTime(timeVal)

	return nil
}

// Value 写入数据库前的日期时间格式转换
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time

	if time.Time(t).UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}

	return time.Time(t), nil
}
