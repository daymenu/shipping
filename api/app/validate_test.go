package app

import (
	"net/url"
	"testing"
)

func TestCheckEmail(t *testing.T) {
	urlReq, _ := url.Parse("http://phpcap.com/hello/kit?name=hj")
	t.Error("url path", urlReq.Path)
	email := "hanjian2018@126.com"
	if err := Check("email", email); err != nil {
		t.Fail()
	}
}

func TestCheckUserName(t *testing.T) {
	username := "hanjian2018"
	if err := Check("username", username); err != nil {
		t.Fail()
	}
}
func TestCheckNotEmpty(t *testing.T) {
	email := "wd"
	if err := Check("notempty", email); err != nil {
		t.Fatal(err)
	}
	empty := ""

	if err := Check("notempty", empty); err == nil {
		t.Fatal(err)
	}
}
