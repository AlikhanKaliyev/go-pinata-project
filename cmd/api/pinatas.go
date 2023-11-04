package main

import (
	"PinataService.alikhankaliyev.net/internal/data"
	"PinataService.alikhankaliyev.net/internal/validator"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) createPinataHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Color      string      `json:"color"`
		Shape      string      `json:"shape"`
		Contents   []string    `json:"contents"`
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
	fmt.Print(pinata)
	if input.Color != "" {
		pinata.Color = input.Color
	}
	if input.Shape != "" {
		pinata.Shape = input.Shape
	}
	if input.Contents != nil {
		pinata.Contents = input.Contents
	}
	if input.Weight != 0 {
		pinata.Weight = input.Weight
	}
	if &input.Dimensions != nil {
		if input.Dimensions.Height != 0 {
			pinata.Dimensions.Height = input.Dimensions.Height
		}
		if input.Dimensions.Width != 0 {
			pinata.Dimensions.Width = input.Dimensions.Width
		}
		if input.Dimensions.Depth != 0 {
			pinata.Dimensions.Depth = input.Dimensions.Depth
		}
	}

	v := validator.New()
	if data.ValidatePinata(v, pinata); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Pinatas.Update(pinata)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
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

func (app *application) listPinatasHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Color    string
		Contents []string
		data.Filters
	}
	v := validator.New()
	qs := r.URL.Query()

	input.Color = app.readString(qs, "color", "")
	input.Contents = app.readCSV(qs, "contents", []string{})

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	// Read the sort query string value into the embedded struct.
	input.Filters.Sort = app.readString(qs, "sort", "id")

	input.Filters.SortSafelist = []string{"id", "-id", "weight", "-weight"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	pinatas, metadata, err := app.models.Pinatas.GetAll(input.Color, input.Contents, input.Filters)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"pinatas": pinatas, "metadata": metadata}, nil)

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
