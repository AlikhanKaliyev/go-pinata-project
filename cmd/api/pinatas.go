package main

import (
	"PinataService.alikhankaliyev.net/internal/data"
	"fmt"
	"net/http"
)

func (app *application) createPinataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showPinataHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	pinata := data.Pinata{
		ID:       1,
		Color:    "multicolor",
		Shape:    "donkey",
		Contents: []string{"candy", "toys"},
		IsBroken: false,
		Weight:   2.5,
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
