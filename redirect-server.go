package main

import (
	"fmt"
	"net/http"
	"golang.org/x/net/publicsuffix"
	"net/url"
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
	urlParts, _ := url.Parse(fmt.Sprintf("https://%s", host))
	tld, _ := publicsuffix.PublicSuffix(urlParts.Host)

	hostParts := strings.Split(host, fmt.Sprintf("%s", tld))
	hostParts = strings.Split(strings.TrimSuffix(hostParts[0], "."), ".")

	domain := fmt.Sprintf("%s.%s", last(hostParts), tld)
	if hostParts[0] != "www" {
		domain = fmt.Sprintf("www.%s", domain)
	}

	domain = fmt.Sprintf("%s%s", domain, urlParts.Path)
	if urlParts.RawQuery != "" {
		domain = fmt.Sprintf("%s?%s", domain, urlParts.RawQuery)
	}
	if urlParts.Fragment != "" {
		domain = fmt.Sprintf("%s#%s", domain, urlParts.Fragment)
	}

	return domain
}

func last(list []string) string {
	return list[len(list)-1]
}
