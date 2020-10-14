package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

var domains string
var output string
var hosts []string

func main() {
	flag.StringVar(&domains, "w", "", "List with the domains")
	flag.StringVar(&output, "o", "", "The output file")
	flag.Parse()

	if domains != "" {
		file, err := os.Open(domains)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		reader := bufio.NewScanner(file)

		for reader.Scan() {
			domain := reader.Text()
			lookup, _ := net.LookupHost(domain)
			hosts = append(hosts, fmt.Sprintf("%v\n", strings.Join(lookup, "\n")))
			fmt.Printf("%v\n", strings.Join(lookup, "\n"))
		}
		if output != "" {
			defer file.Close()
			file, err := os.Create(output)
			if err != nil {
				panic(err)
			}
			file.WriteString(strings.Join(hosts, ""))
		}

	} else {
		fmt.Println("You need to pass a file with domains!")
	}
}
