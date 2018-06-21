package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchHandler(t *testing.T) {
	t.Log("should get the 200 OK status")

	ts := httptest.NewServer(http.HandlerFunc(searchHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL + "/search/")
	if err != nil {
		t.Error("unexpected response")
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	t.Log(string(body))

	if res.StatusCode != 200 {
		t.Error("Incorrect status code:", res.StatusCode)
	}
}
