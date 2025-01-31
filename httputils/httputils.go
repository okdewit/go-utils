package httputils

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/url"
)

func BasicAuth(handler http.HandlerFunc, username, password, message string) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != username || hashCompare(password, pass) != nil {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+message+`"`)
			w.WriteHeader(401)
			_, _ = w.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(w, r)
	}
}

func hashCompare(password string, pass string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return bcrypt.CompareHashAndPassword(hash, []byte(pass))
}

func AddParams(uri *url.URL, params map[string]string) {
	p := uri.Query()
	for k, v := range params {
		p.Add(k, v)
	}
	uri.RawQuery = p.Encode()
}