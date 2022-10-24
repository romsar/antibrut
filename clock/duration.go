package clock

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Duration struct {
	time.Duration
}

func (d Duration) Value() (driver.Value, error) {
	return driver.Value(int64(d.Seconds())), nil
}

func (d *Duration) Scan(raw any) error {
	switch v := raw.(type) {
	case int64:
		*d = Duration{
			Duration: time.Duration(v) * time.Second,
		}
	case nil:
		*d = Duration{
			Duration: time.Duration(0),
		}
	default:
		return errors.New("cannot sql.Scan() antibrut.Duration")
	}
	return nil
}
