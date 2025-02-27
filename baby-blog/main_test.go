// Filename: cmd/web/main_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)
func TestWeek1Handler(t *testing.T) {
    // set up request and response
    req := httptest.NewRequest("GET", "/week1", nil)
    rr := httptest.NewRecorder()
    // convert the home  function to a http.Handler type
    handler := http.HandlerFunc(week1Handler)
    // pass the fake HTTP request to the handler
    handler.ServeHTTP(rr, req)
    // perform  the assertions
    //1. check the status code for 200 OK
    status := rr.Code
    if status != http.StatusOK {
      t.Errorf("got %v, expected %v", status, http.StatusOK)
    }
    //2. check the response 
    expected := "<h1>Life in Adavance Database</h1>"
    //got := rr.Body.String()
    if got != expected {
      t.Errorf("got %v, expected %v", got, expected)
    }
}

