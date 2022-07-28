package jsonpath

import (
	"encoding/json"
	"testing"
)

func TestGetPathValueKeyWithSpaces(t *testing.T) {
	jp := MustParsePath(`$["key with spaces"]`)
	data := map[string]interface{}{
		"key with spaces": "v",
	}
	value, err := GetPathValue(data, jp)
	if err != nil {
		t.Error(err)
		return
	}
	stringVal, ok := value.(string)
	if !ok {
		t.Error("value was not string")
		return
	}
	if stringVal != "v" {
		t.Error("value is incorrect")
		return
	}
}

func TestJsonUnmarshal(t *testing.T) {
	jsonEncodedExpression := []byte(`"$[\"key with spaces\"]"`)
	var target JsonPath
	err := json.Unmarshal(jsonEncodedExpression, &target)
	if err != nil {
		t.Error(err)
		return
	}
	if target.str != `$["key with spaces"]` {
		t.Error("parsing error")
		return
	}
}

func TestJsonMarshal(t *testing.T) {
	jp := MustParsePath(`$["key with spaces"]`)
	b, err := json.Marshal(jp)
	if err != nil {
		t.Error(err)
		return
	}
	if string(b) != `"$[\"key with spaces\"]"` {
		t.Error("value is incorrect")
		return
	}
}
