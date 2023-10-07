package main

import (
	"PinataService.alikhankaliyev.net/internal/data"
	"PinataService.alikhankaliyev.net/internal/validator"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createPinataHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Color      string      `json:"color"`
		Shape      string      `json:"shape"`
		Contents   []string    `json:"contents"`
		IsBroken   bool        `json:"broken"`
		Weight     data.Weight `json:"weight"`
		Dimensions struct {
			Height float64 `json:"height,string"`
			Width  float64 `json:"width,string"`
			Depth  float64 `json:"depth,string"`
		} `json:"dimensions"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	pinata := &data.Pinata{
		Color:      input.Color,
		Shape:      input.Shape,
		Contents:   input.Contents,
		IsBroken:   input.IsBroken,
		Weight:     input.Weight,
		Dimensions: input.Dimensions,
	}

	v := validator.New()

	if data.ValidateMovie(v, pinata); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showPinataHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	pinata := data.Pinata{
		ID:        1,
		CreatedAt: time.Now(),
		Color:     "multicolor",
		Shape:     "donkey",
		Contents:  []string{"candy", "toys"},
		IsBroken:  false,
		Weight:    2.5,
		Dimensions: struct {
			Height float64 `json:"height,string"`
			Width  float64 `json:"width,string"`
			Depth  float64 `json:"depth,string"`
		}{Height: 50.0, Width: 30.0, Depth: 15.0},
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"pinata": pinata}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
