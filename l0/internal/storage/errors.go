package storage

import "fmt"

var ErrUniqueViolation = fmt.Errorf("unique violation")
var ErrNotFound = fmt.Errorf("not found")
