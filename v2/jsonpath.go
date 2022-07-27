package jsonpath

/*
We wrap "github.com/PaesslerAG/jsonpath" so that we can implement some json interface methods on top of the base type.
*/

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/PaesslerAG/gval"
	"github.com/PaesslerAG/jsonpath"
)

func GetPathValue(data interface{}, jsonPath JsonPath) (interface{}, error) {
	return jsonPath.Path(context.Background(), data)
}

// MustParsePath passes the jsonpath expression or panics on error
func MustParsePath(path string) JsonPath {
	compiled, err := jsonpath.New(path)
	if err != nil {
		panic(err)
	}
	return JsonPath{compiled, path}
}

type JsonPath struct {
	Path gval.Evaluable
	str  string
}

func (jp *JsonPath) UnmarshalJSON(path []byte) error {
	// primary
	err := json.Unmarshal(path, &jp.str)
	if err != nil {
		return err
	}
	compiled, err := jsonpath.New(jp.str)
	if err != nil {
		return err
	}
	jp.Path = compiled

	return nil
}

func (jp *JsonPath) MarshalJSON() ([]byte, error) {
	return []byte(`"` + strings.TrimRight(jp.String(), `"`) + `"`), nil
}

func (jp *JsonPath) String() string {
	return jp.str
}
