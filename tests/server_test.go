package tests

import (
	"go-avanzado/handler"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHelloHandler(t *testing.T) {

	r := gin.Default()

	r.GET("/hello", handler.HelloHandler)

	req, _ := http.NewRequest(
		"GET",
		"/hello",
		nil,
	)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 200 {

		t.Errorf(
			"expected 200 got %d",
			w.Code,
		)
	}
}
