package handler

import (
	"fmt"
	"net/http"

	"github.com/golang-module/carbon/v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	now := carbon.Now().ToString()
	fmt.Fprintf(w, "<h1>Hello from Go! %s</h1>", now)
}
