package handlers

import (
    "testing"
    "net/http/httptest"
    "net/http"
    "github.com/labstack/echo"
)

func TestHealthCheck(t *testing.T) {
    e := echo.New()
    e.GET("/", HealthCheck)

    rec := httptest.NewRecorder()
    req := httptest.NewRequest("GET", "/", nil)

    e.ServeHTTP(rec, req)
    resp := rec.Result()
    if resp.StatusCode != http.StatusOK {
        t.Error("unexpected status code: ", resp.Status)
    }
}
