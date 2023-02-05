package CheckDomain

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func CheckDomain(domain string) {

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string
	score := 0

	//	1. checking for MX Records
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error : %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
		score++
	}

	//	2. checking for TXT Records
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error : %v", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			score++
			spfRecord = record
			break
		}
	}
	//	3. Checking for dmarc records
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Printf("Error : %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			score++
			dmarcRecord = record
			break
		}
	}

	//Displaying records
	fmt.Println("-------------------------------------")
	fmt.Printf("Domain Name : %v ", domain)
	fmt.Printf("\nMX Recrod Present  : %v ", hasMX)
	fmt.Printf("\nSPF Record Present  : %v ", hasSPF)
	fmt.Printf("\nSPF : %v ", spfRecord)
	fmt.Printf("\nDMARC Recrod Present  : %v ", hasDMARC)
	fmt.Printf("\nDMARC : %v \n", dmarcRecord)
	fmt.Println("-------------------------------------")

	fmt.Println("\nEnter another domain name : ")

}
