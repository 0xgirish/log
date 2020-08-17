package log

import (
	"testing"
)

func TestJsonParse(t *testing.T) {
	s := `{"time":"2019-07-10T05:35:54.277Z","level":"info","caller":"pretty.go:42","error":"i am test error","foo":"bar","n":42,"t":true,"f":false,"o":null,"a":[1,2,3],"obj":{"a":[1,2], "b":{"c":3}},"message":"hello json console color writer"}`

	results, err := parseJson(s)
	if err != nil {
		t.Errorf("parseJson error: %+v", err)
	}

	for _, v := range results {
		t.Logf("%#v\n", v)
	}
}
