package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const Help = "NAME:\n\tcli_net - client displays network interfaces and info.\n\n" +
	"USAGE:\n\tcli_net [global options] command [command options] [arguments...]\n\n" +
	"VERSION:\n\t0.0.0\n\n" +
	"COMMANDS:\n\thelp, h Shows a list of commands or help for one command\t...\n\n" +
	"GLOBAL OPTIONS:\n\t--version Shows version information\n\n"

func StartCli(r io.Reader, w io.Writer) {
	version := Version{}
	lstInterfaces := LstInterfaces{}
	netInterface := DetailInterface{}
	apiError := ApiError{}
	var server string
	scanner := bufio.NewScanner(r)
	fmt.Println("Hello from the console")
	validIpAddressRegex := `(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`
	validRequestParams := `--server\s+` + validIpAddressRegex + `\s+--port\s+[0-9]{4}$`
	// var text string
	for scanner.Scan() {
		if scanner.Text() == "help" || scanner.Text() == "h" {
			io.WriteString(w, Help)
		}
		if scanner.Text() == "--version" {
			res, err := http.Get("http://server:8080/service/version")
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			json.NewDecoder(res.Body).Decode(&version)
			io.WriteString(w, version.Version+"\n")
			//fmt.Println(version.Version)
		}
		r, err := regexp.Compile(`^list\s+` + validRequestParams)
		if err != nil {
			panic(err)
		}
		if r.MatchString(scanner.Text()) {
			// list --server 127.0.0.1 --port 8080
			text := strings.Split(strings.Split(strings.Replace(scanner.Text(), " ", "", -1), "--server")[1], "--port")
			if text[0] == "127.0.0.1" {
				server = "server"
			} else {
				server = text[0]
			}
			res, err := http.Get(fmt.Sprintf("http://%s:%s/service/v1/interfaces", server, text[1]))
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			json.NewDecoder(res.Body).Decode(&lstInterfaces)
			fmt.Println(lstInterfaces.Intrefaces)
		}
		r, err = regexp.Compile(`^show\s+\w+\s+` + validRequestParams)
		if err != nil {
			panic(err)
		}
		if r.MatchString(scanner.Text()) {
			// show lo --server 127.0.0.1 --port 8080
			splitingText := strings.Split(strings.Replace(scanner.Text(), " ", "", -1), "--server")
			var inter string = strings.Replace(splitingText[0], "show", "", -1)
			var network []string = strings.Split(splitingText[1], "--port")
			if network[0] == "127.0.0.1" {
				server = "server"
			} else {
				server = network[0]
			}
			res, err := http.Get(fmt.Sprintf("http://%s:%s/service/v1/interfaces/%s", server, network[1], inter))
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()
			if res.StatusCode == http.StatusOK {
				json.NewDecoder(res.Body).Decode(&netInterface)
				r, err := regexp.Compile(validIpAddressRegex)
				if err != nil {
					panic(err)
				}
				var ipv4 string
				var ipv6 string
				for _, addr := range netInterface.Inet_addr {
					if r.MatchString(addr) {
						ipv4 = addr
					} else {
						ipv6 = addr
					}
				}
				fmt.Printf("%s:\tHW_addr: %s\n\tIPv4: %s\n\tIPv6: %s\n\tMTU: %d\n",
					netInterface.Name, netInterface.Hw_addr,
					ipv4, ipv6, netInterface.MTU)
			} else if res.StatusCode == http.StatusNotFound || res.StatusCode == http.StatusInternalServerError {
				json.NewDecoder(res.Body).Decode(&apiError)
				fmt.Printf("Error found:\n\t%s\n", apiError.Error)
			}
		}
	}
}

func main() {
	StartCli(os.Stdin, os.Stdout)
}
