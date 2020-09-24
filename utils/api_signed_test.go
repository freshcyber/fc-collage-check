package utils

import (
	"net/url"
	"strings"

	"testing"
)

func TestBindJSON(t *testing.T) {
	_reader := strings.NewReader(string("{\"id\":\"user01\",\"name\":\"eric\",\"key\":\"092982528307ED4B4C7654DF0790E4B2\"}"))
	var props map[string]interface{}
	BindJSON(_reader, &props)
	t.Logf("%v\n", props)

	if props["id"] != "user01" {
		t.Errorf("id error")
	}
}

func TestAPICalcSign(t *testing.T) {
	_reader := strings.NewReader(string("{\"id\":123,\"name\":\"\",\"api_token\":\"499456AFE7368947C23C5AA8EF64F939\"}"))
	var props map[string]interface{}
	BindJSON(_reader, &props)
	t.Logf("%v\n", props)

	// VeZ16GuXyKaYhKbA?id=123&name=
	_sign := APICalcSign(props, "VeZ16GuXyKaYhKbA?", "api_token")

	t.Logf("%s\n", _sign)

	t.Logf("%s\n", props["api_token"])

	if props["api_token"] != _sign {
		t.Errorf("signed error")
	}
}

func TestFormAPICalcSign(t *testing.T) {
	query := "id=123&name=eric&api_token=02C32E5F1EB9BB0707F3B185719C42E5&a="

	_maps, _ := url.ParseQuery(query)
	_props := make(map[string]interface{})
	for _key, _value := range _maps {
		_props[_key] = _value[0]
	}

	t.Logf("%v\n", _props)

	// VeZ16GuXyKaYhKbA?a=&id=123&name=eric
	_sign := APICalcSign(_props, "VeZ16GuXyKaYhKbA?", "api_token")

	t.Logf("%s\n", _sign)

	t.Logf("%s\n", _props["api_token"])

	if _props["api_token"] != _sign {
		t.Errorf("signed error")
	}
}

func TestFormAPICalcSignNil(t *testing.T) {
	query := "api_token=53F29A0E5A243DD78639D2DC7120A8AF&a="

	_maps, _ := url.ParseQuery(query)
	_props := make(map[string]interface{})
	for _key, _value := range _maps {
		_props[_key] = _value[0]
	}
	t.Logf("%v\n", _props)

	// VeZ16GuXyKaYhKbA?a=&id=123&name=eric
	_sign := APICalcSign(_props, "VeZ16GuXyKaYhKbA?", "api_token")

	t.Logf("%s\n", _sign)

	t.Logf("%s\n", _props["api_token"])

	if _props["api_token"] != _sign {
		t.Errorf("signed error")
	}
}
