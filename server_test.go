package main

import (
	"log"
	"net/http/httptest"
	"strings"
	"testing"
)

// func TestHello(t *testing.T) {

// 	e := router()

// 	req := httptest.NewRequest("GET", "/", nil)
// 	rec := httptest.NewRecorder()

// 	e.ServeHTTP(rec, req)

// 	log.Print(req)
// }

func TestSaveConcentration(t *testing.T) {
	e := router()
	req := httptest.NewRequest("POST", "/save_concent", strings.NewReader(`{
		type: "nyan", id: 5fda34af565f987a66992587, measurement: "test", concentration_data: [{"kuso", "kuso"}]`))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	log.Print(req)
}

func TestGetID(t *testing.T) {
	e := router()
	req := httptest.NewRequest("GET", "/get_id", strings.NewReader(`{
		"type": "nyan", "measurement": "test", "concentration_data": []`))
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	log.Print(req)
}
