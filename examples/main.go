package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	wappalyzer "github.com/ducksify/wappalyzergo"
)

func main() {
	resp, err := http.DefaultClient.Get("https://www.hackerone.com")
	if err != nil {
		log.Fatal(err)
	}
	data, _ := io.ReadAll(resp.Body) // Ignoring error for example

	wappalyzerClient, err := wappalyzer.New()
	if err != nil {
		log.Fatal(err)
	}

	// Basic fingerprinting - returns map[string]struct{} (just technology names)
	fingerprints := wappalyzerClient.Fingerprint(resp.Header, data)
	fmt.Printf("Basic fingerprints: %v\n", fingerprints)
	// Output: map[Acquia Cloud Platform:{} Amazon EC2:{} Apache:{} Cloudflare:{} Drupal:{} PHP:{} Percona:{} React:{} Varnish:{}]

	// Fingerprinting with categories - returns map[string]CatsInfo
	fingerprintsWithCats := wappalyzerClient.FingerprintWithCats(resp.Header, data)
	fmt.Printf("Fingerprints with categories: %v\n", fingerprintsWithCats)

	// Fingerprinting with detailed info - returns map[string]AppInfo with description, website, icon, etc.
	fingerprintsWithInfo := wappalyzerClient.FingerprintWithInfo(resp.Header, data)
	fmt.Printf("Fingerprints with detailed info:\n")
	for tech, info := range fingerprintsWithInfo {
		fmt.Printf("  %s:\n", tech)
		fmt.Printf("    Description: %s\n", info.Description)
		fmt.Printf("    Website: %s\n", info.Website)
		fmt.Printf("    Icon: %s\n", info.Icon)
		fmt.Printf("    CPE: %s\n", info.CPE)
		fmt.Printf("    Categories: %v\n", info.Categories)
		fmt.Println()
	}

	// Fingerprinting with title - returns (map[string]struct{}, string)
	fingerprintsWithTitle, title := wappalyzerClient.FingerprintWithTitle(resp.Header, data)
	fmt.Printf("Page title: %s\n", title)
	fmt.Printf("Technologies: %v\n", fingerprintsWithTitle)
}
