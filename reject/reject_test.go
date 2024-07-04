package reject

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

var logger = slog.New(slog.NewTextHandler(io.Discard, nil))

var ls = LoggerSupport{
	Logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
}

func checkError(t *testing.T, errors map[string]string, key, value string) {
	t.Helper()

	if errors[key] != value {
		t.Errorf("got %v want %v", errors[key], value)
	}
}

func TestBasicRejections(t *testing.T) {
	var cases = []struct {
		rejectFn        func(w http.ResponseWriter, r *http.Request)
		expectedStatus  int
		expectedMessage string
	}{
		{ls.NotFound, http.StatusNotFound, "not found"},
		{ls.BadRequest, http.StatusBadRequest, "bad request"},
		{ls.InternalServerError, http.StatusInternalServerError, "internal server error"},
		{ls.Forbidden, http.StatusForbidden, "forbidden"},
		{ls.Unauthorized, http.StatusUnauthorized, "unauthorized"},
	}

	for _, c := range cases {
		r := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		c.rejectFn(w, r)

		resp := w.Result()

		if resp.StatusCode != c.expectedStatus {
			t.Errorf("failed case: %v: got %v want %v", resp.StatusCode, resp.StatusCode, c.expectedStatus)
		}

		var body map[string]string

		err := json.NewDecoder(resp.Body).Decode(&body)
		if err != nil {
			t.Fatal(err)
		}

		checkError(t, body, "error", c.expectedMessage)
	}
}

func TestValidationRejection(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	errors := map[string]string{"field": "must not be empty", "another": "must be greater then 10"}

	ls.ValidationError(w, r, errors)

	resp := w.Result()

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("got %v want %v", resp.StatusCode, http.StatusBadRequest)
	}

	var body map[string]map[string]string

	err := json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		t.Fatal(err)
	}

	e := body["error"]
	if len(e) != 2 {
		t.Errorf("got %v want %v", len(e), 2)
	}

	checkError(t, e, "field", "must not be empty")
	checkError(t, e, "another", "must be greater then 10")
}
