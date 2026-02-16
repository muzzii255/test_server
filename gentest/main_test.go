package gentests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

var testApp *your.App // your app

func TestMain(m *testing.M) {
	// configure your app here
}

func setup() *your.App {
	return testApp
}

func testHelloWorld(t *testing.T) {
	app := setup()
	payload := map[string]string{"name": "Deeznuts"}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/test", bytes.NewReader(body))
	req.Header.Set("content-type", "application/json")

	resp, err := app.Test(req)
	require.NoError(t, err)

	require.Equal(t, http.StatusOK, resp.StatusCode)

	var responseMap map[string]string
	json.NewDecoder(resp.Body).Decode(&responseMap)
	require.Equal(t, "hello Deeznuts", responseMap["result"])
}

// func TestSequentially(t *testing.T) {
// 	t.Run("HelloWorld", testHelloWorld)
// 	// Run your tests here
// }
