package raspy

import (
	"fmt"
	"net/http"
    "github.com/tsoonjin/raspy/pkg/simplex/simplify"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "someone"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}

SimplifyHandler := simplify.Handler;
