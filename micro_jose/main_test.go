package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(apiHandler)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Código de estado incorrecto: obtuve %v en lugar de %v",
			status, http.StatusOK)
	}

	var response Response
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expected := Response{
		Message: "get!",
	}
	if response != expected {
		t.Errorf("Respuesta incorrecta: obtuve '%v' en lugar de '%v'",
			response, expected)
	}
}

func TestPostHandler(t *testing.T) {
	requestData := RequestData{
		Data: "example data",
	}
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/api/post", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(postHandler) // Corregimos el error aquí, debe ser postHandler
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Código de estado incorrecto: obtuve %v en lugar de %v",
			status, http.StatusOK)
	}

	var response Response
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	expected := Response{
		Message: "POST!",
		Data:    "example data",
	}
	if response != expected {
		t.Errorf("Respuesta incorrecta: obtuve '%v' en lugar de '%v'",
			response, expected)
	}
}
