package reject

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

// NewLoggerSupport creates a new instance of the LoggerSupport struct.
type LoggerSupport struct {
	Logger         *slog.Logger // Logger instance
	RequestEnabled bool         // Enable request logging
}

type envelope map[string]any

// Rejects the request with HTTP status code 404 Not Found and a JSON response body.
func (app *LoggerSupport) NotFound(w http.ResponseWriter, r *http.Request) {
	data := envelope{"error": "not found"}
	if app.RequestEnabled {
		app.logError(r, errors.New("not found"))
	}
	app.rejectRequest(w, r, http.StatusNotFound, data)
}

// Rejects the request with HTTP status code 400 Bad Request and a JSON response body containing a validation error message.
func (app *LoggerSupport) ValidationError(w http.ResponseWriter, r *http.Request, errorMessages map[string]string) {
	data := envelope{"error": errorMessages}
	if app.RequestEnabled {
		app.logError(r, errors.New("validation error"))
	}
	app.rejectRequest(w, r, http.StatusBadRequest, data)
}

// Rejects the request with HTTP status code 400 Bad Request and a JSON response body.
func (app *LoggerSupport) BadRequest(w http.ResponseWriter, r *http.Request) {
	data := envelope{"error": "bad request"}
	if app.RequestEnabled {
		app.logError(r, errors.New("bad request"))
	}
	app.rejectRequest(w, r, http.StatusBadRequest, data)
}

// Rejects the request with HTTP status code 500 Internal Server Error and a JSON response body.
func (app *LoggerSupport) InternalServerError(w http.ResponseWriter, r *http.Request) {
	data := envelope{"error": "internal server error"}
	if app.RequestEnabled {
		app.logError(r, errors.New("internal server error"))
	}
	app.rejectRequest(w, r, http.StatusInternalServerError, data)
}

// Rejects the request with HTTP status code 401 Unauthorized and a JSON response body.
func (app *LoggerSupport) Unauthorized(w http.ResponseWriter, r *http.Request) {
	data := envelope{"error": "unauthorized"}
	if app.RequestEnabled {
		app.logError(r, errors.New("unauthorized"))
	}
	app.rejectRequest(w, r, http.StatusUnauthorized, data)
}

// Rejects the request with HTTP status code 403 Forbidden and a JSON response body.
func (app *LoggerSupport) Forbidden(w http.ResponseWriter, r *http.Request) {
	data := envelope{"error": "forbidden"}
	if app.RequestEnabled {
		app.logError(r, errors.New("forbidden"))
	}
	app.rejectRequest(w, r, http.StatusForbidden, data)
}

func (app *LoggerSupport) rejectRequest(w http.ResponseWriter, r *http.Request, status int, data envelope) {
	js, err := json.Marshal(data)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func (app *LoggerSupport) logError(r *http.Request, err error) {
	var (
		url    = r.URL.RequestURI()
		method = r.Method
	)

	err = fmt.Errorf("reject: %w", err)
	app.Logger.Error(err.Error(), "method", method, "url", url)
}
