package objects

import (
	"net/http"
	"os"
	"strings"
	"log"
	"io"
)

func put(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Print(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(f, r.Body)
}