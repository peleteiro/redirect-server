package main

import (
	"fmt"
	"golang.org/x/net/publicsuffix"
	"net"
	"net/http"
	"os"
	"strings"
)

var PORT = "8080"
var SERVICE_FQDN string = "foo.com"
var SERVICE_FQDN_SUFFIX string = ".foo.com."

func redirect(w http.ResponseWriter, domain string, path string) {
	url := strings.Join([]string{"https://", domain, path}, "")
	w.Header().Set("Location", url)
	w.WriteHeader(301)
}

func handler(w http.ResponseWriter, r *http.Request) {
	redirect(w, getRedirectHost(r.Host), r.URL.Path)
}

func main() {
	if os.Getenv("PORT") == "" {
		PORT = os.Getenv("PORT")
	}
	SERVICE_FQDN = os.Getenv("SERVICE_FQDN")
	if SERVICE_FQDN == "" {
		print("Must have an env variable `SERVICE_FQDN`.")
		os.Exit(1)
	}
	SERVICE_FQDN_SUFFIX = fmt.Sprintf(".%s.", SERVICE_FQDN)
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil)
}

func getRedirectHost(host string) string {
	cname, err := net.LookupCNAME(host)
	if err != nil || cname == "" {
		return getRedirectHostByHost(host)
	}
	redirectHost := GetRedirectHostByCNAME(cname)
	if redirectHost == "" {
		return getRedirectHostByHost(host)
	}
	return redirectHost
}

func getRedirectHostByHost(host string) string {
	tld, _ := publicsuffix.PublicSuffix(host)

	hostParts := strings.Split(strings.TrimSuffix(host, fmt.Sprintf(".%s", tld)), ".")
	name := last(hostParts)

	if hostParts[0] == "www" {
		return fmt.Sprintf("%s.%s", name, tld)
	}

	return fmt.Sprintf("www.%s.%s", name, tld)
}

func GetRedirectHostByCNAME(cname string) string {
	if cname == "" {
		return ""
	}

	if !strings.HasSuffix(cname, SERVICE_FQDN_SUFFIX) {
		return ""
	}

	return strings.TrimSuffix(cname, SERVICE_FQDN_SUFFIX)
}

func last(list []string) string {
	return list[len(list)-1]
}
