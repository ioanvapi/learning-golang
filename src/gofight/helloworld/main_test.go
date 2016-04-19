package main

import (
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestBasicHelloWorld(t *testing.T) {
	r := gofight.New()

	r.GET("/").
		// trun on the debug mode.
		SetDebug(true).
		Run(BasicEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {

			assert.Equal(t, "Hello World", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
