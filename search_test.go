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

func TestSearch(t *testing.T) {
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