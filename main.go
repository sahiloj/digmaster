package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func performDNSLookup(subdomain string) {
	fmt.Printf("Performing DNS lookup for %s\n", subdomain)
	cmd := exec.Command("dig", "@1.1.1.1", subdomain)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error performing DNS lookup for %s: %v\n", subdomain, err)
		return
	}
	fmt.Println(string(output))
	fmt.Println("-----------------------------------------")
}

func main() {
	
	subdomainsFile := flag.String("s", "", "Path to subdomains file")
	flag.Parse()

	if *subdomainsFile != "" {
	fmt.Println(`
     _ _                        _            
  __| (_) __ _  /\/\   __ _ ___| |_ ___ _ __ 
 / _` + "`" + ` | |/ _` + "`" + ` |/    \ / _` + "`" + ` / __| __/ _ | '__|
| (_| | | (_| / /\/\ | (_| \__ | ||  __| |   
 \__,_|_|\__, \/    \/\__,_|___/\__\___|_|   
         |___/                               

            by @Sahil Ojha`)
}

	// Read subdomains file
	fileData, err := ioutil.ReadFile(*subdomainsFile)
	if err != nil {
		fmt.Println("Error reading subdomains file:", err)
		os.Exit(1)
	}

	subdomains := strings.Split(string(fileData), "\n")

	for _, subdomain := range subdomains {
		subdomain = strings.TrimSpace(subdomain)
		if subdomain != "" {
			performDNSLookup(subdomain)
		}
	}
}
