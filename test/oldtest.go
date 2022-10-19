//package test
//
//import (
//	"demo/repo"
//	"log"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//package main
//
//import (
//"github.com/go-playground/assert/v2"
//"log"
//"net/http"
//"net/http/httptest"
//"testing"
//)
//
//func TestEcommerceApi_GETCollezioni(t *testing.T) {
//	router := repo.New()
//	w := httptest.NewRecorder()
//	req, err := http.NewRequest("GET", "/collezione/", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	router.ServeHTTP(w, req)
//	assert.Equal(t, 200, w.Code)
//	collezioneString := "[{\"Id\":\"2\",\"Name\":\"Estate\",\"Articles\":null},{\"Id\":\"1\",\"Name\":\"Inverno\",\"Articles\":null}]"
//	assert.Equal(t, collezioneString, w.Body.String())
//}
//
//func TestEcommerceApi_GETCollezione(t *testing.T) {
//	router := repo.New()
//	w := httptest.NewRecorder()
//	req, err := http.NewRequest("GET", "/collezione/1", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	router.ServeHTTP(w, req)
//	assert.Equal(t, 200, w.Code)
//	collezioneString := "[{\"Id\":3,\"Name\":\"Abito Bello Invernale\",\"Sku\":\"1111\"},{\"Id\":5,\"Name\":\"Cosa Invernale\",\"Sku\":\"213\"}]"
//	assert.Equal(t, collezioneString, w.Body.String())
//}
//
//func TestEcommerceApi_GETCollezione2(t *testing.T) {
//	router := repo.New()
//	w := httptest.NewRecorder()
//	req, err := http.NewRequest("GET", "/collezione/2", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//	router.ServeHTTP(w, req)
//	assert.Equal(t, 200, w.Code)
//	collezioneString := "[{\"Id\":4,\"Name\":\"Abito Brutto Estivo\",\"Sku\":\"222\"},{\"Id\":1,\"Name\":\"Estive Nike\",\"Sku\":\"333\"},{\"Id\":2,\"Name\":\"Estive Adidas\",\"Sku\":\"444\"}]"
//	assert.Equal(t, collezioneString, w.Body.String())
//}
//func TestEcommerceApi_GETArticoli(t *testing.T) {
//	router := initServer()
//	w := httptest.NewRecorder()
//	req, _ := http.NewRequest("GET", "/articoli/", nil)
//	router.ServeHTTP(w, req)
//	assert.Equal(t, 200, w.Code)
//	articoliString := "[{\"Id\":3,\"Name\":\"Abito Bello Invernale\",\"Sku\":\"\"},{\"Id\":4,\"Name\":\"Abito Brutto Estivo\",\"Sku\":\"\"},{\"Id\":1,\"Name\":\"Estive Nike\",\"Sku\":\"\"},{\"Id\":2,\"Name\":\"Estive Adidas\",\"Sku\":\"\"},{\"Id\":5,\"Name\":\"Cosa Invernale\",\"Sku\":\"\"}]"
//	assert.Equal(t, articoliString, w.Body.String())
//}
