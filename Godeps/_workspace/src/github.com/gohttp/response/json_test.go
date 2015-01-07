package response

import "github.com/bmizerany/assert"
import "net/http/httptest"
import "testing"

type User struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

func TestJSONPretty(t *testing.T) {
	Pretty = true
	res := httptest.NewRecorder()
	JSON(res, &User{"Tobi", "Ferret"})
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "{\n  \"first\": \"Tobi\",\n  \"last\": \"Ferret\"\n}", string(res.Body.Bytes()))
	assert.Equal(t, "application/json", res.HeaderMap["Content-Type"][0])
}

func TestJSON(t *testing.T) {
	Pretty = false
	res := httptest.NewRecorder()
	JSON(res, &User{"Tobi", "Ferret"})
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, `{"first":"Tobi","last":"Ferret"}`, string(res.Body.Bytes()))
	assert.Equal(t, "application/json", res.HeaderMap["Content-Type"][0])
}
