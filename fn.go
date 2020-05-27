package raspy

import (
	"fmt"
	"net/http"
    "github.com/tsoonjin/raspy/pkg/simplex"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "someone"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

var SimplifyHandler = simplex.Handler;
