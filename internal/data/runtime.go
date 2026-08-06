package data

import (
	"fmt"
	"strconv"
)

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("%d mins", r)
	formatted = strconv.Quote(formatted)
	return []byte(formatted), nil
}
