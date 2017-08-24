package backend

import (
	"net/http"
	"fmt"
)

func Recover(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
