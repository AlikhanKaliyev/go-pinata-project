package data

import "time"

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
