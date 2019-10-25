package binary

import (
	"errors"
	"testing"
)

func NewArr () []int { //техническая функция для тестов. Делает новый упорядоченный массив
	Arr := make([]int, 5)

	for i := 0; i < len(Arr); i++ {
		Arr[i] = (i * 2) + 2
	}

	return Arr
}

func Test_deepSearch (t *testing.T) {
	Arr := NewArr()

	t.Run("Искомого числа нет в массиве", func(t *testing.T) {
		desiredNum := 5

		wantElement := 0
		wantErr := errors.New("запрашиваемого числа не существует в списке")
		wantCountIteration := 0

		gotElememt, gotErr, gotCountIteration := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if wantCountIteration != gotCountIteration {
			t.Error("возвращает неправильный счетчик шагов", gotCountIteration, ". Должен возвращать", wantCountIteration)
		}

		if gotErr == nil {
			t.Errorf("не возвращает ошибку. Должен вернуть %q", wantErr)
		}
	})

	t.Run("Искомое число есть в массиве", func(t *testing.T) {
		desiredNum := 6

		wantElement := 2
		wantCountIteration := 1

		gotElememt, gotErr, gotCountIteration := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if wantCountIteration != gotCountIteration {
			t.Error("возвращает неправильный счетчик шагов", gotCountIteration, ". Должен возвращать", wantCountIteration)
		}

		if gotErr != nil {
			t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
		}
	})
}

func Test_checkSort (t *testing.T) {

	t.Run("Массив из 1 элемента", func(t *testing.T) {
		Arr := []int {1}

		if !checkSort(Arr) {
			t.Error("Считает, что одноэлементый массив не отсортирован. КАК?!?")
		}
	})

	t.Run("Сортированный массив", func(t *testing.T) {
		Arr := []int {1,2,3}

		if !checkSort(Arr) {
			t.Error("Сортированный массив принимает за несортированный")
		}
	})

	t.Run("Несортированный массив", func(t *testing.T) {
		Arr := []int {1,2,3,2}

		if checkSort(Arr) {
			t.Error("Несортированный массив принимает за сортированный")
		}
	})

	t.Run("Неинициализированный массив", func(t *testing.T) {
		var Arr []int

		if !checkSort(Arr) {
			t.Error("Считает, что неинициализированный массив не отсортирован. КАК?!?")
		}
	})
}