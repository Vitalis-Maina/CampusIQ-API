package main

import (
	"fmt"
	"sort"
)

type Applicant struct {
	FirstName string
	LastName  string
	Gpa       float64
}

func main() {
	var totalNum, acceptable int
	fmt.Scan(&totalNum, &acceptable)
	var applicants Applicant

	var applicantList []Applicant

	fmt.Println("Enter applicant details")
	for i := 0; i < totalNum; i++ {
		fmt.Scan(&applicants.FirstName, &applicants.LastName, &applicants.Gpa)
		applicantList = append(applicantList, applicants)
	}
	sort.Slice(applicantList, func(i, j int) bool {
		if applicantList[i].Gpa != applicantList[j].Gpa {
			return applicantList[i].Gpa > applicantList[j].Gpa
		} else if applicantList[i].FirstName != applicantList[j].FirstName {
			return applicantList[i].FirstName < applicantList[j].FirstName
		}
		return applicantList[i].LastName > applicantList[j].LastName
	})
	fmt.Println("successfull applicants:")
	for i := 0; i < acceptable; i++ {
		fmt.Println(applicantList[i].FirstName, applicantList[i].LastName, applicantList[i].Gpa)
	}
}
