package spec_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/swaggest/assertjson"
	"github.com/swaggest/go-asyncapi/spec-2.0.0"
)

func TestInfo_MarshalJSON(t *testing.T) {
	i := spec.Info{
		Title:   "foo",
		Version: "v1",
		MapOfAnything: map[string]interface{}{
			"x-two": "two",
			"x-one": 1,
		},
	}

	res, err := json.Marshal(i)
	assert.NoError(t, err)
	assert.Equal(t, `{"title":"foo","version":"v1","x-one":1,"x-two":"two"}`, string(res))
}

func TestInfo_MarshalJSON_Nil(t *testing.T) {
	i := spec.Info{
		Title:   "foo",
		Version: "v1",
	}

	res, err := json.Marshal(i)
	assert.NoError(t, err)
	assert.Equal(t, `{"title":"foo","version":"v1"}`, string(res))
}

func TestInfo_UnmarshalJSON(t *testing.T) {
	i := spec.Info{
		Title: "foo",
	}

	err := json.Unmarshal([]byte(`{"title":"foo","version":"v1","x-one":1,"x-two":"two"}`), &i)
	assert.NoError(t, err)
	assert.Equal(t, 1.0, i.MapOfAnything["x-one"].(float64))
	assert.Equal(t, "two", i.MapOfAnything["x-two"])
}

func TestAsyncAPI_UnmarshalJSON_roundTrip(t *testing.T) {
	data, err := os.ReadFile("../resources/fixtures/streetlights-2.0.0.json")
	require.NoError(t, err)

	var a spec.AsyncAPI
	err = json.Unmarshal(data, &a)
	require.NoError(t, err)

	roundTripped, err := json.Marshal(a)
	require.NoError(t, err)

	assertjson.Equal(t, data, roundTripped)
}

func TestAsyncAPI_UnmarshalYAML(t *testing.T) {
	data, err := os.ReadFile("../resources/fixtures/streetlights-2.0.0.yml")
	require.NoError(t, err)

	var a spec.AsyncAPI
	err = a.UnmarshalYAML(data)
	require.NoError(t, err)

	assert.Equal(t, "#/components/messages/lightMeasured", a.Channels["smartylighting/streetlights/1/0/event/{streetlightId}/lighting/measured"].Publish.Message.Reference.Ref)

	marshaledData, err := a.MarshalYAML()
	require.NoError(t, err)

	assert.Equal(t, string(data), string(marshaledData))
}
