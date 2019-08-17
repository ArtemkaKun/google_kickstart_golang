package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func input_parser() (cases[][][]string){
	file, err := os.Open("Problem A\\2.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	number_of_cases, err := strconv.Atoi(scanner.Text())

	var test_cases[][][]string
	for test_case := 0; test_case < number_of_cases; test_case++ {
		var one_case[][]string
		//first line of each case shows number of pairs in this case
		scanner.Scan()
		number_of_strings, err := strconv.Atoi(scanner.Text())
		if err == nil {
			for one_pair := 0; one_pair < number_of_strings; one_pair++ {
				scanner.Scan()
				var one_char[]string
				buffer := strings.Fields(scanner.Text())
				//each line contains no more then 2 names (in complex names use snake case)
				one_char = append(one_char, buffer[0])
				one_char = append(one_char, buffer[1])
				one_case = append(one_case, one_char)
			}

			test_cases = append(test_cases, one_case)
		}

	}

	return test_cases
}

type bad_guy struct {
	this_guy string
	other_guys[]string
}

func split_process(parsed_data[][][]string) (can_split[]bool) {

	for one_case, _ := range parsed_data {
		var bad_guys[]bad_guy
		for _, value_pair := range parsed_data[one_case] {
			guy := bad_guy{}
			len_buffer := len(bad_guys)
			if len_buffer == 0 {
				guy.this_guy = value_pair[0]
				guy.other_guys = append(guy.other_guys, value_pair[1])
				bad_guys = append(bad_guys, guy)
			} else {
				var added = false;
				for id,_ := range bad_guys {
					if bad_guys[id].this_guy == value_pair[0] {
						bad_guys[id].other_guys = append(bad_guys[id].other_guys, value_pair[1])
						added = true;
					}
				}
				if !added {
					guy.this_guy = value_pair[0]
					guy.other_guys = append(guy.other_guys, value_pair[1])
					bad_guys = append(bad_guys, guy)
				}
			}
		}

			for _, value_pair := range parsed_data[one_case] {
				guy := bad_guy{}
					var added= false;
					for id, _ := range bad_guys {
						if bad_guys[id].this_guy == value_pair[1] {
							bad_guys[id].other_guys = append(bad_guys[id].other_guys, value_pair[0])
							added = true;
						}
					}
					if !added {
						guy.this_guy = value_pair[1]
						guy.other_guys = append(guy.other_guys, value_pair[0])
						bad_guys = append(bad_guys, guy)
					}
				}

		var first_group[]bad_guy
		var second_group = bad_guys[:]

		conflict : for _, guy_value := range second_group {
			for _, value := range second_group {
				for _, val := range value.other_guys {
					if guy_value.this_guy == val {
						first_group = append(first_group, guy_value)
						second_group = second_group[1:]
						goto conflict
					}
				}
			}
		}

		for _, guy_value := range first_group {
			for _, value := range first_group {
				for _, val := range value.other_guys {
					if guy_value.this_guy == val {
						can_split = append(can_split, false)
						goto cant_split
					}
				}
			}
		}
		can_split = append(can_split, true)
		continue

		cant_split :
			continue
	}

	return can_split
}

func result_writer(decide_split[]bool) {
	f, err := os.Create("result.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for case_id,decision := range decide_split {
		if decision {
			_, err := f.WriteString("Case #" + strconv.Itoa(case_id + 1) + ": Yes\n")
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
		} else {
			_, err := f.WriteString("Case #" + strconv.Itoa(case_id + 1) + ": No\n")
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
		}
	}
	f.Close();
}

func main() {
	var result = split_process(input_parser())
	result_writer(result)
}
