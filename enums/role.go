package enums

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)

func (r *Role) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to scan role from the database driver: ", value))
	}
	*r = Role(bytes)
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}
