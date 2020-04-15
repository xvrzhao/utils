package gorm

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Time 继承 time.Time 的所有方法，并重写 MarshalJSON 方法，
// Time 进行 json 序列化时，会转为字符串 "2006-01-02 15:04:05" 的格式。
type Time struct {
	time.Time
}

// MarshalJSON 将 Time 类型转为 "2006-01-02 15:04:05" 格式。
func (t Time) MarshalJSON() (data []byte, err error) {
	layout := "2006-01-02 15:04:05"
	data = make([]byte, 0, len(layout)+2)
	data = append(data, '"')
	data = t.AppendFormat(data, layout)
	data = append(data, '"')
	return
}

// Scan 在 gorm 进行读操作将数据库时间值转换为 Time 结构体字段时用到，
// 如果 Time 没有此方法，则会报错：无法将 time.Time 类型赋值到 Time 类型。
func (t *Time) Scan(v interface{}) error {
	if tt, ok := v.(time.Time); ok {
		*t = Time{tt}
		return nil
	}
	return fmt.Errorf("can not convert %v to Time", v)
}

// Value 在 gorm 进行写操作将 Time 结构体字段转为数据库时间值时用到，
// 如果 Time 没有此方法，则会报错：不支持 Time 类型。
func (t Time) Value() (driver.Value, error) {
	var zTime time.Time
	if t.UnixNano() == zTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}
