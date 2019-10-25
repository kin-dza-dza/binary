package Binary

import (
	"errors"
)

func Search (desiredNum int, Arr []int) (int, error, int) {
	// определяем границы поиска
	low := 0
	high := len(Arr) - 1

	for countIteration := 1; low <= high ; countIteration++{

		mid := (low + high) / 2 // ищем средний элемент
		guess := Arr[mid]		 // по элементу достаем число
		if guess == desiredNum { // нашли число, ура!
			return mid, nil, countIteration
		}

		if guess > desiredNum {  // если центральное число больше искомого, сдвигаем верхнюю границу
			high = mid -1
		}

		if guess < desiredNum {				 // если центрально число меньше искомого, сдвигаем нижнюю границу
			low = mid + 1
		}
	}

	return 0, errors.New("запрашиваемого числа не существует в списке"), 0
}