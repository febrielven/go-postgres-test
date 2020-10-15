// +build unit

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AddResult struct {
	x        int
	y        int
	expected int
}

var addResult = []AddResult{
	{1, 2, 3},
	{5, 5, 10},
	{2, 2, 4},
}

func TestAdd(t *testing.T) {
	for _, test := range addResult {
		res := Add(test.x, test.y)
		if res != test.expected {
			t.Fatal("Expected result not given")
		}
	}
}

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/test.data")
	if err != nil {
		t.Fatal("Could not open file")
	}
	if string(data) != "hello" {
		t.Fatalf("string contents do not match expected")
	}
}

func TestHttpRequest(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "{\"status\": \"goods\"}")
	}

	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	if 200 != resp.StatusCode {
		t.Fatal("status code not oke")
	}
}
