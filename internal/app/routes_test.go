package app_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"super-heroes/internal/app"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHealth(t *testing.T) {
	t.Parallel()

	sut := app.New(app.Config{})
	srv := httptest.NewServer(sut.Routes())

	defer srv.Close()

	resp, err := srv.Client().Get(srv.URL + "/health")
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, "OK", string(body))

	requestID := resp.Header.Get("X-Request-ID")
	assert.NotEmpty(t, requestID)
}

func TestHealthWithARequestID(t *testing.T) {
	t.Parallel()

	sut := app.New(app.Config{})
	srv := httptest.NewServer(sut.Routes())

	defer srv.Close()

	client := srv.Client()

	req, err := http.NewRequest("GET", srv.URL+"/health", nil)
	require.NoError(t, err)

	const requestID = "test-1234"

	req.Header.Set("X-Request-ID", requestID)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)
	assert.Equal(t, "OK", string(body))

	actualRequestID := resp.Header.Get("X-Request-ID")
	assert.Equal(t, requestID, actualRequestID)
}
