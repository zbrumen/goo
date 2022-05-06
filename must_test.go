package goo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDo(t *testing.T) {
	var out string
	err := Do(http.Get("https://google.com")).Then(func(r *http.Response) error {
		defer r.Body.Close()
		return Do(ioutil.ReadAll(r.Body)).Then(func(d []byte) error {
			out = string(d)
			return nil
		})
	})
	fmt.Println(out, err)
}

func TestCheck(t *testing.T) {
	data := map[string]any{
		"foo": "bar",
	}
	fmt.Println(Check(Get(data, "foo")).Then(func(data any) error {
		fmt.Println(data)
		return nil
	}))
}
