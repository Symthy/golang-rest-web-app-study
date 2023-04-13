package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMux(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/health", nil)
	sut := newMux()
	sut.ServeHTTP(w, r)

	resp := w.Result()
	t.Cleanup(func() {
		_ = resp.Body.Close()
	})

	require.Equal(t, resp.StatusCode, http.StatusOK)
	got, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	want := `{"status": "ok"}`
	require.Equal(t, want, string(got))
}
