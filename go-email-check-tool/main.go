package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// os.Stdin is open Files pointing to the standard input
// NewScanner returns a new Scanner to read from os.Stdin(입력받은 정보)
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	//Scan advances the Scanner to the next token, token 값을 불러오고 stop되면 false 반환
	for scanner.Scan() {
		checkDomain(scanner.Text()) //Text() returns most recetn token generated by a call to Scan
	}

	if err := scanner.Err(); err != nil { //에러 발생시 메세지
		log.Fatal("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	//LookupMX returns the DNS MX records for the given domain name sorted by preference.
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	if len(mxRecords) > 0 { //mxRecords에 length가 0보다 크면 즉 받아온 DNS MX records가 있으면
		hasMX = true //hasMX는 true
	}

	//LookupTXT returns the DNS TXT records for the given domain name.
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range txtRecords { //range는 주어진 array의 index와 value를 순서대로 반환
		if strings.HasPrefix(record, "v=spf1") { //strings.HasPrefix는 record에서 주어진 prefix("v=spf1")로 시작되는 data가 있으면 true반환
			hasSPF = true
			spfRecord = record
			break
		}
	}

	//LookupTXT returns the DNS TXT records for the given domain name.
	dmarcRecords, err := net.LookupTXT("_dmarc" + domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, record := range dmarcRecords { //range는 주어진 array의 index와 value를 순서대로 반환
		if strings.HasPrefix(record, "v=DMARC1") { //strings.HasPrefix는 record에서 주어진 prefix("v=DMARC1")로 시작되는 data가 있으면 true반환
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v", domain, hasMX, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
