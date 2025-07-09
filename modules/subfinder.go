package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Subfinder struct
type Subfinder struct {
	Timeout time.Duration
}

// Public source: crt.sh
func (sf *Subfinder) FetchFromCrtSh(domain string) []string {
	url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)

	client := http.Client{
		Timeout: sf.Timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[-] Error fetching subdomains: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var results []map[string]interface{}
	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Printf("[-] JSON parse error: %v\n", err)
		return nil
	}

	subdomains := make(map[string]bool)
	for _, entry := range results {
		if nameValue, ok := entry["name_value"].(string); ok {
			for _, sub := range strings.Split(nameValue, "\n") {
				if strings.HasSuffix(sub, domain) {
					subdomains[sub] = true
				}
			}
		}
	}

	finalList := []string{}
	for sub := range subdomains {
		finalList = append(finalList, sub)
	}

	return finalList
}

// Entry method
func (sf *Subfinder) Start(domain string) {
	fmt.Printf("[~] Searching subdomains for: %s\n", domain)
	subs := sf.FetchFromCrtSh(domain)
	if len(subs) == 0 {
		fmt.Println("[-] No subdomains found.")
		return
	}

	for _, sub := range subs {
		fmt.Println("[+] Found:", sub)
	}
}
