package binary

import (
	"errors"
	"testing"
)

func TestSearch (t *testing.T) {

	t.Run("Искомого числа нет в массиве", func(t *testing.T) {
		Arr := []int {2,4,6,8,10}
		desiredNum := 5

		wantElement := 0
		wantErr := errors.New("запрашиваемого числа не существует в списке")

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr == nil {
			t.Errorf("не возвращает ошибку. Должен вернуть %q", wantErr)
		}
	})

	t.Run("Искомое число есть в массиве", func(t *testing.T) {
		Arr := []int {2,4,6,8,10}
		desiredNum := 6

		wantElement := 2

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr != nil {
			t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
		}
	})

	t.Run("Массив не сортирован, искомое число есть", func(t *testing.T) {
		Arr := []int {2,6,4,8,10}
		desiredNum := 6

		wantElement := 2

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr != nil {
			t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
		}
	})

	t.Run("Искомого числа нет в массиве", func(t *testing.T) {
		Arr := []int {2,6,4,8,10}
		desiredNum := 5

		wantElement := 0
		wantErr := errors.New("запрашиваемого числа не существует в списке")

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr == nil {
			t.Errorf("не возвращает ошибку. Должен вернуть %q", wantErr)
		}
	})

	t.Run("массив из одного элемента", func(t *testing.T) {
		Arr := []int {2}
		desiredNum := 2

		wantElement := 0

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr != nil {
			t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
		}
	})

	t.Run("не инициализированный массив", func(t *testing.T) {
		var Arr []int
		desiredNum := 5

		wantElement := 0
		wantErr := errors.New("запрашиваемого числа не существует в списке")

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr == nil {
			t.Errorf("не возвращает ошибку. Должен вернуть %q", wantErr)
		}
	})
}