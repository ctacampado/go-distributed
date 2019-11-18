package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	url := "http://localhost:8080"

	exp := "Welcome!"

	req, _ := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("%+v", err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	resp := string(body)
	if resp == exp {
		fmt.Printf("exp: %s | res: %s\n", exp, resp)
	}
}

func TestHandleRegisterPassThrough(t *testing.T) {

	url := "http://localhost:8080/register"

	in := Credentials{Username: "user", Password: "password"}
	exp := Credentials{Username: "user", Password: "password"}

	b, _ := json.Marshal(in)
	payload := strings.NewReader(string(b))

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		t.Errorf("%+v", err)
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("%+v", err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	resp := Credentials{}
	json.Unmarshal(body, &resp)
	if reflect.DeepEqual(resp, exp) {
		fmt.Printf("exp: %+v | res: %+v\n", exp, resp)
	}
}
