package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/EFForg/starttls-check/checker"
)

func manualValidation(csvFilename string) {
	b, err := ioutil.ReadFile(csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	in := string(b)

	r := csv.NewReader(strings.NewReader(in))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		domain := record[0]
		mxHostnames := strings.Split(record[2], ",")
		result := checker.CheckDomain(domain, mxHostnames)
		fmt.Println(result)
	}
	fmt.Println(csvFilename)
}
