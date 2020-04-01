package jsonpath

/*
We wrap "github.com/oliveagle/jsonpath" so that we can implement some json interface methods on top of the base type.
*/

// stdlib
import (
	// "bytes"
	// "io"
	"strings"
	// "encoding/json"
)

// external
import (
	"github.com/oliveagle/jsonpath"
)

func GetPathValue(data interface{}, jsonPath JsonPath) (interface{}, error) {
	return jsonPath.Path.Lookup(data)
}

func MustParsePath(path string) JsonPath {
	compiled := jsonpath.MustCompile(path)
	return JsonPath{compiled, path}
}

type JsonPath struct {
	Path *jsonpath.Compiled
	str   string
}

func (jp *JsonPath) UnmarshalJSON(path []byte) error {
	// primary
	jp.str = strings.Trim(string(path), `"`)
	compiled, err := jsonpath.Compile(jp.str)
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
