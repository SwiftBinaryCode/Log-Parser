package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}
type parser struct {
	sum     map[string]result
	domains []string
	total   int
	lines   int
}

func main() {
	//log parser using strucs

	p := parser{
		sum: make(map[string]result),
	}

	//scan the standard-in line by line
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		p.lines++
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input:  %v (line #%d)\n", fields, p.lines)
			return
		}

		domain := fields[0]

		//sum the ttal visits per domain
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input:  %v (line #%d)\n", fields[1], p.lines)
			return
		}
		//Collect the unique domains
		if _, ok := p.sum[domain]; !ok {

			p.domains = append(p.domains, domain)
		}

		//Keep track of total and per domain visits
		p.total += visits

		p.sum[domain] = result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}
	}

	//print te visits per domain
	fmt.Printf("%-30s %10s\n ", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))
	sort.Strings(p.domains)

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)

	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

	if err := in.Err(); err != nil {
		fmt.Println("> ERR:", err)
	}
}
