package types

import (
	"database/sql/driver"
	"fmt"
	"strconv"
)

type Status uint8

func (s Status) Value() (driver.Value, error) {
	return int64(s), nil
}

func (s *Status) Scan(v any) error {
	s2 := asString(v)
	parseUint, err := strconv.ParseUint(s2, 10, 8)
	if err != nil {
		return fmt.Errorf("invalid database type: %T %v", v, v)
	}
	*s = Status(parseUint)
	return nil
}

const (
	StatusNormal Status = iota
	StatusBanned
)
