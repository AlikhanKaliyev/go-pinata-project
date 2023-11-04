package main

import (
	"PinataService.alikhankaliyev.net/internal/data"
	"PinataService.alikhankaliyev.net/internal/validator"
	"errors"
	"net/http"
)

func (app *application) createPinataHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Color      string      `json:"color"`
		Shape      string      `json:"shape"`
		Contents   []string    `json:"contents"`
		IsBroken   bool        `json:"broken"`
		Weight     data.Weight `json:"weight"`
		Dimensions struct {
			Height float32 `json:"height,string"`
			Width  float32 `json:"width,string"`
			Depth  float32 `json:"depth,string"`
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

	if data.ValidatePinata(v, pinata); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	//fmt.Fprintf(w, "%+v\n", input)

	err = app.models.Pinatas.Insert(pinata)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	//headers.Set("Location", fmt.Sprintf("/v1/pinatas/%d", pinata.ID))
	// Write a JSON response with a 201 Created status code, the movie data in the // response body, and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"pinata": pinata}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showPinataHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	pinata, err := app.models.Pinatas.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"pinata": pinata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updatePinataHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	pinata, err := app.models.Pinatas.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Color      string      `json:"color"`
		Shape      string      `json:"shape"`
		Contents   []string    `json:"contents"`
		IsBroken   bool        `json:"broken"`
		Weight     data.Weight `json:"weight"`
		Dimensions struct {
			Height float32 `json:"height,string"`
			Width  float32 `json:"width,string"`
			Depth  float32 `json:"depth,string"`
		} `json:"dimensions"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	pinata = &data.Pinata{
		ID:         id,
		Color:      input.Color,
		Shape:      input.Shape,
		Contents:   input.Contents,
		IsBroken:   input.IsBroken,
		Weight:     input.Weight,
		Dimensions: input.Dimensions,
	}

	v := validator.New()
	if data.ValidatePinata(v, pinata); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Pinatas.Update(pinata)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"pinata": pinata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deletePinataHandler(w http.ResponseWriter, r *http.Request) { // Extract the movie ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Delete the movie from the database, sending a 404 Not Found response to the // client if there isn't a matching record.
	err = app.models.Pinatas.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Return a 200 OK status code along with a success message.
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "pinata successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
