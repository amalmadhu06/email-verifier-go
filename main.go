package main

import (
	"bufio"
	"fmt"
	check "github.com/amalmadhu06/email-verifier-go/CheckDomain"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter a domain name here : ")
	//receiving input from command line
	for scanner.Scan() {
		fmt.Println("Checking records ....")
		check.CheckDomain(scanner.Text())

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Unable to read the input : %v \n", err)
	}

}
