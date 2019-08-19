package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func input_parser() (cases[][]float64){
	file, err := os.Open("Problem B\\1.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	number_of_cases, err := strconv.Atoi(scanner.Text())

	for test_case := 0; test_case < number_of_cases; test_case++ {
		scanner.Scan();
		buffer := strings.Fields(scanner.Text())
		var one_case[]float64
		s, err := strconv.ParseFloat(buffer[0], 64)
		if err == nil {
			one_case = append(one_case, s)
		}

		s, err = strconv.ParseFloat(buffer[1], 64)
		if err == nil {
			one_case = append(one_case, s)
		}

		cases = append(cases, one_case)
	}

	return cases
}

func calc_angle(cases[][]float64) (angles[]float64) {
	for _,values := range cases {
		angle := math.Asin((values[1] * 9.8) / (values[0] * values[0])) / 2
		angles = append(angles, angle)
	}

	return angles
}

func result_writer(angles[]float64) {
	f, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for case_id,angle := range angles {
		_, err := f.WriteString("Case #" + strconv.Itoa(case_id + 1) +": " + strconv.FormatFloat(angle, 'f', 6, 64) +"\n")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
	}
	f.Close();
}

func main() {
	result_writer(calc_angle(input_parser()))
}
