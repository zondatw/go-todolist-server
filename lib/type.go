package lib

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthFunc func(*gin.Context) (interface{}, error)

type Date time.Time

func (date Date) Value() (driver.Value, error) {
	return []byte(time.Time(date).Format(`"2006-01-02"`)), nil
}

func (date *Date) Scan(v interface{}) error {
	switch s := v.(type) {
	case time.Time:
		*date = Date(s)
	case []byte:
		newTime, err := time.Parse("2006-01-02", string(s))
		if err != nil {
			return err
		}
		*date = Date(newTime)
	case string:
		newTime, err := time.Parse("2006-01-02", s)
		if err != nil {
			return err
		}
		*date = Date(newTime)
	default:
		return fmt.Errorf("date: Unsupport scanning type %T", v)
	}
	return nil
}

func (date *Date) UnmarshalJSON(input []byte) error {
	newTime, err := time.Parse(`"2006-01-02"`, string(input))
	if err != nil {
		return err
	}

	*date = Date(newTime)
	return nil
}

func (date Date) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(date).Format(`"2006-01-02"`)), nil
}
