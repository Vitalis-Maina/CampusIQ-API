package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Applicant struct {
	FirstName string
	LastName  string
	Gpa       float64
}

var applicants []Applicant

func main() {
	GetApplicants()
	GenerateAdmitted(applicants)
	IsAccepted(applicants)

}
func GetApplicants() {
	var totalNum, acceptable int

	fmt.Scan(&totalNum, &acceptable)
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < totalNum; i++ {

		scanner.Scan()
		field := strings.Split(scanner.Text(), " ")

		gpa, err := strconv.ParseFloat(field[2], 64)
		if err != nil {
			log.Fatal(err)
		}
		newApplicant := Applicant{
			FirstName: field[0],
			LastName:  field[1],
			Gpa:       gpa,
		}
		applicants = append(applicants, newApplicant)

	}

}

func GenerateAdmitted(applicants []Applicant) {

	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].Gpa != applicants[j].Gpa {
			return applicants[i].Gpa > applicants[j].Gpa
		} else if applicants[i].FirstName != applicants[j].FirstName {
			return applicants[i].FirstName < applicants[j].FirstName
		}
		return applicants[i].LastName < applicants[i].LastName
	})
}

func IsAccepted(applicants []Applicant) {
	for i := range applicants {
		if applicants[i].Gpa >= 60 {
			fmt.Println("Congratulations, you are accepted!")
			fmt.Println("successfull applicants:")

			fmt.Println(applicants[i].FirstName, applicants[i].LastName, applicants[i].Gpa)
			break
		} else {
			fmt.Printf("Hello %v %v, We regret to inform you that we will not be able to offer you admission\n", applicants[i].FirstName, applicants[i].LastName)

		}
	}

}
