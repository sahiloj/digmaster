<pre>
     _ _                        _            
  __| (_) __ _  /\/\   __ _ ___| |_ ___ _ __ 
 / _` | |/ _` |/    \ / _` / __| __/ _ | '__|
| (_| | | (_| / /\/\ | (_| \__ | ||  __| |   
 \__,_|_|\__, \/    \/\__,_|___/\__\___|_|   
         |___/                               

            by @Sahil Ojha
</pre>
# digMaster
digMaster is a command-line tool written in Go for performing DNS lookups. It allows users to quickly query DNS records for multiple subdomains specified in a file or entered interactively. The tool utilizes the dig command-line utility to retrieve DNS information and presents the results in a readable format.

# Installation
* From source:
  ```From source:
  go install github.com/sahiloj/digmaster@latest
  ```
# Usage
   ```
    digmaster -s subdomains.txt
   ```

# Features:

* Perform DNS lookups for multiple subdomains specified in a file or entered interactively.
* Displays DNS records such as A, AAAA, CNAME, MX, TXT, etc., for each subdomain.
* Supports both interactive mode and file input for flexibility in usage.
* Simple and easy-to-use interface, ideal for both beginners and experienced users.
* Developed with Go for cross-platform compatibility.
