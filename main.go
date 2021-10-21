package main

import (
	"fmt"
	"strconv"
)

func main() {

	var (
		number, sum int
		nums        = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		vars        [][]int
		str_result  string
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

	result, _ := find(vars, number)

	for i, elem := range result {
		sum += elem
		if i != 0 && elem >= 0 {
			str_result += "+"
		}
		str_result += strconv.Itoa(elem)
	}

	if sum != number {
		fmt.Println("Fail")
		return
	}

	str_result += " = "
	str_result += strconv.Itoa(number)

	fmt.Println(str_result)

}

func find(vars [][]int, number int) ([]int, int) {

	var result []int
	ln := len(vars[0])

	if len(vars) == 2 {
		for i := 0; i < ln; i++ {

			if number == 0 {
				return []int{vars[0][0]}, number
			}

			if number-vars[0][i] == 0 {
				if vars[0][i]/10 != 0 {
					return []int{vars[0][i]}, vars[0][i]
				} else {
					return []int{vars[0][i], 0}, vars[0][i]
				}
			}
		}
	} else if len(vars) == 1 {
		return []int{vars[0][0]}, vars[0][0]
	} else {
		for i := 0; i < ln; i++ {
			if vars[0][i]/10 == 0 {
				if new_vars, new_num := find(vars[1:], number-vars[0][i]); new_num+vars[0][i] == number {
					result = append(result, vars[0][i])
					for _, elem := range new_vars {
						result = append(result, elem)
					}
					return result, new_num + vars[0][i]
				}
			} else {
				if new_vars, new_num := find(vars[2:], number-vars[0][i]); new_num+vars[0][i] == number {
					result = append(result, vars[0][i])
					for _, elem := range new_vars {
						result = append(result, elem)
					}
					return result, new_num + vars[0][i]
				}
			}
		}
		return result, 0
	}
	return result, 0
}
