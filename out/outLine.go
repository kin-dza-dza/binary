package out

import "fmt"

func Output(element int, err error) {
	if err == nil {
		fmt.Println("ваше число под номером", element)
	} else {
		fmt.Println(err)
	}
}