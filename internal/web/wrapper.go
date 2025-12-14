package web

import "net/http"

func invertChain(original, added http.Handler, source string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			if r.Method != http.MethodPost {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
		}

		if source == r.URL.Path {
			added.ServeHTTP(w, r)
		}

		original.ServeHTTP(w, r)
	})
}
