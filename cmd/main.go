package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printBanner() {
	fmt.Println("╔╗ ╦ ╦╔═╗╔═╗╔═╗╔═╗╔╗╔═╗ ╦")
	fmt.Println("╠╩╗║ ║║ ╦╚═╗║  ╠═╣║║║╔╩╦╝")
	fmt.Println("╚═╝╚═╝╚═╝╚═╝╚═╝╩ ╩╝╚╝╩ ╚═")
}

func showMenu() {
	fmt.Println("\n[1] HOST SCANNER         # Advanced bug host scanner with multiple modes")
	fmt.Println("[2] SUBFINDER            # Subdomain enumeration with passive discovery modes")
	fmt.Println("[3] IP LOOKUP            # Reverse IP lookup")
	fmt.Println("[4] FILE TOOLKIT         # Bug host list management")
	fmt.Println("[5] PORT SCANNER         # Discover open ports")
	fmt.Println("[6] DNS RECORD           # DNS record gathering")
	fmt.Println("[7] HOST INFO            # Detailed bug host analysis")
	fmt.Println("[8] HELP                 # Documentation and usage examples")
	fmt.Println("[9] UPDATE               # Self-update tool")
	fmt.Println("[0] EXIT                 # Quit application")
	fmt.Print("\nSelect an option: ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		printBanner()
		showMenu()

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Println("\n[+] Running HOST SCANNER...\n")
			// Call hostscanner.Start() (later)
		case "2":
			fmt.Println("\n[+] Running SUBFINDER...\n")
			// Call subfinder.Start() (later)
		case "3":
			fmt.Println("\n[+] Running IP LOOKUP...\n")
			// Call iplookup.Start() (later)
		case "4":
			fmt.Println("\n[+] Running FILE TOOLKIT...\n")
			// Call filetoolkit.Start() (later)
		case "5":
			fmt.Println("\n[+] Running PORT SCANNER...\n")
			// Call portscanner.Start() (later)
		case "6":
			fmt.Println("\n[+] Running DNS RECORD FETCHER...\n")
			// Call dnsrecord.Start() (later)
		case "7":
			fmt.Println("\n[+] Running HOST INFO MODULE...\n")
			// Call hostinfo.Start() (later)
		case "8":
			fmt.Println("\n[?] HELP SECTION:")
			fmt.Println(" - Use this tool to find valid bug hosts, subdomains, open ports, etc.")
			fmt.Println(" - Each module runs individually and outputs to /output/ directory.")
			fmt.Println(" - Customize scan settings in config/config.json\n")
		case "9":
			fmt.Println("\n[~] UPDATE: Coming soon...\n")
		case "0":
			fmt.Println("\n[-] Exiting... Goodbye!")
			return
		default:
			fmt.Println("\n[!] Invalid option, please try again.")
		}

		fmt.Print("\nPress ENTER to return to menu...")
		reader.ReadString('\n')
	}
}
