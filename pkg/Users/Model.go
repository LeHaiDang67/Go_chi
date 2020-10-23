package Users

import "time"

type User struct {
	ID       string
	Address  string
	Birthday time.Time
	Name     string
}
