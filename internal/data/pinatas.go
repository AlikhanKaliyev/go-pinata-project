package data

import (
	"PinataService.alikhankaliyev.net/internal/validator"
	"strings"
	"time"
)

type Pinata struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"-"`
	Color      string    `json:"color,omitempty"`
	Shape      string    `json:"shape,omitempty"`
	Contents   []string  `json:"contents,omitempty"`
	IsBroken   bool      `json:"broken"`
	Weight     Weight    `json:"weight,omitempty,string"`
	Dimensions struct {
		Height float64 `json:"height,string"`
		Width  float64 `json:"width,string"`
		Depth  float64 `json:"depth,string"`
	} `json:"dimensions,omitempty"`
}

func ValidateMovie(v *validator.Validator, pinata *Pinata) {
	v.Check(pinata.Color != "", "color", "must be provided")
	v.Check(len(strings.Split(pinata.Color, " ")) > 1, "color", "there must be one word")

	v.Check(pinata.Contents != nil, "contents", "must be provided")
	v.Check(len(pinata.Contents) >= 1, "contents", "must contain at least 1 content")
	v.Check(len(pinata.Contents) <= 100, "contents", "must not contain more than 100 contents")

	v.Check(validator.Unique(pinata.Contents), "contents", "must not contain duplicate values")
}
