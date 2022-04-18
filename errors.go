package ultra_table

import (
	"errors"
)

var (
	RecordNotFound = errors.New("record not found")
	OnlySupportPtr = errors.New("only support ptr")
	ValueNotNil    = errors.New("value can not be nil")
	UniqueIndex    = errors.New("unique index duplicate")
	
)
