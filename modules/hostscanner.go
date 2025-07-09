package modules

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

// Core scanner struct
type HostScanner struct {
	Timeout time.Duration
	Non302  bool
	SSLMode bool
}

// StartScan scans a list of domains
func (hs *HostScanner) StartScan(domains []string) {
	for _, domain := range domains {
		url := "https://" + domain
		if hs.SSLMode {
			hs.checkSSL(domain)
		} else {
			hs.checkHTTP(url)
		}
	}
}

func (hs *HostScanner) checkHTTP(url string) {
	client := http.Client{
		Timeout: hs.Timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if hs.Non302 {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[-] %s -> %v\n", url, err)
		return
	}
	defer resp.Body.Close()

	if hs.Non302 && (resp.StatusCode == 301 || resp.StatusCode == 302) {
		fmt.Printf("[~] %s skipped (redirect %d)\n", url, resp.StatusCode)
		return
	}

	fmt.Printf("[+] %s -> %d %s\n", url, resp.StatusCode, http.StatusText(resp.StatusCode))
}

func (hs *HostScanner) checkSSL(domain string) {
	conn, err := tls.DialWithDialer(&tls.Dialer{Timeout: hs.Timeout}, "tcp", domain+":443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		fmt.Printf("[-] TLS Failed: %s -> %v\n", domain, err)
		return
	}
	defer conn.Close()

	state := conn.ConnectionState()
	fmt.Printf("[🔐] %s\n", domain)
	fmt.Printf("     TLS Version: %x\n", state.Version)
	fmt.Printf("     Cipher Suite: %x\n", state.CipherSuite)
	if len(state.PeerCertificates) > 0 {
		fmt.Printf("     Certificate CN: %s\n", state.PeerCertificates[0].Subject.CommonName)
	}
}
