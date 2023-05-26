package main

import "fmt"

func main() {
	var res, sum, average float64

	for i := 0; i < 3; i++ {
		fmt.Scan(&res)
		sum += res
	}

	average = sum / 3

	fmt.Println("mean score is:", average)

	if average >= 60 {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}

}
