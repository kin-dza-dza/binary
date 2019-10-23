package out

import "fmt"

func Output(element int, err error, countIteration int) {
	if err == nil {
		fmt.Println("ваше число под номером", element)
		fmt.Println("шагов сделано", countIteration)
	} else {
		fmt.Println(err)
		fmt.Println("шагов сделано", countIteration)
	}
}