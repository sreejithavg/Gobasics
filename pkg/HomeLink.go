package pkg

import (
	"fmt"
	"net/http"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!....\n Welcome Home")
}
