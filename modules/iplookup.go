package modules

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// IPLookup struct
type IPLookup struct {
	Timeout time.Duration
}

// Perform reverse IP lookup using HackerTarget API
func (il *IPLookup) reverseLookup(ip string) {
	api := fmt.Sprintf("https://api.hackertarget.com/reverseiplookup/?q=%s", ip)

	client := http.Client{
		Timeout: il.Timeout,
	}
	resp, err := client.Get(api)
	if err != nil {
		fmt.Printf("[-] Error on lookup for %s: %v\n", ip, err)
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ",") {
			domain := strings.Split(line, ",")[0]
			fmt.Println("[+] Found:", domain)
		} else if strings.Contains(line, "error") {
			fmt.Printf("[-] %s => %s\n", ip, line)
			break
		} else {
			fmt.Println("[+] Found:", line)
		}
	}
}

// Expand CIDR range and run reverse lookup for each IP
func (il *IPLookup) ProcessCIDR(cidr string) {
	ips := []string{}
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Printf("[-] Invalid CIDR: %s\n", cidr)
		return
	}

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// Skip network/broadcast IPs
	if len(ips) > 2 {
		ips = ips[1 : len(ips)-1]
	}

	for _, ip := range ips {
		fmt.Printf("\n[~] Looking up %s\n", ip)
		il.reverseLookup(ip)
	}
}

// Helper to increment IP
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// Start function entry point
func (il *IPLookup) Start(input string) {
	if strings.Contains(input, "/") {
		il.ProcessCIDR(input)
	} else {
		il.reverseLookup(input)
	}
}
