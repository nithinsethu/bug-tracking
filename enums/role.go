package enums

import (
	"database/sql/driver"
	"fmt"
)

type Role string

const (
	Admin  Role = "admin"
	Member Role = "member"
)

func (r *Role) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan role from the database driver: %v", value)
	}
	*r = Role(v)
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}
