package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	server *httptest.Server
	usrUrl string
)

func init() {
	server = httptest.NewServer(Handlers())
}

func TestGetVersion(t *testing.T) {
	usrUrl = fmt.Sprintf("%s/service/version", server.URL)
	checkResponceCode(t, http.StatusOK, getResponse(t, usrUrl).StatusCode)
}

func TestNotFoundInterface(t *testing.T) {
	usrUrl = fmt.Sprintf("%s/service/%s/interfaces/adsdadass", server.URL, V)
	checkResponceCode(t, http.StatusNotFound, getResponse(t, usrUrl).StatusCode)
}

func TestDetailInterface(t *testing.T) {
	inf, err := net.Interfaces()
	if err != nil {
		t.Error(err)
	}
	for _, i := range inf {
		usrUrl = fmt.Sprintf("%s/service/v1/interfaces/%s", server.URL, i.Name)
		checkResponceCode(t, http.StatusOK, getResponse(t, usrUrl).StatusCode)
	}
}

func TestInternalServerErrorWithWrongVersion(t *testing.T) {
	usrUrl = fmt.Sprintf("%s/service/v100500/interfaces/adsdadass", server.URL)
	checkResponceCode(t, http.StatusInternalServerError, getResponse(t, usrUrl).StatusCode)
}

func TestInternalServerErrorDefault(t *testing.T) {
	usrUrl = fmt.Sprintf("%s/service", server.URL)
	checkResponceCode(t, http.StatusInternalServerError, getResponse(t, usrUrl).StatusCode)
}

func TestGetInterfaces(t *testing.T) {
	usrUrl = fmt.Sprintf("%s/service/%s/interfaces", server.URL, V)
	checkResponceCode(t, http.StatusOK, getResponse(t, usrUrl).StatusCode)
}

func getResponse(t *testing.T, url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	return res
}

func checkResponceCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected responce code %d. Got %d\n", expected, actual)
	}
}
