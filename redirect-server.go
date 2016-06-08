package main

import (
	"net/http"
	"strings"
)

func redirect(w http.ResponseWriter, domain string, path string) {
	url := strings.Join([]string{"https://", domain, path}, "")
	w.Header().Set("Location", url)
	w.WriteHeader(301)
}

func handler(w http.ResponseWriter, r *http.Request) {
	redirect(w, getRedirectHost(r.Host), r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func getRedirectHost(host string) string {
	domainParts := strings.Split(strings.Split(host, ":")[0], ".")

	var domain string

	if last(domainParts) == "br" {
		domain = strings.Join(domainParts[len(domainParts)-3:], ".")
	} else {
		domain = strings.Join(domainParts[len(domainParts)-2:], ".")
	}

	if domainParts[0] != "www" {
		domain = strings.Join([]string{"www", domain}, ".")
	}

	return domain
}

func last(list []string) string {
	return list[len(list)-1]
}
