package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Weight float64

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (w Weight) MarshalJSON() ([]byte, error) {

	jsonValue := fmt.Sprintf("%f kg", w)

	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}

func (w *Weight) UnmarshalJSON(jsonValue []byte) error {

	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "kg" {
		return ErrInvalidRuntimeFormat
	}
	// Otherwise, parse the string containing the number into an int32. Again, if this // fails return the ErrInvalidRuntimeFormat error.
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	// Convert the int32 to a Runtime type and assign this to the receiver. Note that we // use the * operator to deference the receiver (which is a pointer to a Runtime
	// type) in order to set the underlying value of the pointer.
	*w = Weight(i)
	return nil
}
