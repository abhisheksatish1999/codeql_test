package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Insecure: certificate verification is disabled, which should be
	// flagged by CodeQL's go/disabled-certificate-check query.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://example.com")
	if err != nil {
		fmt.Println("request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
