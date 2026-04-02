package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func SearchGoogleBooks(query string) (*GoogleBooksResponse, error) {
	// 1. Escapar a query para ser usada na URL (ex: transforma espaços em +)
	safeQuery := url.QueryEscape(query)
	apiURL := fmt.Sprintf("https://www.googleapis.com/books/v1/volumes?q=%s", safeQuery)

	// 2. Fazer a requisição GET
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 3. Decodificar o JSON para nossa Struct
	var result GoogleBooksResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}