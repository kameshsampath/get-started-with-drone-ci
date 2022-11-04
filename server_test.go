package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAddPost(t *testing.T) {
	postJSON := `{"name":"kamesh","text":"get started with drone and golang"}`
	want := `{"id":1,"name":"kamesh","text":"get started with drone and golang"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/add", strings.NewReader(postJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, addPost(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		got := strings.TrimSpace(rec.Body.String())
		assert.Equalf(t, want, got, "want %s != got %s", want, got)
	}
}
