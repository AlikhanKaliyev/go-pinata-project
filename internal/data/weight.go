package data

import (
	"fmt"
	"strconv"
)

type Weight float64

func (w Weight) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%f kg", w)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}
