package mango

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/a/b", nil)
	r.addRoute("GET", "/assets/*filepath", nil)

	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	if !ok {
		t.Fatal("test parse pattern failed. /p/:name")
	}
	ok = reflect.DeepEqual(parsePattern("/assets/*"), []string{"assets", "*"})
	if !ok {
		t.Fatal("test parse pattern failed. /assets/*")
	}
	ok = reflect.DeepEqual(parsePattern("/abc/*ddd/kjd"), []string{"abc", "*ddd"})
	if !ok {
		t.Fatal("test parse pattern failed. /abc/*ddd/kjd")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()

	n, ps := r.getRoute("GET", "/hello/abc")

	if n == nil {
		t.Fatal("nil shouldn't be returned.")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("pattern should be match /hello/:name; now is: ", n.pattern)
	}

	if ps["name"] != "abc" {
		t.Fatal("param name should be abc", ps["name"])
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])

}
