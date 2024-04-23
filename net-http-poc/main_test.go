package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	handler "net-http-poc/internal/v1"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestHandlerAlive tests the /v1/alive endpoint just calling the handler
func TestHandlerGetAlive(t *testing.T) {
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/v1/alive", nil)

	handler.Alive(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}
	if !strings.Contains(wr.Body.String(), "OK") {
		t.Errorf(`response body "%s" does not contain "OK"`, wr.Body.String())
	}
}

// TestGetAliveMessage tests the /v1/alive endpoint using a client and a server
func TestGetAliveMessage(t *testing.T) {
	http.HandleFunc("GET /v1/alive", handler.Alive)

	svr := httptest.NewServer(http.DefaultServeMux)
	defer svr.Close()

	c := NewClient(http.DefaultClient, svr.URL)
	m, err := c.GetAliveMessage()
	if err != nil {
		t.Fatalf("error in GetAliveMessage: %v", err)
	}
	if m.Message != "OK" {
		t.Errorf(`message %s should contain string "OK"`, m.Message)
	}
}

type Client struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient(httpClient *http.Client, baseURL string) Client {
	return Client{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

// GetAliveMessage sends a GET request to the /v1/alive endpoint and returns the response
func (c *Client) GetAliveMessage() (*handler.AliveResponse, error) {
	res, err := c.httpClient.Get(c.baseURL + "/v1/alive")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got status code %d", res.StatusCode)
	}
	var m handler.AliveResponse
	if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

var CompilationPayload = handler.Compilation{
	Fqbn: "esp8266:esp8266:esp8266",
	OTA:  false,
	Sketch: handler.Sketch{
		Files: []handler.File{},
		Ino: handler.File{
			Name: "test.ino",
			Data: "#include <Arduino.h>\nvoid setup() {}\nvoid loop() {}",
		},
		Metadata: []handler.Library{},
		Name:     "test",
	},
	Verbose: false,
}

// TestHandlerPostCompilations tests the /v1/compilations endpoint just calling the handler
func TestHandlerPostCompilations(t *testing.T) {

	//initializes the compilations queue
	handler.Init()
	body, _ := json.Marshal(CompilationPayload)
	bodyReader := bytes.NewReader(body)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/compilations", bodyReader)

	handler.CreateCompilation(wr, req)
	if wr.Code != http.StatusCreated {
		t.Errorf("got HTTP status code %d, expected 201", wr.Code)
	}
	if !strings.Contains(wr.Body.String(), "status") {
		t.Errorf(`response body "%s" does not contain "status"`, wr.Body.String())
	}
}

// TestPostCompilationsMessage tests the /v1/compilations endpoint using a client and a server
func TestPostCompilationsMessage(t *testing.T) {

	handler.Init()
	http.HandleFunc("POST /v1/compilations", handler.CreateCompilation)

	svr := httptest.NewServer(http.DefaultServeMux)
	// defer svr.Close()

	c := NewClient(http.DefaultClient, svr.URL)
	m, err := c.PostCompilationsMessage()
	if err != nil {
		t.Fatalf("error in GetCompilationsMessage: %v", err)
	}
	if m.Status != "created" {
		t.Errorf(`status %s should contain string "created"`, m.Status)
	}
}

// PostCompilationsMessage sends a POST request to the /v1/compilations endpoint and returns the response
func (c *Client) PostCompilationsMessage() (*handler.CompilationResponse, error) {
	body, _ := json.Marshal(CompilationPayload)
	bodyReader := bytes.NewReader(body)

	res, err := c.httpClient.Post(c.baseURL+"/v1/compilations", "application/json", bodyReader)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("got status code %d", res.StatusCode)
	}
	var m handler.CompilationResponse
	if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

// TestHandlerStopCompilations tests the /v1/compilations/{id}/cancel endpoint just calling the handler
func TestHandlerStopCompilations(t *testing.T) {
	// first we have to create a compilation:
	//initializes the compilations queue
	handler.Init()
	body, _ := json.Marshal(CompilationPayload)
	bodyReader := bytes.NewReader(body)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/compilations", bodyReader)

	handler.CreateCompilation(wr, req)
	if wr.Code != http.StatusCreated {
		t.Errorf("got HTTP status code %d, expected 201", wr.Code)
	}
	if !strings.Contains(wr.Body.String(), "status") {
		t.Errorf(`response body "%s" does not contain "status"`, wr.Body.String())
	}

	// we marshall the json to get the UUID
	var compilationResponse handler.CompilationResponse
	if err := json.NewDecoder(wr.Body).Decode(&compilationResponse); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}
	defer req.Body.Close()

	// then we stop the compilation
	wr = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/v1/compilations/{id}/cancel", nil)
	req.SetPathValue("id", compilationResponse.UUID)
	handler.StopCompilation(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}
	if !strings.Contains(wr.Body.String(), "cancelled") {
		t.Errorf(`response body "%s" does not contain "cancelled"`, wr.Body.String())
	}
}

// TestHandlerGetCompilations tests the /v1/compilations{id} endpoint just calling the handler
func TestHandlerGetCompilations(t *testing.T) {
	// first we have to create a compilation:
	//initializes the compilations queue
	handler.Init()
	body, _ := json.Marshal(CompilationPayload)
	bodyReader := bytes.NewReader(body)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/compilations", bodyReader)

	handler.CreateCompilation(wr, req)
	if wr.Code != http.StatusCreated {
		t.Errorf("got HTTP status code %d, expected 201", wr.Code)
	}
	if !strings.Contains(wr.Body.String(), "status") {
		t.Errorf(`response body "%s" does not contain "status"`, wr.Body.String())
	}

	// we marshall the json to get the UUID
	var compilationResponse handler.CompilationResponse
	if err := json.NewDecoder(wr.Body).Decode(&compilationResponse); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}
	defer req.Body.Close()

	// then we get the compilation
	wr = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/v1/compilations/{id}/cancel", nil)
	req.SetPathValue("id", compilationResponse.UUID)
	handler.GetCompilations(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}
	if !strings.Contains(wr.Body.String(), "created") {
		t.Errorf(`response body "%s" does not contain "cancelled"`, wr.Body.String())
	}
}

