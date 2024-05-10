package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
	database "github.com/juliazuin/AluraChallenge/app/dabatase"
)

func TestPingRoute(t *testing.T) {

	db := database.NewDB()
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(setupRouter(db.Db))

	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	/*body := map[string]string{
		"descricao":  "teste camil",
		"valor":      "14.50",
		"data_atual": "18/01/2022",
	}
	body_data, _ := json.Marshal(body)*/

	resp, err := http.Get(fmt.Sprintf("%s/despesas", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.Status)

}
func TestPostDespesa(t *testing.T) {
	db := database.NewDB()
	// The setupServer method, that we previously refactored
	// is injected into a test server
	ts := httptest.NewServer(setupRouter(db.Db))

	// Shut down the server and block until all requests have gone through
	defer ts.Close()

	// Make a request to our server with the {base url}/ping
	body, _ := json.Marshal(map[string]string{
		"descricao":  "teste",
		"valor":      "3000.00",
		"data_atual": "01/02/2022",
	})

	responseBody := bytes.NewBuffer(body)
	println(fmt.Sprintf("%s/despesas", ts.URL))

	res, err := ts.Client().Post(fmt.Sprintf("%s/despesas", ts.URL), "", responseBody)

	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.Status)

}
