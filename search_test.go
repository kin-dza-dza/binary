package binary

import (
	"errors"
	"math/rand"
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

	t.Run("Массив не сортирован, искомого числа нет", func(t *testing.T) {
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

	t.Run("Массив из одного элемента", func(t *testing.T) {
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

	t.Run("Массив из одного нуля", func(t *testing.T) {
		Arr := []int {0}
		desiredNum := 0

		wantElement := 0

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr != nil {
			t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
		}
	})

	t.Run("Не инициализированный массив", func(t *testing.T) {
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

	t.Run("Ищем отрицательное число", func(t *testing.T) {
		Arr := []int {-10,-8,-6,-4,-2}
		desiredNum := -4

		wantElement := 3

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr != nil {
			t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
		}
	})

	t.Run("Ищем ноль", func(t *testing.T) {
		Arr := []int {-10,-8,0,2,4}
		desiredNum := 0

		wantElement := 2

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr != nil {
			t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
		}
	})

	t.Run("Ищем большое положительное", func(t *testing.T) {
		Arr := []int {5,100,32767,32768,100000000000000000}
		desiredNum := 100000000000000000

		wantElement := 4

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr != nil {
			t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
		}
	})

	t.Run("Ищем большое отрицательное", func(t *testing.T) {
		Arr := []int {-100000000000000000,-32768,-32767,-100,-10}
		desiredNum := -100000000000000000

		wantElement := 0

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr != nil {
			t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
		}
	})

	t.Run("В массиве есть ноль", func(t *testing.T) {
		Arr := []int {0,100,101,102,1000}
		desiredNum := 100

		wantElement := 1

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr != nil {
			t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
		}
	})

	t.Run("Массив с отрицательными и положительными значениями", func(t *testing.T) {
		Arr := []int {-20,-10,0,102,1000}
		desiredNum := 102

		wantElement := 3

		gotElememt, gotErr := Search(desiredNum, Arr)

		if wantElement != gotElememt {
			t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
		}

		if gotErr != nil {
			t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
		}
	})

	t.Run("Несколько искомых значений", func(t *testing.T) {
		Arr := []int {5,6,6,6,7}
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

	t.Run("Большие массивы", func(t *testing.T) {

		t.Run("Массив на 100 элементов", func(t *testing.T) {
			Arr := makeArr(100)
			desiredNum := 50

			wantElement := 49
			gotElememt, gotErr := Search(desiredNum, Arr)

			if wantElement != gotElememt {
				t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
			}

			if gotErr != nil {
				t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
			}
		})

		t.Run("Массив на 10000 элементов", func(t *testing.T) {
			Arr := makeArr(10000)
			desiredNum := 5000

			wantElement := 4999
			gotElememt, gotErr := Search(desiredNum, Arr)

			if wantElement != gotElememt {
				t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
			}

			if gotErr != nil {
				t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
			}
		})

		t.Run("Массив на 1000000 элементов", func(t *testing.T) {
			Arr := makeArr(1000000)
			desiredNum := 500000

			wantElement := 499999
			gotElememt, gotErr := Search(desiredNum, Arr)

			if wantElement != gotElememt {
				t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
			}

			if gotErr != nil {
				t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
			}
		})

		t.Run("Массив на 100000000 элементов", func(t *testing.T) {
			Arr := makeArr(100000000)
			desiredNum := 5000000

			wantElement := 4999999
			gotElememt, gotErr := Search(desiredNum, Arr)

			if wantElement != gotElememt {
				t.Error("возвращает неправильный элемент", gotElememt, ". Должен возвращать", wantElement)
			}

			if gotErr != nil {
				t.Errorf("возвращает ошибку %q. Должно быть nil", gotErr)
			}
		})
	})
}

// техническая функция. Создает массив. Нужна для тестов.
func makeArr (amountOfElements int) []int {
	Arr := make([]int, amountOfElements)
	for i := 0; i < amountOfElements; i++ {
		Arr[i] = i + 1
	}

	return Arr
}

func makeArrRand () []int {
	amountOfElements := rand.Intn(1000000)
	Arr := make([]int, amountOfElements)
	for i := 0; i < amountOfElements; i++ {
		Arr[i] = rand.Int()
	}

	return Arr
}