// TestHandlerGetArtifacts tests the /v1/compilations/{id}/artifacts endpoint just calling the handler
func TestHandlerGetArtifacts(t *testing.T) {
	// first we have to create a compilation:
	//initializes the compilations queue
	handler.Init()
	body, _ := json.Marshal(CompilationPayload)
	bodyReader := bytes.NewReader(body)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/compilations", bodyReader)

	handler.CreateCompilation(wr, req)
	if wr.Code != http.StatusCreated {
		t.Errorf("got HTTP status code %d, expected 201", wr.Code)
	}
	if !strings.Contains(wr.Body.String(), "status") {
		t.Errorf(`response body "%s" does not contain "status"`, wr.Body.String())
	}

	// we marshall the json to get the UUID
	var compilationResponse handler.CompilationResponse
	if err := json.NewDecoder(wr.Body).Decode(&compilationResponse); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}
	defer req.Body.Close()

	// then we get the artifatcs
	wr = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/v1/compilations/{id}/artifact", nil)
	req.SetPathValue("id", compilationResponse.UUID)
	handler.GetArtifacts(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}
	if !strings.Contains(wr.Body.String(), "bin") {
		t.Errorf(`response body "%s" does not contain "bin"`, wr.Body.String())
	}

	// then we get the artifatcs, filtering by type
	wr = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/v1/compilations/{id}/artifact", nil)
	req.SetPathValue("id", compilationResponse.UUID)
	req.URL.RawQuery = "type=elf"
	handler.GetArtifacts(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}
	if !strings.Contains(wr.Body.String(), "elf") {
		t.Errorf(`response body "%s" does not contain "elf"`, wr.Body.String())
	}
}

// TestHandlerGetLogs tests the /v1/compilations/{id}/logs endpoint just calling the handler
func TestHandlerGetLogs(t *testing.T) {
	// first we have to create a compilation:
	//initializes the compilations queue
	handler.Init()
	body, _ := json.Marshal(CompilationPayload)
	bodyReader := bytes.NewReader(body)
	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/v1/compilations", bodyReader)

	handler.CreateCompilation(wr, req)
	if wr.Code != http.StatusCreated {
		t.Errorf("got HTTP status code %d, expected 201", wr.Code)
	}
	if !strings.Contains(wr.Body.String(), "status") {
		t.Errorf(`response body "%s" does not contain "status"`, wr.Body.String())
	}

	// we marshall the json to get the UUID
	var compilationResponse handler.CompilationResponse
	if err := json.NewDecoder(wr.Body).Decode(&compilationResponse); err != nil {
		t.Fatalf("error decoding response: %v", err)
	}
	defer req.Body.Close()

	// then we get the artifatcs
	wr = httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/v1/compilations/{id}/artifact", nil)
	req.SetPathValue("id", compilationResponse.UUID)
	handler.GetLogs(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}
	if !strings.Contains(wr.Body.String(), "stdout") {
		t.Errorf(`response body "%s" does not contain "stdout"`, wr.Body.String())
	}
}
