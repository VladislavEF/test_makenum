package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		number     int
		nums       = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		vars       [][]int
		result     [][]int
		str_result []string
	)
	number = 200
	//fmt.Println("Enter number")
	//fmt.Scan(&number)

	for i := 0; i < len(nums); i++ {
		vars = append(vars, []int{})
		for j, sum := i, 0; j < len(nums); j++ {
			sum = sum*10 + nums[j]
			if sum == 0 {
				vars[i] = append(vars[i], sum)
			} else if sum <= number || sum/10 == 0 {
				vars[i] = append(vars[i], sum)
				vars[i] = append(vars[i], -sum)
			}
		}
	}

	result = find(vars, number)

	for i, sub_arr := range result {
		str_result = append(str_result, "")
		for j, elem := range sub_arr {
			if j != 0 && elem >= 0 {
				str_result[i] += "+"
			}
			str_result[i] += strconv.Itoa(elem)
		}
		str_result[i] += " = "
		str_result[i] += strconv.Itoa(number)
	}

	for _, sub := range str_result {
		fmt.Println(sub)
	}
}

func find(comb [][]int, number int) [][]int {

	var result [][]int

	if len(comb) == 1 {

		result = append(result, []int{})

		if number == 0 {
			result[0] = append(result[0], comb[0][0])
			return result
		} else {
			return nil
		}
	} else if len(comb) == 2 {

		result = append(result, []int{})

		for i := 0; i < len(comb[0]); i++ {

			if number-comb[0][i] == 0 {
				if comb[0][i]%10 == 0 {
					result[0] = append(result[0], comb[0][i])
					return result
				} else {
					result[0] = append(result[0], comb[0][i], 0)
					return result
				}
			}
		}
		return nil
	} else {

		var local_res [][]int
		var cnt int

		for i := 0; i < len(comb[0]); i++ {

			if comb[0][i]/10 == 0 {
				local_res = find(comb[1:], number-comb[0][i])
			} else {
				local_res = find(comb[2:], number-comb[0][i])
			}

			if local_res != nil {

				for _, sub_arr := range local_res {
					var buf []int
					buf = append(buf, comb[0][i])
					for _, elem := range sub_arr {
						buf = append(buf, elem)
					}
					cnt++
					result = append(result, buf)
				}
			}
		}
	}
	return result
}
